package fleetdbapi

import (
	"encoding/json"
	"fmt"
	"maps"
	"strconv"
	"strings"

	"github.com/bmc-toolbox/common"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// CollectionMethod indicates how the data was collected
type CollectionMethod string

// ComponentSlugMap holds a lookup map for component slug to component type objs
type ComponentSlugMap map[string]*ServerComponentType

// Converter converts from the common.Device type to the fleetdbapi Server,Component types
//
// This is an exported type to enable fleetdbapi clients to publish server inventory data.
type Converter struct {
	skipSlugCheck    bool
	CollectionMethod CollectionMethod
	slugs            ComponentSlugMap
}

// Initializes and returns a new common.Device to fleetdbapi.Server converter
//
// skipSlugCheck when set will cause the convertor to not verify the components are of a valid ComponentSlugType in fleetdbapi
// this check should not be disabled for when the converted inventory has to be stored in fleetdb.
func NewComponentConverter(method CollectionMethod, slugs ComponentSlugMap, skipSlugCheck bool) *Converter {
	return &Converter{CollectionMethod: method, slugs: slugs, skipSlugCheck: skipSlugCheck}
}

var (
	// Inband identifies data collected through the ironlib image running on the host OS
	Inband CollectionMethod = "inband"
	// Outofband identifies data collected through the host BMC
	Outofband CollectionMethod = "outofband"
	ErrSlugs                   = errors.New("component slug error")
)

// FromCommonDevice returns a fleetdbapi.Server equivalent for a common.Device type
func (r *Converter) FromCommonDevice(serverID uuid.UUID, hw *common.Device) (*Server, error) {
	components, err := r.toComponentSlice(serverID, hw)
	if err != nil {
		return nil, err
	}

	return &Server{
		UUID:       serverID,
		Serial:     hw.Serial,
		Model:      common.FormatProductName(hw.Model),
		Vendor:     common.FormatVendorName(hw.Vendor),
		Components: components,
	}, nil
}

// toComponentSlice converts a common.Device object into a slice of components along with its attributes
func (r *Converter) toComponentSlice(serverID uuid.UUID, hw *common.Device) (ServerComponentSlice, error) {
	hwVendor := common.FormatVendorName(hw.Vendor)
	hwModel := common.FormatVendorName(hw.Model)

	// populate singular components
	componentsTmp := []*ServerComponent{
		r.bios(hwVendor, hwModel, hw.BIOS),
		r.bmc(hw.Vendor, hw.Model, hw.BMC),
		r.mainboard(hw.Vendor, hw.Model, hw.Mainboard),
	}

	// populate multiples of components
	componentsTmp = append(componentsTmp, r.dimms(hw.Memory)...)
	componentsTmp = append(componentsTmp, r.nics(hw.NICs)...)
	componentsTmp = append(componentsTmp, r.drives(hw.Drives)...)
	componentsTmp = append(componentsTmp, r.psus(hw.PSUs)...)
	componentsTmp = append(componentsTmp, r.cpus(hw.CPUs)...)
	componentsTmp = append(componentsTmp, r.tpms(hw.TPMs)...)
	componentsTmp = append(componentsTmp, r.cplds(hw.CPLDs)...)
	componentsTmp = append(componentsTmp, r.gpus(hw.GPUs)...)
	componentsTmp = append(componentsTmp, r.storageControllers(hw.StorageControllers)...)
	componentsTmp = append(componentsTmp, r.enclosures(hwVendor, hwModel, hw.Enclosures)...)

	final := []*ServerComponent{}
	for _, component := range componentsTmp {
		if component == nil {
			continue
		}

		if r.isRequiredAttributesEmpty(component) {
			zap.L().Info("component ignored, required attributes missing: " + component.Name)
			continue
		}

		component.ServerUUID = serverID
		final = append(final, component)
	}

	return final, nil
}

func (r *Converter) isRequiredAttributesEmpty(component *ServerComponent) bool {
	return component.Serial == "0" &&
		component.Model == "" &&
		component.Vendor == ""
}

func setFirmware(c *ServerComponent, fw *common.Firmware) {
	if fw == nil || fw.Installed == "" {
		return
	}

	c.InstalledFirmware = &InstalledFirmware{Version: fw.Installed}
}

func setStatus(c *ServerComponent, status *common.Status) {
	if status == nil || (status.Health == "" && status.State == "") {
		return
	}

	c.Status = &ComponentStatus{
		Health: status.Health,
		State:  status.State,
	}
}

func setCapabilities(c *ServerComponent, caps []*common.Capability) {
	found := []*ComponentCapability{}

	for _, cap := range caps {
		if cap == nil || cap.Name == "" {
			continue
		}

		found = append(
			found,
			&ComponentCapability{
				Name:        cap.Name,
				Description: cap.Description,
				Enabled:     cap.Enabled,
			},
		)
	}

	if len(found) > 0 {
		c.Capabilities = found
	}
}

func setMetadata(c *ServerComponent, metadata map[string]string) {
	if metadata == nil {
		return
	}

	// Note: The input metadata map should not be re-initialized,
	// since it may be called multiple times on a component.
	found := map[string]string{}

	for k, v := range metadata {
		if k == "" || v == "" {
			continue
		}

		found[k] = v
	}

	if len(found) == 0 {
		return
	}

	data, err := json.Marshal(found)
	if err != nil {
		zap.L().Warn("error in conversion of generic metadata attributes to json",
			zap.String("name", c.Name),
			zap.String("kind", fmt.Sprintf("%T", data)),
			zap.Error(err),
		)
	}

	c.Metadata = append(
		c.Metadata,
		&ComponentMetadata{
			Namespace: ComponentMetadataGenericNS,
			Data:      data,
		},
	)
}

// TODO: add OEM field
func (r *Converter) newComponent(
	slug,
	cvendor,
	cmodel,
	cserial,
	cproduct,
	description string,
	firmware *common.Firmware,
	status *common.Status,
	capabilities []*common.Capability,
	metadata map[string]string,
) (*ServerComponent, error) {
	// lower case slug to changeObj how its stored in server service
	slug = strings.ToLower(slug)

	if !r.skipSlugCheck {
		// component slug lookup map is expected
		if len(r.slugs) == 0 {
			return nil, errors.Wrap(ErrSlugs, "component slugs lookup map empty")
		}

		// component slug is part of the lookup map
		_, exists := r.slugs[slug]
		if !exists {
			return nil, errors.Wrap(ErrSlugs, "unknown component slug: "+slug)
		}
	}

	// use the product name when model number is empty
	if strings.TrimSpace(cmodel) == "" && strings.TrimSpace(cproduct) != "" {
		cmodel = cproduct
	}

	component := &ServerComponent{
		Name:        slug,
		Description: description,
		Vendor:      common.FormatVendorName(cvendor),
		Model:       common.FormatProductName(cmodel),
		Serial:      cserial,
	}

	setFirmware(component, firmware)
	setStatus(component, status)
	setCapabilities(component, capabilities)
	setMetadata(component, metadata)

	return component, nil
}

// clone map or create a new one if current is nil
func (r *Converter) cloneMap(current map[string]string) map[string]string {
	if current == nil {
		return map[string]string{}
	}

	return maps.Clone(current)
}

func (r *Converter) logz(err error, slug, serial string) {
	zap.L().Warn("component conversion error", zap.Error(err), zap.String("slug", slug), zap.String("serial", serial))
}

func (r *Converter) gpus(gpus []*common.GPU) []*ServerComponent {
	if gpus == nil {
		return nil
	}

	components := make([]*ServerComponent, 0, len(gpus))

	slug := common.SlugGPU
	for idx, c := range gpus {
		if strings.TrimSpace(c.Serial) == "" {
			c.Serial = strconv.Itoa(idx)
		}

		component, err := r.newComponent(
			slug,
			c.Vendor,
			c.Model,
			c.Serial,
			c.ProductName,
			c.Description,
			c.Firmware,
			c.Status,
			c.Capabilities,
			c.Metadata,
		)
		if err != nil {
			r.logz(err, slug, c.Serial)
			return nil
		}

		components = append(components, component)
	}

	return components
}

func (r *Converter) cplds(cplds []*common.CPLD) []*ServerComponent {
	if cplds == nil {
		return nil
	}

	components := make([]*ServerComponent, 0, len(cplds))

	slug := common.SlugCPLD
	for idx, c := range cplds {
		if strings.TrimSpace(c.Serial) == "" {
			c.Serial = strconv.Itoa(idx)
		}

		component, err := r.newComponent(
			slug,
			c.Vendor,
			c.Model,
			c.Serial,
			c.ProductName,
			c.Description,
			c.Firmware,
			c.Status,
			c.Capabilities,
			c.Metadata,
		)
		if err != nil {
			r.logz(err, slug, c.Serial)
			return nil
		}

		components = append(components, component)
	}

	return components
}

func (r *Converter) tpms(tpms []*common.TPM) []*ServerComponent {
	if tpms == nil {
		return nil
	}

	components := make([]*ServerComponent, 0, len(tpms))

	slug := common.SlugTPM
	for idx, c := range tpms {
		if strings.TrimSpace(c.Serial) == "" {
			c.Serial = strconv.Itoa(idx)
		}

		metadata := r.cloneMap(c.Metadata)
		metadata["interface_type"] = c.InterfaceType

		component, err := r.newComponent(
			slug,
			c.Vendor,
			c.Model,
			c.Serial,
			c.ProductName,
			c.Description,
			c.Firmware,
			c.Status,
			c.Capabilities,
			metadata,
		)
		if err != nil {
			r.logz(err, slug, c.Serial)
			return nil
		}

		components = append(components, component)
	}

	return components
}

func (r *Converter) cpus(cpus []*common.CPU) []*ServerComponent {
	if cpus == nil {
		return nil
	}

	components := make([]*ServerComponent, 0, len(cpus))

	slug := common.SlugCPU
	for idx, c := range cpus {
		if strings.TrimSpace(c.Serial) == "" {
			c.Serial = strconv.Itoa(idx)
		}

		metadata := r.cloneMap(c.Metadata)
		metadata["slot"] = c.Slot
		metadata["architecture"] = c.Architecture
		metadata["clock_speed_hz"] = fmt.Sprintf("%d", c.ClockSpeedHz)
		metadata["cores"] = fmt.Sprintf("%d", c.Cores)
		metadata["threads"] = fmt.Sprintf("%d", c.Threads)

		component, err := r.newComponent(
			slug,
			c.Vendor,
			c.Model,
			c.Serial,
			c.ProductName,
			c.Description,
			c.Firmware,
			c.Status,
			c.Capabilities,
			metadata,
		)
		if err != nil {
			r.logz(err, slug, c.Serial)
			return nil
		}

		components = append(components, component)
	}

	return components
}

func (r *Converter) storageControllers(controllers []*common.StorageController) []*ServerComponent {
	if controllers == nil {
		return nil
	}

	components := make([]*ServerComponent, 0, len(controllers))

	serials := map[string]bool{}

	slug := common.SlugStorageController
	for idx, c := range controllers {
		if strings.TrimSpace(c.Serial) == "" {
			c.Serial = strconv.Itoa(idx)
		}

		// Storage controllers can show up with pci IDs as their serial number
		// set a unique serial on those components
		_, exists := serials[c.Serial]
		if exists {
			c.Serial = c.Serial + "-fleetdb-" + strconv.Itoa(idx)
		} else {
			serials[c.Serial] = true
		}

		metadata := r.cloneMap(c.Metadata)
		metadata["id"] = c.ID
		metadata["supported_controller_protocols"] = c.SupportedControllerProtocols
		metadata["supported_device_protocols"] = c.SupportedDeviceProtocols
		metadata["supported_raid_types"] = c.SupportedRAIDTypes
		metadata["physical_id"] = c.PhysicalID
		metadata["bus_info"] = c.BusInfo
		metadata["speed_bytes_per_sec"] = fmt.Sprintf("%d", int64(c.SpeedGbps*125000000)) // Gb/s to bytes

		component, err := r.newComponent(slug, c.Vendor,
			c.Model,
			c.Serial,
			c.ProductName,
			c.Description,
			c.Firmware,
			c.Status,
			c.Capabilities,
			metadata,
		)
		if err != nil {
			r.logz(err, slug, c.Serial)
			return nil
		}

		// some controller show up with model numbers in the description field.
		if component.Model == "" && c.Description != "" {
			component.Model = c.Description
		}

		components = append(components, component)
	}

	return components
}

func (r *Converter) psus(psus []*common.PSU) []*ServerComponent {
	if psus == nil {
		return nil
	}

	components := make([]*ServerComponent, 0, len(psus))

	slug := common.SlugPSU
	for idx, c := range psus {
		trimedSerial := strings.TrimSpace(c.Serial)
		if trimedSerial == "" || strings.Contains(trimedSerial, "To Be Filled By O.E.M.") {
			c.Serial = strconv.Itoa(idx)
		}

		metadata := r.cloneMap(c.Metadata)
		metadata["id"] = c.ID
		metadata["power_capacity_watts"] = fmt.Sprintf("%d", c.PowerCapacityWatts)

		component, err := r.newComponent(slug, c.Vendor,
			c.Model,
			c.Serial,
			c.ProductName,
			c.Description,
			c.Firmware,
			c.Status,
			c.Capabilities,
			metadata,
		)
		if err != nil {
			r.logz(err, slug, c.Serial)
			return nil
		}

		components = append(components, component)
	}

	return components
}

func (r *Converter) drives(drives []*common.Drive) []*ServerComponent {
	if drives == nil {
		return nil
	}

	components := make([]*ServerComponent, 0, len(drives))

	slug := common.SlugDrive
	for idx, c := range drives {
		if strings.TrimSpace(c.Serial) == "" {
			c.Serial = strconv.Itoa(idx)
		}

		metadata := r.cloneMap(c.Metadata)
		metadata["bus_info"] = c.BusInfo
		metadata["oem_id"] = c.OemID
		metadata["storage_controller"] = c.StorageController
		metadata["protocol"] = c.Protocol
		metadata["type"] = c.Type
		metadata["wwn"] = c.WWN
		metadata["capacity_bytes"] = fmt.Sprintf("%d", c.CapacityBytes)
		metadata["blocksize_bytes"] = fmt.Sprintf("%d", c.BlockSizeBytes)
		metadata["capable_speed_bytes_per_sec"] = fmt.Sprintf("%d", int64(c.CapableSpeedGbps*125000000)) // Gb/s to bytes - 1000000000/8
		metadata["negotiated_speed_bytes_per_sec"] = fmt.Sprintf("%d", int64(c.NegotiatedSpeedGbps*125000000))
		metadata["smart_status"] = c.SmartStatus

		component, err := r.newComponent(common.SlugDrive, c.Vendor,
			c.Model,
			c.Serial,
			c.ProductName,
			c.Description,
			c.Firmware,
			c.Status,
			c.Capabilities,
			metadata,
		)
		if err != nil {
			r.logz(err, slug, c.Serial)
			return nil
		}

		// some drives show up with model numbers in the description field.
		if component.Model == "" && c.Description != "" {
			component.Model = c.Description
		}

		components = append(components, component)
	}

	return components
}

func (r *Converter) nics(nics []*common.NIC) []*ServerComponent {
	if nics == nil {
		return nil
	}

	components := make([]*ServerComponent, 0, len(nics))

	slug := common.SlugNIC
	for idx, c := range nics {
		if strings.TrimSpace(c.Serial) == "" {
			c.Serial = strconv.Itoa(idx)
		}

		component, err := r.newComponent(slug, c.Vendor,
			c.Model,
			c.Serial,
			c.ProductName,
			c.Description,
			c.Firmware,
			c.Status,
			c.Capabilities,
			c.Metadata,
		)
		if err != nil {
			r.logz(err, slug, c.Serial)
			return nil
		}

		components = append(components, r.nicPorts(c, c.NICPorts)...)
		components = append(components, component)
	}

	return components
}

func (r *Converter) nicPorts(nic *common.NIC, nicPorts []*common.NICPort) []*ServerComponent {
	components := make([]*ServerComponent, 0, len(nicPorts))

	slug := common.SlugNICPort
	for _, p := range nicPorts {
		if strings.TrimSpace(p.Serial) == "" && p.MacAddress != "" {
			p.Serial = p.MacAddress
		} else {
			// NIC ports with no serial/macaddresses are ignored
			return nil
		}

		metadata := r.cloneMap(p.Metadata)
		metadata["link_status"] = p.LinkStatus
		metadata["active_link_technology"] = p.ActiveLinkTechnology
		metadata["auto_negotiate"] = strconv.FormatBool(p.AutoNeg)
		metadata["mtu_size"] = fmt.Sprintf("%d", p.MTUSize)
		metadata["bus_info"] = p.BusInfo
		metadata["physical_id"] = p.PhysicalID
		metadata["speed_bits_per_second"] = fmt.Sprintf("%d", p.SpeedBits)

		component, err := r.newComponent(
			slug,
			nic.Vendor,
			nic.Model,
			p.Serial,
			p.ProductName,
			p.Description,
			p.Firmware,
			p.Status,
			p.Capabilities,
			metadata,
		)
		if err != nil {
			r.logz(err, slug, p.Serial)
			return nil
		}

		components = append(components, component)
	}

	return components
}

func (r *Converter) dimms(dimms []*common.Memory) []*ServerComponent {
	if dimms == nil {
		return nil
	}

	components := make([]*ServerComponent, 0, len(dimms))

	slug := common.SlugPhysicalMem
	for idx, c := range dimms {
		// skip empty dimm slots
		if c.Vendor == "" && c.ProductName == "" && c.SizeBytes == 0 && c.ClockSpeedHz == 0 {
			continue
		}

		// set incrementing serial when one isn't found
		if strings.TrimSpace(c.Serial) == "" {
			c.Serial = strconv.Itoa(idx)
		}

		// trim redundant prefix
		c.Slot = strings.TrimPrefix(c.Slot, "DIMM.Socket.")

		metadata := r.cloneMap(c.Metadata)
		metadata["slot"] = c.Slot
		metadata["clock_speed_hz"] = fmt.Sprintf("%d", c.ClockSpeedHz)
		metadata["size_bytes"] = fmt.Sprintf("%d", c.SizeBytes)
		metadata["form_factor"] = c.FormFactor
		metadata["part_number"] = c.PartNumber

		component, err := r.newComponent(
			slug,
			c.Vendor,
			c.Model,
			c.Serial,
			c.ProductName,
			c.Description,
			c.Firmware,
			c.Status,
			c.Capabilities,
			metadata,
		)
		if err != nil {
			r.logz(err, slug, c.Serial)
			return nil
		}

		components = append(components, component)
	}

	return components
}

func (r *Converter) mainboard(hwVendor, hwModel string, c *common.Mainboard) *ServerComponent {
	if c == nil {
		return nil
	}

	if strings.TrimSpace(c.Serial) == "" {
		c.Serial = "0"
	}

	if c.Vendor == "" {
		c.Vendor = hwVendor
	}

	if c.Model == "" {
		c.Model = hwModel
	}

	slug := common.SlugMainboard
	component, err := r.newComponent(
		slug,
		c.Vendor,
		c.Model,
		c.Serial,
		c.ProductName,
		c.Description,
		c.Firmware,
		c.Status,
		c.Capabilities,
		c.Metadata,
	)
	if err != nil {
		r.logz(err, slug, c.Serial)

		return nil
	}

	return component
}

func (r *Converter) enclosures(hwVendor, hwModel string, enclosures []*common.Enclosure) []*ServerComponent {
	if enclosures == nil {
		return nil
	}

	components := make([]*ServerComponent, 0, len(enclosures))

	slug := common.SlugEnclosure
	for idx, c := range enclosures {
		if strings.TrimSpace(c.Serial) == "" {
			c.Serial = strconv.Itoa(idx)
		}

		if c.Vendor == "" {
			c.Vendor = hwVendor
		}

		if c.Model == "" {
			c.Model = hwModel
		}

		metadata := r.cloneMap(c.Metadata)
		metadata["chassis_type"] = c.ChassisType

		component, err := r.newComponent(
			slug,
			c.Vendor,
			c.Model,
			c.Serial,
			c.ProductName,
			c.Description,
			c.Firmware,
			c.Status,
			c.Capabilities,
			metadata,
		)
		if err != nil {
			r.logz(err, slug, c.Serial)
			return nil
		}

		components = append(components, component)
	}

	return components
}

func (r *Converter) bmc(hwVendor, hwModel string, c *common.BMC) *ServerComponent {
	if c == nil {
		return nil
	}

	if strings.TrimSpace(c.Serial) == "" {
		c.Serial = "0"
	}

	if c.Vendor == "" {
		c.Vendor = hwVendor
	}

	if c.Model == "" {
		c.Model = hwModel
	}

	slug := common.SlugBMC
	component, err := r.newComponent(
		slug,
		c.Vendor,
		c.Model,
		c.Serial,
		c.ProductName,
		c.Description,
		c.Firmware,
		c.Status,
		c.Capabilities,
		c.Metadata,
	)
	if err != nil {
		r.logz(err, slug, c.Serial)

		return nil
	}

	return component
}

func (r *Converter) bios(hwVendor, hwModel string, c *common.BIOS) *ServerComponent {
	if c == nil {
		return nil
	}

	if strings.TrimSpace(c.Serial) == "" {
		c.Serial = "0"
	}

	if c.Vendor == "" {
		c.Vendor = hwVendor
	}

	if c.Model == "" {
		c.Model = hwModel
	}

	slug := common.SlugBIOS

	metadata := r.cloneMap(c.Metadata)
	metadata["size_bytes"] = fmt.Sprintf("%d", c.SizeBytes)
	metadata["capacity_bytes"] = fmt.Sprintf("%d", c.CapacityBytes)

	component, err := r.newComponent(common.SlugBIOS, c.Vendor,
		c.Model,
		c.Serial,
		c.ProductName,
		c.Description,
		c.Firmware,
		c.Status,
		c.Capabilities,
		metadata,
	)
	if err != nil {
		r.logz(err, slug, c.Serial)

		return nil
	}

	return component
}

package fleetdbapi

import (
	"context"
	"fmt"
	"path"
	"time"

	"github.com/google/uuid"
)

const (
	serversEndpoint                     = "servers"
	serverAttributesEndpoint            = "attributes"
	serverComponentsEndpoint            = "components"
	serverVersionedAttributesEndpoint   = "versioned-attributes"
	serverComponentFirmwaresEndpoint    = "server-component-firmwares"
	serverCredentialsEndpoint           = "credentials"
	serverCredentialTypeEndpoint        = "server-credential-types" // nolint:gosec //false positive
	serverComponentFirmwareSetsEndpoint = "server-component-firmware-sets"
	serverBiosConfigSetEndpoint         = "server-bios-config-sets"
	bomInfoEndpoint                     = "bill-of-materials"
	uploadFileEndpoint                  = "batch-upload"
	bomByMacAOCAddressEndpoint          = "aoc-mac-address"
	bomByMacBMCAddressEndpoint          = "bmc-mac-address"
	inventoryEndpoint                   = "inventory"
	hardwareVendorsEndpoint             = "hardware-vendors"
	hardwareModelsEndpoint              = "hardware-models"
	serverBMCsEndpoint                  = "bmc"
	installedFirmwareEndpoint           = "installed-firmware"
	componentStatusEndpoint             = "component-status"
	serverStatusEndpoint                = "server-status"
	componentCapabilityEndpoint         = "component-capability"
	componentMetadataEndpoint           = "component-metadata"
	componentChangesEndpoint            = "component-changes"
)

// ClientInterface provides an interface for the expected calls to interact with a fleetdb api
type ClientInterface interface {
	Create(context.Context, Server) (*uuid.UUID, *ServerResponse, error)
	Delete(context.Context, Server) (*ServerResponse, error)
	Get(context.Context, uuid.UUID) (*Server, *ServerResponse, error)
	ListServers(context.Context, *ServerQueryParams) ([]Server, *ServerResponse, error)
	Update(context.Context, uuid.UUID, Server) (*ServerResponse, error)

	GetComponents(context.Context, uuid.UUID, *PaginationParams) ([]ServerComponent, *ServerResponse, error)
	ListComponents(context.Context, *ServerComponentListParams) ([]ServerComponent, *ServerResponse, error)
	CreateComponents(context.Context, uuid.UUID, ServerComponentSlice) (*ServerResponse, error)
	UpdateComponents(context.Context, uuid.UUID, ServerComponentSlice) (*ServerResponse, error)
	DeleteServerComponents(context.Context, uuid.UUID) (*ServerResponse, error)

	CreateServerComponentFirmware(context.Context, ComponentFirmwareVersion) (*uuid.UUID, *ServerResponse, error)
	DeleteServerComponentFirmware(context.Context, ComponentFirmwareVersion) (*ServerResponse, error)
	GetServerComponentFirmware(context.Context, uuid.UUID) (*ComponentFirmwareVersion, *ServerResponse, error)
	ListServerComponentFirmware(context.Context, *ComponentFirmwareVersionListParams) ([]ComponentFirmwareVersion, *ServerResponse, error)
	UpdateServerComponentFirmware(context.Context, uuid.UUID, ComponentFirmwareVersion) (*ServerResponse, error)

	CreateServerComponentFirmwareSet(context.Context, ComponentFirmwareSetRequest) (*uuid.UUID, *ServerResponse, error)
	UpdateComponentFirmwareSetRequest(context.Context, ComponentFirmwareSetRequest) (*uuid.UUID, *ServerResponse, error)
	GetServerComponentFirmwareSet(context.Context, uuid.UUID) (*ComponentFirmwareSet, *ServerResponse, error)
	ListServerComponentFirmwareSet(context.Context, *ComponentFirmwareSetListParams) ([]ComponentFirmwareSet, *ServerResponse, error)
	ListFirmwareSets(context.Context, *ComponentFirmwareSetListParams) ([]ComponentFirmwareSet, *ServerResponse, error)
	DeleteServerComponentFirmwareSet(context.Context, uuid.UUID) (*ServerResponse, error)
	ValidateFirmwareSet(context.Context, uuid.UUID, uuid.UUID, time.Time) error

	GetCredential(context.Context, uuid.UUID, string) (*ServerCredential, *ServerResponse, error)
	SetCredential(context.Context, uuid.UUID, string, string) (*ServerResponse, error)
	DeleteCredential(context.Context, uuid.UUID, string) (*ServerResponse, error)
	ListServerCredentialTypes(context.Context) (*ServerResponse, error)

	GetHistoryByID(context.Context, uuid.UUID) (*Event, *ServerResponse, error)
	GetServerEvents(context.Context, uuid.UUID) ([]*Event, *ServerResponse, error)
	UpdateEvent(context.Context, *Event) (*ServerResponse, error)

	CreateServerBiosConfigSet(context.Context, BiosConfigSet) (*uuid.UUID, *ServerResponse, error)
	GetServerBiosConfigSet(context.Context, uuid.UUID) (*BiosConfigSet, *ServerResponse, error)
	DeleteServerBiosConfigSet(context.Context, uuid.UUID) (*ServerResponse, error)
	ListServerBiosConfigSet(context.Context) (*ServerResponse, error)
	UpdateServerBiosConfigSet(context.Context, uuid.UUID, BiosConfigSet) (*ServerResponse, error)

	CreateHardwareVendor(context.Context, *HardwareVendor) (*ServerResponse, error)
	GetHardwareVendor(context.Context, string) (*HardwareVendor, *ServerResponse, error)
	ListHardwareVendors(context.Context) ([]*HardwareVendor, *ServerResponse, error)
	DeleteHardwareVendor(context.Context, string) (*ServerResponse, error)

	CreateHardwareModel(context.Context, *HardwareModel) (*ServerResponse, error)
	GetHardwareModel(context.Context, string) (*HardwareModel, *ServerResponse, error)
	ListHardwareModels(context.Context) ([]*HardwareModel, *ServerResponse, error)
	DeleteHardwareModel(context.Context, string) (*ServerResponse, error)
}

// Create will attempt to create a server in Hollow and return the new server's UUID
func (c *Client) Create(ctx context.Context, srv Server) (*uuid.UUID, *ServerResponse, error) {
	resp, err := c.post(ctx, serversEndpoint, srv)
	if err != nil {
		return nil, nil, err
	}

	u, err := uuid.Parse(resp.Slug)
	if err != nil {
		return nil, resp, err
	}

	return &u, resp, nil
}

// Delete will attempt to delete a server in Hollow and return an error on failure
func (c *Client) Delete(ctx context.Context, srv Server) (*ServerResponse, error) {
	return c.delete(ctx, fmt.Sprintf("%s/%s", serversEndpoint, srv.UUID))
}

// GetServer will return a server by it's UUID
func (c *Client) GetServer(ctx context.Context, srvUUID uuid.UUID, params *ServerQueryParams) (*Server, *ServerResponse, error) {
	endpoint := fmt.Sprintf("%s/%s", serversEndpoint, srvUUID)
	srv := &Server{}
	r := ServerResponse{Record: srv}

	if err := c.getWithParams(ctx, endpoint, params, &r); err != nil {
		return nil, nil, err
	}

	return srv, &r, nil
}

// ListServers will return all servers with optional params to filter the results
func (c *Client) ListServers(ctx context.Context, params *ServerQueryParams) ([]Server, *ServerResponse, error) {
	servers := &[]Server{}
	r := ServerResponse{Records: servers}

	if err := c.list(ctx, serversEndpoint, params, &r); err != nil {
		return nil, nil, err
	}

	return *servers, &r, nil
}

// Update will to update a server with the new values passed in
func (c *Client) Update(ctx context.Context, srvUUID uuid.UUID, srv Server) (*ServerResponse, error) {
	endpoint := fmt.Sprintf("%s/%s", serversEndpoint, srvUUID)
	return c.put(ctx, endpoint, srv)
}

// GetComponents will get all the components for a given server
func (c *Client) GetComponents(ctx context.Context, srvUUID uuid.UUID, params *ServerComponentGetParams) (ServerComponentSlice, *ServerResponse, error) {
	sc := &ServerComponentSlice{}
	r := ServerResponse{Records: sc}

	endpoint := fmt.Sprintf("%s/%s/%s", serversEndpoint, srvUUID, serverComponentsEndpoint)
	if err := c.list(ctx, endpoint, params, &r); err != nil {
		return nil, nil, err
	}

	return *sc, &r, nil
}

// ListComponents will get all the components matching the given parameters
func (c *Client) ListComponents(ctx context.Context, params *ServerComponentListParams) (ServerComponentSlice, *ServerResponse, error) {
	sc := &ServerComponentSlice{}
	r := ServerResponse{Records: sc}

	endpoint := fmt.Sprintf("%s/%s", serversEndpoint, serverComponentsEndpoint)
	if err := c.list(ctx, endpoint, params, &r); err != nil {
		return nil, nil, err
	}

	return *sc, &r, nil
}

// InitComponentCollection is to be called to initialize component and relational records for a server
// This will only create component records only if none already exist
// collectionMethod is one of inband/outofband
//
// note: The {Report,Accept}ComponentChanges methods are used to add/delete components
func (c *Client) InitComponentCollection(ctx context.Context, srvUUID uuid.UUID, components ServerComponentSlice, collectionMethod CollectionMethod) (*ServerResponse, error) {
	endpoint := path.Join(serversEndpoint, srvUUID.String(), serverComponentsEndpoint, "init", string(collectionMethod))
	return c.post(ctx, endpoint, components)
}

// UpdateComponentCollection will update existing component and related records for a server - this will not delete any or add any components.
//
// note: The {Report,Accept}ComponentChanges methods are used to add/delete components
func (c *Client) UpdateComponentCollection(ctx context.Context, srvUUID uuid.UUID, components ServerComponentSlice, collectionMethod CollectionMethod) (*ServerResponse, error) {
	endpoint := path.Join(serversEndpoint, srvUUID.String(), serverComponentsEndpoint, "update", string(collectionMethod))
	return c.put(ctx, endpoint, components)
}

// DeleteServerComponents will delete all components for the given server identifier.
func (c *Client) DeleteServerComponents(ctx context.Context, srvUUID uuid.UUID) (*ServerResponse, error) {
	return c.delete(ctx, fmt.Sprintf("%s/%s/%s", serversEndpoint, srvUUID, serverComponentsEndpoint))
}

// ReportComponentChanges creates records for server component additions/deletes - these have to be accepted first
func (c *Client) ReportComponentChanges(ctx context.Context, serverID string, change *ComponentChangeReport) (*ComponentChangeReportResponse, *ServerResponse, error) {
	cr := &ComponentChangeReportResponse{}
	res := &ServerResponse{Data: &cr}
	endpoint := path.Join(serversEndpoint, serverID, componentChangesEndpoint, "report")

	if err := c.postWithReciever(ctx, endpoint, change, res); err != nil {
		return nil, nil, err
	}

	return cr, res, nil
}

// AcceptComponentChanges accepts and merges the specified changeIDs which reference component addition/deletion component change reports
func (c *Client) AcceptComponentChanges(ctx context.Context, serverID string, changeIDs []string) (*ServerResponse, error) {
	endpoint := path.Join(serversEndpoint, serverID, componentChangesEndpoint, "accept")
	report := &ComponentChangeAccept{
		ChangeIDs: changeIDs,
	}

	return c.post(ctx, endpoint, report)
}

// CreateServerComponentFirmware will attempt to create a firmware in Hollow and return the firmware UUID
func (c *Client) CreateServerComponentFirmware(ctx context.Context, firmware ComponentFirmwareVersion) (*uuid.UUID, *ServerResponse, error) {
	resp, err := c.post(ctx, serverComponentFirmwaresEndpoint, firmware)
	if err != nil {
		return nil, nil, err
	}

	u, err := uuid.Parse(resp.Slug)
	if err != nil {
		return nil, resp, err
	}

	return &u, resp, nil
}

// DeleteServerComponentFirmware will attempt to delete firmware and return an error on failure
func (c *Client) DeleteServerComponentFirmware(ctx context.Context, firmware ComponentFirmwareVersion) (*ServerResponse, error) {
	return c.delete(ctx, fmt.Sprintf("%s/%s", serverComponentFirmwaresEndpoint, firmware.UUID))
}

// GetServerComponentFirmware will return a firmware by its UUID
func (c *Client) GetServerComponentFirmware(ctx context.Context, fwUUID uuid.UUID) (*ComponentFirmwareVersion, *ServerResponse, error) {
	endpoint := fmt.Sprintf("%s/%s", serverComponentFirmwaresEndpoint, fwUUID)
	fw := &ComponentFirmwareVersion{}
	r := ServerResponse{Record: fw}

	if err := c.get(ctx, endpoint, &r); err != nil {
		return nil, nil, err
	}

	return fw, &r, nil
}

// ListServerComponentFirmware will return all firmwares with optional params to filter the results
func (c *Client) ListServerComponentFirmware(ctx context.Context, params *ComponentFirmwareVersionListParams) ([]ComponentFirmwareVersion, *ServerResponse, error) {
	firmwares := &[]ComponentFirmwareVersion{}
	r := ServerResponse{Records: firmwares}

	if err := c.list(ctx, serverComponentFirmwaresEndpoint, params, &r); err != nil {
		return nil, nil, err
	}

	return *firmwares, &r, nil
}

// UpdateServerComponentFirmware will to update a firmware with the new values passed in
func (c *Client) UpdateServerComponentFirmware(ctx context.Context, fwUUID uuid.UUID, firmware ComponentFirmwareVersion) (*ServerResponse, error) {
	endpoint := fmt.Sprintf("%s/%s", serverComponentFirmwaresEndpoint, fwUUID)
	return c.put(ctx, endpoint, firmware)
}

// CreateServerComponentFirmwareSet will attempt to create a firmware set in Hollow and return the firmware UUID
func (c *Client) CreateServerComponentFirmwareSet(ctx context.Context, set ComponentFirmwareSetRequest) (*uuid.UUID, *ServerResponse, error) {
	resp, err := c.post(ctx, serverComponentFirmwareSetsEndpoint, set)
	if err != nil {
		return nil, nil, err
	}

	u, err := uuid.Parse(resp.Slug)
	if err != nil {
		return nil, resp, err
	}

	return &u, resp, nil
}

// DeleteServerComponentFirmwareSet will attempt to delete a firmware set and return an error on failure
func (c *Client) DeleteServerComponentFirmwareSet(ctx context.Context, firmwareSetID uuid.UUID) (*ServerResponse, error) {
	return c.delete(ctx, fmt.Sprintf("%s/%s", serverComponentFirmwareSetsEndpoint, firmwareSetID))
}

// GetServerComponentFirmwareSet will return a firmware by its UUID
func (c *Client) GetServerComponentFirmwareSet(ctx context.Context, fwSetUUID uuid.UUID) (*ComponentFirmwareSet, *ServerResponse, error) {
	endpoint := fmt.Sprintf("%s/%s", serverComponentFirmwareSetsEndpoint, fwSetUUID)
	fws := &ComponentFirmwareSet{}
	r := ServerResponse{Record: fws}

	if err := c.get(ctx, endpoint, &r); err != nil {
		return nil, nil, err
	}

	return fws, &r, nil
}

// ListServerComponentFirmwareSet will return all firmwares with optional params to filter the results
// if AttributeListParams is defined then ignore the main struct fields (Vendor, Model, Labels)
// otherwise do the selection based on the Vendor, Model, Labelswill
// return all firmwares with optional params to filter the results
// vendor and model should be non-empty. arbitraryLabels is formatted as k1=v1,k2=v2,etc.
// To view the behavior of the default/latest label, please check
// https://fleet-docs.pages.equinixmetal.net/procedures/firmware-install/#firmware-sets
func (c *Client) ListServerComponentFirmwareSet(ctx context.Context, params *ComponentFirmwareSetListParams) ([]ComponentFirmwareSet, *ServerResponse, error) {
	firmwareSets := &[]ComponentFirmwareSet{}
	r := ServerResponse{Records: firmwareSets}

	if err := c.list(ctx, serverComponentFirmwareSetsEndpoint, params, &r); err != nil {
		return nil, nil, err
	}

	return *firmwareSets, &r, nil
}

// UpdateComponentFirmwareSetRequest will add a firmware set with the new firmware id(s) passed in the firmwareSet parameter
func (c *Client) UpdateComponentFirmwareSetRequest(ctx context.Context, fwSetUUID uuid.UUID, firmwareSet ComponentFirmwareSetRequest) (*ServerResponse, error) {
	endpoint := fmt.Sprintf("%s/%s", serverComponentFirmwareSetsEndpoint, fwSetUUID)
	return c.put(ctx, endpoint, firmwareSet)
}

// RemoveServerComponentFirmwareSetFirmware will update a firmware set by removing the mapping for the firmware id(s) passed in the firmwareSet parameter
func (c *Client) RemoveServerComponentFirmwareSetFirmware(ctx context.Context, fwSetUUID uuid.UUID, firmwareSet ComponentFirmwareSetRequest) (*ServerResponse, error) {
	endpoint := fmt.Sprintf("%s/%s/remove-firmware", serverComponentFirmwareSetsEndpoint, fwSetUUID)
	return c.post(ctx, endpoint, firmwareSet)
}

// ValidateFirmwareSet inserts or updates a record containing facts about the validation of this
// particular firmware set. On a successful execution the API returns 204 (http.StatusNoContent), so
// there is nothing useful to put into a ServerResponse.
func (c *Client) ValidateFirmwareSet(ctx context.Context, srvID, fwSetID uuid.UUID, on time.Time) error {
	endpoint := fmt.Sprintf("%s/validate-firmware-set", serverComponentFirmwareSetsEndpoint)
	facts := FirmwareSetValidation{
		TargetServer: srvID,
		FirmwareSet:  fwSetID,
		PerformedOn:  on,
	}
	_, err := c.post(ctx, endpoint, facts)
	return err
}

// GetCredential will return the secret for the secret type for the given server UUID
func (c *Client) GetCredential(ctx context.Context, srvUUID uuid.UUID, secretSlug string) (*ServerCredential, *ServerResponse, error) {
	p := path.Join(serversEndpoint, srvUUID.String(), serverCredentialsEndpoint, secretSlug)
	secret := &ServerCredential{}
	r := ServerResponse{Record: secret}

	if err := c.get(ctx, p, &r); err != nil {
		return nil, nil, err
	}

	return secret, &r, nil
}

// SetCredential will set the secret for a given server UUID and secret type.
func (c *Client) SetCredential(ctx context.Context, srvUUID uuid.UUID, secretSlug, username, password string) (*ServerResponse, error) {
	p := path.Join(serversEndpoint, srvUUID.String(), serverCredentialsEndpoint, secretSlug)
	secret := &serverCredentialValues{
		Password: password,
		Username: username,
	}

	return c.put(ctx, p, secret)
}

// DeleteCredential will remove the secret for a given server UUID and secret type.
func (c *Client) DeleteCredential(ctx context.Context, srvUUID uuid.UUID, secretSlug string) (*ServerResponse, error) {
	p := path.Join(serversEndpoint, srvUUID.String(), serverCredentialsEndpoint, secretSlug)

	return c.delete(ctx, p)
}

// ListServerCredentialTypes will return all server secret types
func (c *Client) ListServerCredentialTypes(ctx context.Context, params *PaginationParams) ([]ServerCredentialType, *ServerResponse, error) {
	types := &[]ServerCredentialType{}
	r := ServerResponse{Records: types}

	if err := c.list(ctx, serverCredentialTypeEndpoint, params, &r); err != nil {
		return nil, nil, err
	}

	return *types, &r, nil
}

// CreateServerCredentialType will create a new server secret type
func (c *Client) CreateServerCredentialType(ctx context.Context, sType *ServerCredentialType) (*ServerResponse, error) {
	return c.post(ctx, serverCredentialTypeEndpoint, sType)
}

// GetHistoryByID returns the details of the event with the given ID
func (c *Client) GetHistoryByID(ctx context.Context, evtID uuid.UUID) ([]*Event, *ServerResponse, error) {
	evts := &[]*Event{}
	r := &ServerResponse{Records: evts}
	endpoint := fmt.Sprintf("events/%s", evtID.String())

	if err := c.get(ctx, endpoint, r); err != nil {
		return nil, nil, err
	}

	return *evts, r, nil
}

// GetServerEvents returns the most recent events for the given server ID
func (c *Client) GetServerEvents(ctx context.Context, srvID uuid.UUID,
	params *PaginationParams) ([]*Event, *ServerResponse, error) {
	evts := &[]*Event{}
	r := &ServerResponse{Records: evts}
	endpoint := fmt.Sprintf("events/by-server/%s", srvID.String())

	if err := c.list(ctx, endpoint, params, r); err != nil {
		return nil, nil, err
	}

	return *evts, r, nil
}

// UpdateEvent adds a new event to the event history
func (c *Client) UpdateEvent(ctx context.Context, evt *Event) (*ServerResponse, error) {
	endpoint := fmt.Sprintf("events/%s", evt.EventID.String())
	return c.put(ctx, endpoint, evt)
}

// CreateServerBiosConfigSet will store the BiosConfigSet, and return the generated UUID of the BiosConfigSet
func (c *Client) CreateServerBiosConfigSet(ctx context.Context, set BiosConfigSet) (*ServerResponse, error) {
	resp, err := c.post(ctx, serverBiosConfigSetEndpoint, set)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetServerBiosConfigSet will retrieve the BiosConfigSet referred to by the given ID if found
func (c *Client) GetServerBiosConfigSet(ctx context.Context, id uuid.UUID) (*ServerResponse, error) {
	endpoint := fmt.Sprintf("%s/%s", serverBiosConfigSetEndpoint, id)
	cfg := &BiosConfigSet{}
	resp := ServerResponse{Record: cfg}

	if err := c.get(ctx, endpoint, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// DeleteServerBiosConfigSet will delete the BiosConfigSet referred to by the given ID if found
func (c *Client) DeleteServerBiosConfigSet(ctx context.Context, id uuid.UUID) (*ServerResponse, error) {
	endpoint := fmt.Sprintf("%s/%s", serverBiosConfigSetEndpoint, id)
	resp, err := c.delete(ctx, endpoint)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// ListServerBiosConfigSet will return a list of BiosConfigSets referred to by the given query. More details about querying at the type definition of BiosConfigSetListParams.
func (c *Client) ListServerBiosConfigSet(ctx context.Context, params *BiosConfigSetListParams) (*ServerResponse, error) {
	cfg := &[]BiosConfigSet{}
	resp := ServerResponse{Records: cfg}

	err := c.list(ctx, serverBiosConfigSetEndpoint, params, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateServerBiosConfigSet will update a config set.
func (c *Client) UpdateServerBiosConfigSet(ctx context.Context, id uuid.UUID, set BiosConfigSet) (*ServerResponse, error) {
	endpoint := fmt.Sprintf("%s/%s", serverBiosConfigSetEndpoint, id)
	resp, err := c.put(ctx, endpoint, set)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateHardwareVendor creates a hardware vendor record
func (c *Client) CreateHardwareVendor(ctx context.Context, hardwareVendor *HardwareVendor) (*ServerResponse, error) {
	return c.post(ctx, hardwareVendorsEndpoint, hardwareVendor)
}

// ListHardwareVendors lists hardware vendor records
func (c *Client) ListHardwareVendors(ctx context.Context) ([]*HardwareVendor, *ServerResponse, error) {
	hardwareVendors := []*HardwareVendor{}
	resp := ServerResponse{Records: &hardwareVendors}

	if err := c.list(ctx, hardwareVendorsEndpoint, nil, &resp); err != nil {
		return nil, nil, err
	}

	return hardwareVendors, &resp, nil
}

// GetHardwareVendor retrieves a hardware vendor record by its name
func (c *Client) GetHardwareVendor(ctx context.Context, name string) (*HardwareVendor, *ServerResponse, error) {
	hardwareVendor := &HardwareVendor{}
	resp := ServerResponse{Record: hardwareVendor}

	endpoint := path.Join(hardwareVendorsEndpoint, name)
	if err := c.get(ctx, endpoint, &resp); err != nil {
		return nil, nil, err
	}

	return hardwareVendor, &resp, nil
}

// DeleteHardwareVendor purges a hardware vendor record by its name
func (c *Client) DeleteHardwareVendor(ctx context.Context, name string) (*ServerResponse, error) {
	endpoint := path.Join(hardwareVendorsEndpoint, name)
	return c.delete(ctx, endpoint)
}

// CreateHardwareModel creates a hardware model record - requires the hardware vendor relation
func (c *Client) CreateHardwareModel(ctx context.Context, hardwareModel *HardwareModel) (*ServerResponse, error) {
	return c.post(ctx, hardwareModelsEndpoint, hardwareModel)
}

// ListHardwareModels lists hardware vendor model records
func (c *Client) ListHardwareModels(ctx context.Context) ([]*HardwareModel, *ServerResponse, error) {
	hardwareModels := []*HardwareModel{}
	resp := ServerResponse{Records: &hardwareModels}

	if err := c.list(ctx, hardwareModelsEndpoint, nil, &resp); err != nil {
		return nil, nil, err
	}

	return hardwareModels, &resp, nil
}

// GetHardwareModel retrieves a hardware vendor model record
func (c *Client) GetHardwareModel(ctx context.Context, name string) (*HardwareModel, *ServerResponse, error) {
	hardwareModel := &HardwareModel{}
	resp := ServerResponse{Record: hardwareModel}

	endpoint := path.Join(hardwareModelsEndpoint, name)
	if err := c.get(ctx, endpoint, &resp); err != nil {
		return nil, nil, err
	}

	return hardwareModel, &resp, nil
}

// DeleteHardwareModel purges a hardware model record
func (c *Client) DeleteHardwareModel(ctx context.Context, name string) (*ServerResponse, error) {
	endpoint := path.Join(hardwareModelsEndpoint, name)
	return c.delete(ctx, endpoint)
}

// CreateServerBMC creates a server BMC record - requires the server relation
func (c *Client) CreateServerBMC(ctx context.Context, serverBMC *ServerBMC) (*ServerResponse, error) {
	endpoint := path.Join(serversEndpoint, serverBMC.ServerID.String(), serverBMCsEndpoint)
	return c.post(ctx, endpoint, serverBMC)
}

// GetServerBMC retrieves a server's BMC record
func (c *Client) GetServerBMC(ctx context.Context, serverID uuid.UUID) (*ServerBMC, *ServerResponse, error) {
	serverBMC := &ServerBMC{}
	resp := ServerResponse{Record: serverBMC}

	endpoint := path.Join(serversEndpoint, serverID.String(), serverBMCsEndpoint)
	if err := c.get(ctx, endpoint, &resp); err != nil {
		return nil, nil, err
	}

	return serverBMC, &resp, nil
}

// DeleteServerBMC purges a server's BMC record
func (c *Client) DeleteServerBMC(ctx context.Context, serverID uuid.UUID) (*ServerResponse, error) {
	endpoint := path.Join(serversEndpoint, serverID.String(), serverBMCsEndpoint)
	return c.delete(ctx, endpoint)
}

// SetInstalledFirmware creates a server component installed firmware record - requires the component relation
func (c *Client) SetInstalledFirmware(ctx context.Context, installedFirmware *InstalledFirmware) (*ServerResponse, error) {
	return c.post(ctx, installedFirmwareEndpoint, installedFirmware)
}

// ListInstalledFirmware lists server component firmware installed records
func (c *Client) ListInstalledFirmware(ctx context.Context) ([]*InstalledFirmware, *ServerResponse, error) {
	installedFirmware := []*InstalledFirmware{}
	resp := ServerResponse{Records: &installedFirmware}

	if err := c.list(ctx, installedFirmwareEndpoint, nil, &resp); err != nil {
		return nil, nil, err
	}

	return installedFirmware, &resp, nil
}

// GetInstalledFirmware retrieves a server component firmware installed record by the componentID
func (c *Client) GetInstalledFirmware(ctx context.Context, componentID uuid.UUID) (*InstalledFirmware, *ServerResponse, error) {
	installedFirmware := &InstalledFirmware{}
	resp := ServerResponse{Record: installedFirmware}

	endpoint := path.Join(installedFirmwareEndpoint, componentID.String())
	if err := c.get(ctx, endpoint, &resp); err != nil {
		return nil, nil, err
	}

	return installedFirmware, &resp, nil
}

// DeleteInstalledFirmware purges a installed firmware record (soft delete)
func (c *Client) DeleteInstalledFirmware(ctx context.Context, componentID uuid.UUID) (*ServerResponse, error) {
	endpoint := path.Join(installedFirmwareEndpoint, componentID.String())
	return c.delete(ctx, endpoint)
}

// SetComponentStatus creates or updates a component status record
func (c *Client) SetComponentStatus(ctx context.Context, componentStatus *ComponentStatus) (*ServerResponse, error) {
	return c.post(ctx, componentStatusEndpoint, componentStatus)
}

// ListComponentStatus lists all component status records
func (c *Client) ListComponentStatus(ctx context.Context) ([]*ComponentStatus, *ServerResponse, error) {
	componentStatus := []*ComponentStatus{}
	resp := ServerResponse{Records: &componentStatus}
	if err := c.list(ctx, componentStatusEndpoint, nil, &resp); err != nil {
		return nil, nil, err
	}
	return componentStatus, &resp, nil
}

// GetComponentStatus retrieves a component status by component ID
func (c *Client) GetComponentStatus(ctx context.Context, componentID uuid.UUID) (*ComponentStatus, *ServerResponse, error) {
	componentStatus := &ComponentStatus{}
	resp := ServerResponse{Record: componentStatus}
	endpoint := path.Join(componentStatusEndpoint, componentID.String())
	if err := c.get(ctx, endpoint, &resp); err != nil {
		return nil, nil, err
	}
	return componentStatus, &resp, nil
}

// DeleteComponentStatus deletes a component status record
func (c *Client) DeleteComponentStatus(ctx context.Context, componentID uuid.UUID) (*ServerResponse, error) {
	endpoint := path.Join(componentStatusEndpoint, componentID.String())
	return c.delete(ctx, endpoint)
}

// SetServerStatus creates or updates a server status record
func (c *Client) SetServerStatus(ctx context.Context, serverStatus *ServerStatus) (*ServerResponse, error) {
	return c.post(ctx, serverStatusEndpoint, serverStatus)
}

// ListServerStatus lists all server status records
func (c *Client) ListServerStatus(ctx context.Context) ([]*ServerStatus, *ServerResponse, error) {
	serverStatus := []*ServerStatus{}
	resp := ServerResponse{Records: &serverStatus}
	if err := c.list(ctx, serverStatusEndpoint, nil, &resp); err != nil {
		return nil, nil, err
	}
	return serverStatus, &resp, nil
}

// GetServerStatus retrieves a server status by server ID
func (c *Client) GetServerStatus(ctx context.Context, serverID uuid.UUID) (*ServerStatus, *ServerResponse, error) {
	serverStatus := &ServerStatus{}
	resp := ServerResponse{Record: serverStatus}
	endpoint := path.Join(serverStatusEndpoint, serverID.String())
	if err := c.get(ctx, endpoint, &resp); err != nil {
		return nil, nil, err
	}
	return serverStatus, &resp, nil
}

// DeleteServerStatus deletes a server status record
func (c *Client) DeleteServerStatus(ctx context.Context, serverID uuid.UUID) (*ServerResponse, error) {
	endpoint := path.Join(serverStatusEndpoint, serverID.String())
	return c.delete(ctx, endpoint)
}

// SetComponentCapability creates or updates a component capability record
func (c *Client) SetComponentCapability(ctx context.Context, capability []*ComponentCapability) (*ServerResponse, error) {
	return c.post(ctx, componentCapabilityEndpoint, capability)
}

// GetComponentCapability retrieves a component capability by component ID and capability name
func (c *Client) GetComponentCapability(ctx context.Context, componentID uuid.UUID) (*ComponentCapability, *ServerResponse, error) {
	capability := &ComponentCapability{}
	resp := ServerResponse{Record: capability}
	endpoint := path.Join(componentCapabilityEndpoint, componentID.String())
	if err := c.get(ctx, endpoint, &resp); err != nil {
		return nil, nil, err
	}
	return capability, &resp, nil
}

// DeleteComponentCapability deletes a component capability record
func (c *Client) DeleteComponentCapability(ctx context.Context, componentID uuid.UUID) (*ServerResponse, error) {
	endpoint := path.Join(componentCapabilityEndpoint, componentID.String())
	return c.delete(ctx, endpoint)
}

// SetComponentMetadata creates or updates a component metadata record
func (c *Client) SetComponentMetadata(ctx context.Context, metadata []*ComponentMetadata) (*ServerResponse, error) {
	return c.post(ctx, componentMetadataEndpoint, metadata)
}

// ListComponentMetadata lists all component metadata records
// componentID and namespace are optional filters
func (c *Client) ListComponentMetadata(ctx context.Context, componentID uuid.UUID, namespace string) ([]*ComponentMetadata, *ServerResponse, error) {
	endpoint := path.Join(componentMetadataEndpoint, componentID.String(), namespace)
	metadata := []*ComponentMetadata{}
	resp := ServerResponse{Records: &metadata}
	if err := c.list(ctx, endpoint, nil, &resp); err != nil {
		return nil, nil, err
	}
	return metadata, &resp, nil
}

// GetComponentMetadata retrieves a component metadata by component ID and namespace
func (c *Client) GetComponentMetadata(ctx context.Context, componentID uuid.UUID, namespace string) (*ComponentMetadata, *ServerResponse, error) {
	metadata := &ComponentMetadata{}
	resp := ServerResponse{Record: metadata}
	endpoint := path.Join(componentMetadataEndpoint, componentID.String(), namespace)
	if err := c.get(ctx, endpoint, &resp); err != nil {
		return nil, nil, err
	}
	return metadata, &resp, nil
}

// DeleteComponentMetadata deletes a component metadata record
func (c *Client) DeleteComponentMetadata(ctx context.Context, componentID uuid.UUID, namespace string) (*ServerResponse, error) {
	endpoint := path.Join(componentMetadataEndpoint, componentID.String(), namespace)
	return c.delete(ctx, endpoint)
}

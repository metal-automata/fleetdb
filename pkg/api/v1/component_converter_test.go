package fleetdbapi_test

import (
	"fmt"
	"slices"
	"strconv"
	"testing"

	"github.com/bmc-toolbox/common"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	fleetdbapi "github.com/metal-automata/fleetdb/pkg/api/v1"
	"github.com/metal-automata/fleetdb/pkg/api/v1/fixtures"
)

func componentSlugTypeMap() map[string]*fleetdbapi.ServerComponentType {
	m := make(map[string]*fleetdbapi.ServerComponentType, len(serverComponentTypes))
	for _, t := range serverComponentTypes {
		m[t.Slug] = t
	}

	return m
}

var (
	serverComponentTypes = fleetdbapi.ServerComponentTypeSlice{
		&fleetdbapi.ServerComponentType{
			ID:   "02dc2503-b64c-439b-9f25-8e130705f14a",
			Name: "Backplane-Expander",
			Slug: "backplane-expander",
		},
		&fleetdbapi.ServerComponentType{
			ID:   "1e0c3417-d63c-4fd5-88f7-4c525c70da12",
			Name: "Mainboard",
			Slug: "mainboard",
		},
		&fleetdbapi.ServerComponentType{
			ID:   "262e1a12-25a0-4d84-8c79-b3941603d48e",
			Name: "BIOS",
			Slug: "bios",
		},
		&fleetdbapi.ServerComponentType{
			ID:   "322b8728-dcc9-44e3-a139-81220c75a339",
			Name: "NVMe-PCIe-SSD",
			Slug: "nvme-pcie-ssd",
		},
		&fleetdbapi.ServerComponentType{
			ID:   "352631d2-b1ed-4d8e-85f7-9c92ddb76679",
			Name: "Sata-SSD",
			Slug: "sata-ssd",
		},
		&fleetdbapi.ServerComponentType{
			ID:   "3717d747-3cc3-4800-822c-4c7a9ac2c314",
			Name: "Drive",
			Slug: "drive",
		},
		&fleetdbapi.ServerComponentType{
			ID:   "3fc448ce-ea68-4f7c-beb1-c376f594db80",
			Name: "Chassis",
			Slug: "chassis",
		},
		&fleetdbapi.ServerComponentType{
			ID:   "4588a8fb-2e0f-4fa1-9634-9819a319b70b",
			Name: "GPU",
			Slug: "gpu",
		},
		&fleetdbapi.ServerComponentType{
			ID:   "5850ede2-d6d6-4df7-89d6-eab9110a9113",
			Name: "NIC",
			Slug: "nic",
		},
		&fleetdbapi.ServerComponentType{
			ID:   "5ac890cc-dd92-4609-9615-ca4b05b62a8e",
			Name: "PhysicalMemory",
			Slug: "physicalmemory",
		},
		&fleetdbapi.ServerComponentType{
			ID:   "75fc736e-fe42-4495-8e62-02d46fd08528",
			Name: "CPU",
			Slug: "cpu",
		},
		&fleetdbapi.ServerComponentType{
			ID:   "79ad53a2-0c05-4912-a156-8311bd54017d",
			Name: "TPM",
			Slug: "tpm",
		},
		&fleetdbapi.ServerComponentType{
			ID:   "9f5f39a4-82ed-4870-ab32-268bec45c8c8",
			Name: "Enclosure",
			Slug: "enclosure",
		},
		&fleetdbapi.ServerComponentType{
			ID:   "cbfbbe99-8d79-49e0-8f5d-c5296932bbd1",
			Name: "Sata-HDD",
			Slug: "sata-hdd",
		},
		&fleetdbapi.ServerComponentType{
			ID:   "ce396912-210e-4f6e-902d-9f07a8efe092",
			Name: "CPLD",
			Slug: "cpld",
		},
		&fleetdbapi.ServerComponentType{
			ID:   "d51b438b-a767-459e-8eda-fd0700a46686",
			Name: "Power-Supply",
			Slug: "power-supply",
		},
		&fleetdbapi.ServerComponentType{
			ID:   "e96c8557-4a71-4887-a3bb-28b6f90e5489",
			Name: "BMC",
			Slug: "bmc",
		},
		&fleetdbapi.ServerComponentType{
			ID:   "eb82dbe3-df77-4409-833b-c44241885410",
			Name: "unknown",
			Slug: "unknown",
		},
		&fleetdbapi.ServerComponentType{
			ID:   "ef563926-8011-4985-bc4a-7ed7e9933971",
			Name: "StorageController",
			Slug: "storagecontroller",
		},
		&fleetdbapi.ServerComponentType{
			ID:   "df563926-8011-4985-bc4a-7ed7e9933972",
			Name: "NICPort",
			Slug: "nicport",
		},
	}
)

func capabilitiesForComponent(t *testing.T, dev *common.Device, component, serial string) []*common.Capability {
	t.Helper()

	switch component {
	case "bios":
		return dev.BIOS.Capabilities
	case "cpu":
		for idx, cpu := range dev.CPUs {
			if cpu.Serial == serial || strconv.Itoa(idx) == serial {
				return cpu.Capabilities
			}
		}
	case "nic":
		for idx, nic := range dev.NICs {
			if nic.Serial == serial || strconv.Itoa(idx) == serial {
				return nic.Capabilities
			}
		}
	case "drive":
		for _, drive := range dev.Drives {
			if drive.Serial == serial {
				return drive.Capabilities
			}
		}

	default:
		return nil
	}

	return nil
}

func TestFromCommonDevice(t *testing.T) {
	serverID := uuid.New()
	testcases := []struct {
		name              string
		dev               *common.Device
		counts            map[string]int // map of slug to component counts
		components        int
		drives            int
		mem               int
		psu               int
		storageController int
		vendors           []string
		firmwareEmpty     []string
		metadataEmpty     []string
	}{
		{
			name:       "device 1",
			dev:        fixtures.CopyDevice(fixtures.DellR6515),
			components: 23,
			counts: map[string]int{
				"bmc":               1,
				"mainboard":         1,
				"power-supply":      2,
				"cpu":               1,
				"storagecontroller": 4,
				"bios":              1,
				"physicalmemory":    8,
				"nic":               1,
				"drive":             4,
			},
			vendors: []string{
				"intel",
				"micron",
				"amd",
				"broadcom",
				"marvell",
				"dell",
				"hynix",
			},
			firmwareEmpty: []string{
				"bmc",
				"mainboard",
				"physicalmemory",
				"power-supply",
				"storagecontroller",
			},
			metadataEmpty: []string{
				"bmc",
				"mainboard",
			},
		},
		{
			name:       "device 2",
			dev:        fixtures.CopyDevice(fixtures.SMCX11DPH),
			components: 41,
			counts: map[string]int{
				"bmc":               1,
				"mainboard":         1,
				"nicport":           1,
				"power-supply":      2,
				"cpu":               2,
				"cpld":              1,
				"storagecontroller": 1,
				"bios":              1,
				"physicalmemory":    12,
				"nic":               3,
				"drive":             16,
			},
			vendors: []string{
				"supermicro",
				"hynix",
				"intel",
				"mellanox",
				"micron",
				"toshiba",
				"hgst",
				"broadcom",
			},
			firmwareEmpty: []string{
				"mainboard",
				"physicalmemory",
				"nicport",
				"power-supply",
				"storagecontroller",
			},
			metadataEmpty: []string{
				"bmc",
				"mainboard",
				"cpld",
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			conv := fleetdbapi.NewComponentConverter(fleetdbapi.Inband, componentSlugTypeMap())
			server, err := conv.FromCommonDevice(serverID, fixtures.CopyDevice(tc.dev))
			assert.Nil(t, err)

			//b, _ := json.MarshalIndent(server, "  ", "  ")
			//fmt.Println(string(b))

			assert.NotEmpty(t, server.Model)
			assert.NotEmpty(t, server.Vendor)
			assert.NotEmpty(t, server.Serial)
			assert.Equal(t, server.Model, common.FormatProductName(server.Model))
			assert.Equal(t, server.Vendor, common.FormatVendorName(server.Vendor))
			assert.Len(t, server.Components, tc.components)

			gotVendors := []string{}

			// map of component slug to counts
			gotCounts := map[string]int{}

			for _, component := range server.Components {
				// expect vendor attribute
				assert.NotEmpty(t, component.Vendor)

				// expect model attribute
				assert.NotEmpty(t, component.Model)

				// expect model, vendor formatted
				assert.Equal(t, component.Model, common.FormatProductName(component.Model))
				assert.Equal(t, component.Vendor, common.FormatVendorName(component.Vendor))

				// expect components not in the installedFirmwareEmpty list to have firmware versions
				if component.InstalledFirmware == nil {
					assert.Contains(t, tc.firmwareEmpty, component.Name)
				}

				// expect components to have metadata set except for the ones in the metadataEmpty slice
				if len(component.Metadata) == 0 {
					assert.Contains(t, tc.metadataEmpty, component.Name)
				} else {
					assert.Equal(t, component.Metadata[0].Namespace, fleetdbapi.ComponentMetadataGenericNS)
					assert.GreaterOrEqual(t, len(component.Metadata[0].Data), 10)
				}

				if !slices.Contains(gotVendors, component.Vendor) {
					gotVendors = append(gotVendors, component.Vendor)
				}

				gotCounts[component.Name]++

				// expect capabilities are populated
				componentCaps := capabilitiesForComponent(t, tc.dev, component.Name, component.Serial)
				assert.Equal(t, len(componentCaps), len(component.Capabilities), fmt.Sprintf("caps not equal %s, serial %s", component.Name, component.Serial))
				for _, cap := range component.Capabilities {
					assert.NotEmpty(t, cap.Name)
				}
			}

			assert.ElementsMatch(t, tc.vendors, gotVendors)
			assert.Equal(t, tc.counts, gotCounts)
		})
	}

}

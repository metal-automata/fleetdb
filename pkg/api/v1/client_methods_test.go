package fleetdbapi_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	fleetdbapi "github.com/metal-automata/fleetdb/pkg/api/v1"
)

func TestFleetdbCreate(t *testing.T) {
	mockClientTests(t, func(ctx context.Context, respCode int, expectError bool) error {
		srv := fleetdbapi.Server{UUID: uuid.New(), FacilityCode: "Test1"}
		jsonResponse := json.RawMessage([]byte(`{"message": "resource created", "slug":"00000000-0000-0000-0000-000000001234"}`))

		c := mockClient(string(jsonResponse), respCode)
		res, _, err := c.Create(ctx, srv)

		if !expectError {
			assert.Equal(t, "00000000-0000-0000-0000-000000001234", res.String())
		}

		return err
	})
}

func TestFleetdbDelete(t *testing.T) {
	mockClientTests(t, func(ctx context.Context, respCode int, _ bool) error {
		jsonResponse := json.RawMessage([]byte(`{"message": "resource deleted"}`))
		c := mockClient(string(jsonResponse), respCode)
		_, err := c.Delete(ctx, fleetdbapi.Server{UUID: uuid.New()})

		return err
	})
}

func TestFleetdbGet(t *testing.T) {
	mockClientTests(t, func(ctx context.Context, respCode int, expectError bool) error {
		srv := fleetdbapi.Server{UUID: uuid.New(), FacilityCode: "Test1"}
		jsonResponse, err := json.Marshal(fleetdbapi.ServerResponse{Record: srv})
		require.Nil(t, err)

		c := mockClient(string(jsonResponse), respCode)
		res, _, err := c.GetServer(ctx, srv.UUID, nil)

		if !expectError {
			assert.Equal(t, srv.UUID, res.UUID)
			assert.Equal(t, srv.FacilityCode, res.FacilityCode)
		}

		return err
	})
}

// TODO: for when list is implemented
//func TestFleetdbList(t *testing.T) {
//	mockClientTests(t, func(ctx context.Context, respCode int, expectError bool) error {
//		srv := []fleetdbapi.Server{{UUID: uuid.New(), FacilityCode: "Test1"}}
//		jsonResponse, err := json.Marshal(fleetdbapi.ServerResponse{Records: srv})
//		require.Nil(t, err)
//
//		c := mockClient(string(jsonResponse), respCode)
//		res, _, err := c.List(ctx, nil)
//
//		if !expectError {
//			assert.ElementsMatch(t, srv, res)
//		}
//
//		return err
//	})
//}

func TestFleetdbUpdate(t *testing.T) {
	mockClientTests(t, func(ctx context.Context, respCode int, _ bool) error {
		jsonResponse, err := json.Marshal(fleetdbapi.ServerResponse{Message: "resource updated"})
		require.Nil(t, err)

		c := mockClient(string(jsonResponse), respCode)
		_, err = c.Update(ctx, uuid.UUID{}, fleetdbapi.Server{Name: "new-name"})

		return err
	})
}

func TestFleetdbComponentsGet(t *testing.T) {
	mockClientTests(t, func(ctx context.Context, respCode int, expectError bool) error {
		sc := []*fleetdbapi.ServerComponent{{Name: "unit-test", Serial: "1234"}}
		jsonResponse, err := json.Marshal(fleetdbapi.ServerResponse{Records: sc})
		require.Nil(t, err)

		c := mockClient(string(jsonResponse), respCode)
		res, _, err := c.GetComponents(ctx, uuid.UUID{}, nil)

		if !expectError {
			assert.ElementsMatch(t, sc, res)
		}

		return err
	})
}

func TestFleetdbComponentsList(t *testing.T) {
	mockClientTests(t, func(ctx context.Context, respCode int, expectError bool) error {
		sc := []*fleetdbapi.ServerComponent{{Name: "unit-test", Serial: "1234"}}
		jsonResponse, err := json.Marshal(fleetdbapi.ServerResponse{Records: sc})
		require.Nil(t, err)

		c := mockClient(string(jsonResponse), respCode)
		res, _, err := c.ListComponents(ctx, &fleetdbapi.ServerComponentListParams{Name: "unit-test", Serial: "1234"})

		if !expectError {
			assert.ElementsMatch(t, sc, res)
		}

		return err
	})
}

func TestFleetdbComponentsCreate(t *testing.T) {
	mockClientTests(t, func(ctx context.Context, respCode int, expectError bool) error {
		jsonResponse, err := json.Marshal(fleetdbapi.ServerResponse{Message: "resource created"})
		require.Nil(t, err)

		c := mockClient(string(jsonResponse), respCode)
		res, err := c.InitComponentCollection(ctx, uuid.New(), fleetdbapi.ServerComponentSlice{{Name: "unit-test"}}, fleetdbapi.Inband)

		if !expectError {
			assert.Contains(t, res.Message, "resource created")
		}

		return err
	})
}

func TestFleetdbComponentsUpdate(t *testing.T) {
	testUUID := uuid.New()
	testComponents := fleetdbapi.ServerComponentSlice{
		{
			Name:       "unit-test",
			Model:      "TestModel",
			Serial:     "TestSerial",
			ServerUUID: testUUID,
		},
	}

	mockClientTests(t, func(ctx context.Context, respCode int, expectError bool) error {
		jsonResponse, err := json.Marshal(fleetdbapi.ServerResponse{Message: "component(s) updated"})
		require.Nil(t, err)

		c := mockClient(string(jsonResponse), respCode)
		res, err := c.UpdateComponentCollection(ctx, testUUID, testComponents, fleetdbapi.Inband)
		if !expectError {
			assert.Contains(t, res.Message, "component(s) updated")
		}
		return err
	})
}

func TestFleetdbCreateServerComponentFirmware(t *testing.T) {
	mockClientTests(t, func(ctx context.Context, respCode int, expectError bool) error {
		firmware := fleetdbapi.ComponentFirmwareVersion{
			UUID:    uuid.New(),
			Vendor:  "Dell",
			Model:   []string{"R615"},
			Version: "21.07.00",
		}
		jsonResponse := json.RawMessage([]byte(`{"message": "resource created", "slug":"00000000-0000-0000-0000-000000001234"}`))

		c := mockClient(string(jsonResponse), respCode)
		res, _, err := c.CreateServerComponentFirmware(ctx, firmware)

		if !expectError {
			assert.Equal(t, "00000000-0000-0000-0000-000000001234", res.String())
		}

		return err
	})
}

func TestFleetdbServerComponentFirmwareDelete(t *testing.T) {
	mockClientTests(t, func(ctx context.Context, respCode int, _ bool) error {
		jsonResponse := json.RawMessage([]byte(`{"message": "resource deleted"}`))
		c := mockClient(string(jsonResponse), respCode)
		_, err := c.DeleteServerComponentFirmware(ctx, fleetdbapi.ComponentFirmwareVersion{UUID: uuid.New()})

		return err
	})
}
func TestFleetdbServerComponentFirmwareGet(t *testing.T) {
	mockClientTests(t, func(ctx context.Context, respCode int, expectError bool) error {
		firmware := fleetdbapi.ComponentFirmwareVersion{
			UUID:    uuid.New(),
			Vendor:  "Dell",
			Model:   []string{"R615"},
			Version: "21.07.00",
		}
		jsonResponse, err := json.Marshal(fleetdbapi.ServerResponse{Record: firmware})
		require.Nil(t, err)

		c := mockClient(string(jsonResponse), respCode)
		res, _, err := c.GetServerComponentFirmware(ctx, firmware.UUID)

		if !expectError {
			assert.Equal(t, firmware.UUID, res.UUID)
			assert.Equal(t, firmware.Vendor, res.Vendor)
			assert.Equal(t, firmware.Model, res.Model)
			assert.Equal(t, firmware.Version, res.Version)
		}

		return err
	})
}

func TestFleetdbServerComponentFirmwareList(t *testing.T) {
	mockClientTests(t, func(ctx context.Context, respCode int, expectError bool) error {
		firmware := []fleetdbapi.ComponentFirmwareVersion{{
			UUID:    uuid.New(),
			Vendor:  "Dell",
			Model:   []string{"R615"},
			Version: "21.07.00",
		}}
		jsonResponse, err := json.Marshal(fleetdbapi.ServerResponse{Records: firmware})
		require.Nil(t, err)

		c := mockClient(string(jsonResponse), respCode)
		res, _, err := c.ListServerComponentFirmware(ctx, nil)

		if !expectError {
			assert.ElementsMatch(t, firmware, res)
		}

		return err
	})
}

func TestFleetdbServerComponentFirmwareUpdate(t *testing.T) {
	mockClientTests(t, func(ctx context.Context, respCode int, _ bool) error {
		jsonResponse, err := json.Marshal(fleetdbapi.ServerResponse{Message: "resource updated"})
		require.Nil(t, err)

		c := mockClient(string(jsonResponse), respCode)
		_, err = c.UpdateServerComponentFirmware(ctx, uuid.UUID{}, fleetdbapi.ComponentFirmwareVersion{UUID: uuid.New()})

		return err
	})
}

func TestBillOfMaterialsBatchUpload(t *testing.T) {
	mockClientTests(t, func(ctx context.Context, respCode int, expectError bool) error {
		bom := []fleetdbapi.Bom{{SerialNum: "fakeSerialNum1", AocMacAddress: "fakeAocMacAddress1", BmcMacAddress: "fakeBmcMacAddress1"}}
		jsonResponse, err := json.Marshal(fleetdbapi.ServerResponse{Record: bom})
		require.Nil(t, err)

		c := mockClient(string(jsonResponse), respCode)
		res, err := c.BillOfMaterialsBatchUpload(ctx, bom)

		if !expectError {
			assert.Equal(t, []interface{}([]interface{}{
				map[string]interface{}{
					"aoc_mac_address": "fakeAocMacAddress1",
					"bmc_mac_address": "fakeBmcMacAddress1",
					"metro":           "",
					"num_def_pwd":     "",
					"num_defi_pmi":    "",
					"serial_num":      "fakeSerialNum1"}}), res.Record)
		}

		return err
	})
}

func TestGetBomInfoByAOCMacAddr(t *testing.T) {
	mockClientTests(t, func(ctx context.Context, respCode int, expectError bool) error {
		bom := fleetdbapi.Bom{SerialNum: "fakeSerialNum1", AocMacAddress: "fakeAocMacAddress1", BmcMacAddress: "fakeBmcMacAddress"}
		jsonResponse, err := json.Marshal(fleetdbapi.ServerResponse{Record: bom})
		require.Nil(t, err)

		c := mockClient(string(jsonResponse), respCode)
		respBom, _, err := c.GetBomInfoByAOCMacAddr(ctx, "fakeAocMacAddress1")

		if !expectError {
			assert.Equal(t, &bom, respBom)
		}

		return err
	})
}

func TestGetBomInfoByBMCMacAddr(t *testing.T) {
	mockClientTests(t, func(ctx context.Context, respCode int, expectError bool) error {
		bom := fleetdbapi.Bom{SerialNum: "fakeSerialNum1", AocMacAddress: "fakeAocMacAddress1", BmcMacAddress: "fakeBmcMacAddress1"}
		jsonResponse, err := json.Marshal(fleetdbapi.ServerResponse{Record: bom})
		require.Nil(t, err)

		c := mockClient(string(jsonResponse), respCode)
		respBom, _, err := c.GetBomInfoByBMCMacAddr(ctx, "fakeBmcMacAddress1")

		if !expectError {
			assert.Equal(t, &bom, respBom)
		}

		return err
	})
}

package fleetdbapi_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/metal-automata/fleetdb/internal/dbtools"
	fleetdbapi "github.com/metal-automata/fleetdb/pkg/api/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestIntegrationHardwareModelCreate(t *testing.T) {
	s := serverTest(t)

	scopes := []string{"create:hardware-models"}
	scopedRealClientTests(t, scopes, func(ctx context.Context, authToken string, respCode int, expectError bool) error {
		s.Client.SetToken(authToken)

		hwm := &fleetdbapi.HardwareModel{Name: "foo123", HardwareVendorName: dbtools.FixtureHardwareVendorNameBar}
		resp, err := s.Client.CreateHardwareModel(ctx, hwm)

		if !expectError {
			require.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, hwm.Name, resp.Slug)
			assert.NotNil(t, resp.Links.Self)
			assert.Equal(t, fmt.Sprintf("http://test.hollow.com/api/v1/hardware-models/%s", resp.Slug), resp.Links.Self.Href)
		}

		return err
	})
}

func TestIntegrationHardwareModelList(t *testing.T) {
	s := serverTest(t)

	scopes := []string{"read:hardware-models"}
	scopedRealClientTests(t, scopes, func(ctx context.Context, authToken string, respCode int, expectError bool) error {
		s.Client.SetToken(authToken)

		boil.DebugMode = true
		hardwareModels, resp, err := s.Client.ListHardwareModels(ctx)

		boil.DebugMode = false
		if !expectError {
			require.NoError(t, err)
			assert.Len(t, hardwareModels, 2)
			assert.EqualValues(t, 2, resp.PageCount)
			assert.EqualValues(t, 1, resp.TotalPages)
			assert.EqualValues(t, 2, resp.TotalRecordCount)
			// We returned everything, so we shouldnt have a next page info
			assert.Nil(t, resp.Links.Next)
			assert.Nil(t, resp.Links.Previous)
		}

		return err
	})
}

func TestIntegrationHardwareModelGet(t *testing.T) {
	s := serverTest(t)

	scopes := []string{"read:hardware-models"}
	scopedRealClientTests(t, scopes, func(ctx context.Context, authToken string, _ int, expectError bool) error {
		s.Client.SetToken(authToken)
		hardwareModel, resp, err := s.Client.GetHardwareModel(ctx, dbtools.FixtureHardwareModelBaz123Name)
		if !expectError {
			require.NoError(t, err)
			assert.NotNil(t, resp.Record)
			assert.NotNil(t, hardwareModel)

			hwm, ok := resp.Record.(*fleetdbapi.HardwareModel)
			assert.True(t, ok)

			assert.Equal(t, dbtools.FixtureHardwareModelBaz123Name, hardwareModel.Name)
			assert.NotNil(t, hwm.ID, dbtools.FixtureHardwareModelBar123)
			// We returned everything, so we shouldnt have a next page info
			assert.Nil(t, resp.Links.Next)
			assert.Nil(t, resp.Links.Previous)
		}

		return err
	})

	_, _, err := s.Client.GetHardwareModel(context.Background(), "non-existant")
	assert.ErrorContainsf(t, err, "404", "")
}

func TestIntegrationHardwareModelDelete(t *testing.T) {
	s := serverTest(t)

	scopes := []string{"delete:hardware-models", "read:hardware-models"}
	scopedRealClientTests(t, scopes, func(ctx context.Context, authToken string, _ int, expectError bool) error {
		s.Client.SetToken(authToken)

		resp, err := s.Client.DeleteHardwareModel(ctx, dbtools.FixtureHardwareModelBaz123Name)
		if !expectError {
			require.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, "resource deleted", resp.Message)

		}
		return err
	})

	_, _, err := s.Client.GetHardwareModel(context.Background(), dbtools.FixtureHardwareModelBaz123Name)
	assert.ErrorContainsf(t, err, "404", "")
}

package fleetdbapi_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/metal-automata/fleetdb/internal/dbtools"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	fleetdbapi "github.com/metal-automata/fleetdb/pkg/api/v1"
)

func TestIntegrationInstalledFirmware_Set(t *testing.T) {
	s := serverTest(t)

	scopes := []string{"create:installed-firmware"}
	scopedRealClientTests(t, scopes, func(ctx context.Context, authToken string, respCode int, expectError bool) error {
		s.Client.SetToken(authToken)
		installedFirmware := &fleetdbapi.InstalledFirmware{
			ID:                uuid.New(),
			ServerComponentID: uuid.MustParse(dbtools.FixtureMarlinLeftFin.ID),
			Version:           "1.0",
		}

		resp, err := s.Client.SetInstalledFirmware(ctx, installedFirmware)
		if !expectError {
			require.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, "resource created", resp.Message)
			assert.NotNil(t, resp.Links.Self)
			assert.Equal(t, fmt.Sprintf("http://test.hollow.com/api/v1/installed-firmware/%s", resp.Slug), resp.Links.Self.Href)
		}

		return err
	})

	testcases := []struct {
		name             string
		payload          *fleetdbapi.InstalledFirmware
		expectError      string
		expectStatusCode string
	}{
		{
			name: "Unique firmware version required",
			payload: &fleetdbapi.InstalledFirmware{
				ServerComponentID: uuid.MustParse(dbtools.FixtureNemoLeftFin.ID),
				Version:           "1.0", // attempt to set the fixture version to the same
			},
			expectError:      "duplicate key",
			expectStatusCode: "500",
		},
		{
			name: "Update record",
			payload: &fleetdbapi.InstalledFirmware{
				ServerComponentID: uuid.MustParse(dbtools.FixtureNemoLeftFin.ID),
				Version:           "3.0",
			},
			expectError:      "",
			expectStatusCode: "200",
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			s.Client.SetToken(validToken(scopes))
			resp, err := s.Client.SetInstalledFirmware(context.TODO(), *&tt.payload)
			if tt.expectError != "" {
				assert.Error(t, err)
				assert.ErrorContains(t, err, tt.expectError)
				assert.ErrorContains(t, err, tt.expectStatusCode)

				return
			}

			assert.Nil(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, "resource created", resp.Message)

			// validate firmware record was updated
			s.Client.SetToken(validToken([]string{"read:installed-firmware"}))
			installedFirmware, _, err := s.Client.GetInstalledFirmware(context.TODO(), uuid.MustParse(dbtools.FixtureNemoLeftFin.ID))
			assert.Nil(t, err)
			assert.Equal(t, tt.payload.Version, installedFirmware.Version)
		})
	}

}

func TestIntegrationInstalledFirmwareList(t *testing.T) {
	s := serverTest(t)

	scopes := []string{"read:installed-firmware"}
	scopedRealClientTests(t, scopes, func(ctx context.Context, authToken string, respCode int, expectError bool) error {
		s.Client.SetToken(authToken)

		expectCount := len(dbtools.FixtureInstalledFirmwares)
		got, resp, err := s.Client.ListInstalledFirmware(ctx)
		if !expectError {
			require.NoError(t, err)
			assert.Len(t, got, expectCount)
			assert.EqualValues(t, expectCount, resp.PageCount)
			assert.EqualValues(t, 1, resp.TotalPages)
			assert.EqualValues(t, expectCount, resp.TotalRecordCount)
			// We returned everything, so we shouldnt have a next page info
			assert.Nil(t, resp.Links.Next)
			assert.Nil(t, resp.Links.Previous)
		}

		return err
	})
}

func TestIntegrationInstalledFirmwareGet(t *testing.T) {
	s := serverTest(t)
	componentID := uuid.MustParse(dbtools.FixtureNemoLeftFin.ID)

	scopes := []string{"read:installed-firmware"}
	scopedRealClientTests(t, scopes, func(ctx context.Context, authToken string, _ int, expectError bool) error {
		s.Client.SetToken(authToken)
		got, resp, err := s.Client.GetInstalledFirmware(ctx, componentID)
		if !expectError {
			require.NoError(t, err)
			assert.NotNil(t, resp.Record)
			assert.NotNil(t, got)
			_, ok := resp.Record.(*fleetdbapi.InstalledFirmware)
			assert.True(t, ok)
			assert.Equal(t, componentID, got.ServerComponentID)
			// We returned everything, so we shouldnt have a next page info
			assert.Nil(t, resp.Links.Next)
			assert.Nil(t, resp.Links.Previous)
		}

		return err
	})

	_, _, err := s.Client.GetInstalledFirmware(context.Background(), uuid.New())
	assert.ErrorContainsf(t, err, "404", "")
}

func TestIntegrationInstalledFirmwareDelete(t *testing.T) {
	s := serverTest(t)
	scopes := []string{"delete:installed-firmware"}
	scopedRealClientTests(t, scopes, func(ctx context.Context, authToken string, _ int, expectError bool) error {
		s.Client.SetToken(authToken)

		resp, err := s.Client.DeleteInstalledFirmware(ctx, uuid.MustParse(dbtools.FixtureNemoLeftFin.ID))
		if !expectError {
			require.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, "resource deleted", resp.Message)

		}
		return err
	})

	// validate firmware record was deleted
	s.Client.SetToken(validToken([]string{"read:installed-firmware"}))
	_, _, err := s.Client.GetInstalledFirmware(context.Background(), uuid.MustParse(dbtools.FixtureNemoLeftFin.ID))
	assert.ErrorContainsf(t, err, "404", "")
}

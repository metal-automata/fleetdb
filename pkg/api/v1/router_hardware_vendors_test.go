package fleetdbapi_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/metal-automata/fleetdb/internal/dbtools"
	fleetdbapi "github.com/metal-automata/fleetdb/pkg/api/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntegrationHardwareVendorCreate(t *testing.T) {
	s := serverTest(t)

	scopes := []string{"create:hardware-vendors"}
	scopedRealClientTests(t, scopes, func(ctx context.Context, authToken string, respCode int, expectError bool) error {
		s.Client.SetToken(authToken)

		hwv := &fleetdbapi.HardwareVendor{Name: "foo"}
		resp, err := s.Client.CreateHardwareVendor(ctx, hwv)
		if !expectError {
			require.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, hwv.Name, resp.Slug)
			assert.NotNil(t, resp.Links.Self)
			assert.Equal(t, fmt.Sprintf("http://test.hollow.com/api/v1/hardware-vendors/%s", resp.Slug), resp.Links.Self.Href)
		}

		return err
	})
}

func TestIntegrationHardwareVendorList(t *testing.T) {
	s := serverTest(t)

	scopes := []string{"read:hardware-vendors"}
	scopedRealClientTests(t, scopes, func(ctx context.Context, authToken string, _ int, expectError bool) error {
		s.Client.SetToken(authToken)
		hardwareVendors, resp, err := s.Client.ListHardwareVendors(ctx)
		if !expectError {
			require.NoError(t, err)
			assert.Len(t, hardwareVendors, 2)
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

func TestIntegrationHardwareVendorGet(t *testing.T) {
	s := serverTest(t)

	scopes := []string{"read:hardware-vendors"}
	scopedRealClientTests(t, scopes, func(ctx context.Context, authToken string, _ int, expectError bool) error {
		s.Client.SetToken(authToken)

		hardwareVendor, resp, err := s.Client.GetHardwareVendor(ctx, dbtools.FixtureHardwareVendorNameBar)
		if !expectError {
			require.NoError(t, err)
			assert.NotNil(t, resp.Record)

			hwv, ok := resp.Record.(*fleetdbapi.HardwareVendor)
			assert.True(t, ok)

			assert.Equal(t, dbtools.FixtureHardwareVendorNameBar, hardwareVendor.Name)
			assert.NotNil(t, hwv.ID, dbtools.FixtureHardwareVendorNameBar)
			// We returned everything, so we shouldnt have a next page info
			assert.Nil(t, resp.Links.Next)
			assert.Nil(t, resp.Links.Previous)
		}

		return err
	})
}

func TestIntegrationHardwareVendorDelete(t *testing.T) {
	s := serverTest(t)

	scopes := []string{"delete:hardware-vendors", "read:hardware-vendors"}
	scopedRealClientTests(t, scopes, func(ctx context.Context, authToken string, _ int, expectError bool) error {
		s.Client.SetToken(authToken)

		resp, err := s.Client.DeleteHardwareVendor(ctx, dbtools.FixtureHardwareVendorNameBar)
		if !expectError {
			require.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, "resource deleted", resp.Message)

		}
		return err
	})

	_, _, err := s.Client.GetHardwareVendor(context.Background(), dbtools.FixtureHardwareVendorNameBar)
	assert.ErrorContainsf(t, err, "404", "")
}

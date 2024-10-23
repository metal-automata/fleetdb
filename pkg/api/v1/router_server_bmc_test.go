package fleetdbapi_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/metal-automata/fleetdb/internal/dbtools"
	fleetdbapi "github.com/metal-automata/fleetdb/pkg/api/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntegrationServerBMCCreate(t *testing.T) {
	s := serverTest(t)

	scopes := []string{"create:server-bmcs"}
	scopedRealClientTests(t, scopes, func(ctx context.Context, authToken string, respCode int, expectError bool) error {
		s.Client.SetToken(authToken)
		serverBMC := &fleetdbapi.ServerBMC{
			ID:                 uuid.New(),
			ServerID:           uuid.MustParse(dbtools.FixtureMarlin.ID),
			HardwareVendorName: dbtools.FixtureHardwareVendorNameBar,
			HardwareModelName:  dbtools.FixtureHardwareModelBar456Name,
			Username:           "user",
			IPAddress:          "127.0.0.1",
			MacAddress:         "de:ad:be:ef:ca:fe",
		}

		resp, err := s.Client.CreateServerBMC(ctx, serverBMC)
		if !expectError {
			require.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, "resource created", resp.Message)
			assert.NotNil(t, resp.Links.Self)
			assert.Equal(t, fmt.Sprintf("http://test.hollow.com/api/v1/server-bmcs/%s", resp.Slug), resp.Links.Self.Href)
		}

		return err
	})
}

func TestIntegrationServerBMCList(t *testing.T) {
	s := serverTest(t)

	scopes := []string{"read:server-bmcs"}
	scopedRealClientTests(t, scopes, func(ctx context.Context, authToken string, respCode int, expectError bool) error {
		s.Client.SetToken(authToken)

		expectCount := len(dbtools.FixtureServerBMCs)
		serverBMCs, resp, err := s.Client.ListServerBMCs(ctx)
		if !expectError {
			require.NoError(t, err)
			assert.Len(t, serverBMCs, expectCount)
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

func TestIntegrationServerBMCGet(t *testing.T) {
	s := serverTest(t)
	serverID := uuid.MustParse(dbtools.FixtureNemo.ID)
	scopes := []string{"read:server-bmcs"}
	scopedRealClientTests(t, scopes, func(ctx context.Context, authToken string, _ int, expectError bool) error {
		s.Client.SetToken(authToken)
		serverBMC, resp, err := s.Client.GetServerBMC(ctx, serverID)
		if !expectError {
			require.NoError(t, err)
			assert.NotNil(t, resp.Record)
			assert.NotNil(t, serverBMC)
			_, ok := resp.Record.(*fleetdbapi.ServerBMC)
			assert.True(t, ok)
			assert.Equal(t, serverID, serverBMC.ServerID)
			// We returned everything, so we shouldnt have a next page info
			assert.Nil(t, resp.Links.Next)
			assert.Nil(t, resp.Links.Previous)
		}

		return err
	})

	_, _, err := s.Client.GetServerBMC(context.Background(), uuid.New())
	assert.ErrorContainsf(t, err, "404", "")
}

func TestIntegrationServerBMCDelete(t *testing.T) {
	s := serverTest(t)
	serverID := uuid.MustParse(dbtools.FixtureNemo.ID)
	scopes := []string{"delete:server-bmcs", "read:server-bmcs"}
	scopedRealClientTests(t, scopes, func(ctx context.Context, authToken string, _ int, expectError bool) error {
		s.Client.SetToken(authToken)

		resp, err := s.Client.DeleteServerBMC(ctx, serverID)
		if !expectError {
			require.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, "resource deleted", resp.Message)

		}
		return err
	})

	_, _, err := s.Client.GetServerBMC(context.Background(), serverID)
	assert.ErrorContainsf(t, err, "404", "")
}

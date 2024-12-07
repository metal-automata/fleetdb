package fleetdbapi_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/metal-automata/fleetdb/internal/dbtools"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	fleetdbapi "github.com/metal-automata/fleetdb/pkg/api/v1"
)

func TestIntegrationServerStatusSet(t *testing.T) {
	s := serverTest(t)

	scopes := []string{"create:server-status"}
	scopedRealClientTests(t, scopes, func(ctx context.Context, authToken string, respCode int, expectError bool) error {
		s.Client.SetToken(authToken)
		serverStatus := &fleetdbapi.ServerStatus{
			ID:       uuid.New(),
			ServerID: uuid.MustParse(dbtools.FixtureMarlin.ID),
			Health:   "healthy",
			State:    "running",
			Info:     "all systems nominal",
		}

		resp, err := s.Client.SetServerStatus(ctx, serverStatus)
		if !expectError {
			require.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, "resource created", resp.Message)
		}

		return err
	})
}

func TestIntegrationServerStatusList(t *testing.T) {
	s := serverTest(t)

	scopes := []string{"read:server-status"}
	scopedRealClientTests(t, scopes, func(ctx context.Context, authToken string, respCode int, expectError bool) error {
		s.Client.SetToken(authToken)

		expectCount := len(dbtools.FixtureServerStatuses)
		got, resp, err := s.Client.ListServerStatus(ctx)
		if !expectError {
			require.NoError(t, err)
			assert.Len(t, got, expectCount)
			assert.EqualValues(t, expectCount, resp.PageCount)
			assert.EqualValues(t, expectCount, resp.TotalRecordCount)
			assert.Nil(t, resp.Links.Next)
			assert.Nil(t, resp.Links.Previous)
		}

		return err
	})
}

func TestIntegrationServerStatusGet(t *testing.T) {
	s := serverTest(t)
	serverID := uuid.MustParse(dbtools.FixtureNemo.ID)

	scopes := []string{"read:server-status"}
	scopedRealClientTests(t, scopes, func(ctx context.Context, authToken string, _ int, expectError bool) error {
		s.Client.SetToken(authToken)
		got, resp, err := s.Client.GetServerStatus(ctx, serverID)
		if !expectError {
			require.NoError(t, err)
			assert.NotNil(t, resp.Record)
			assert.NotNil(t, got)
			assert.Equal(t, serverID, got.ServerID)
		}

		return err
	})

	_, _, err := s.Client.GetServerStatus(context.Background(), uuid.New())
	assert.ErrorContains(t, err, "404")
}

func TestIntegrationServerStatusDelete(t *testing.T) {
	s := serverTest(t)
	scopes := []string{"delete:server-status"}
	scopedRealClientTests(t, scopes, func(ctx context.Context, authToken string, _ int, expectError bool) error {
		s.Client.SetToken(authToken)

		resp, err := s.Client.DeleteServerStatus(ctx, uuid.MustParse(dbtools.FixtureNemo.ID))
		if !expectError {
			require.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, "resource deleted", resp.Message)
		}
		return err
	})

	s.Client.SetToken(validToken([]string{"read:server-status"}))
	_, _, err := s.Client.GetServerStatus(context.Background(), uuid.MustParse(dbtools.FixtureNemo.ID))
	assert.ErrorContains(t, err, "404")
}

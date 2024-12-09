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

func TestIntegrationComponentStatusSet(t *testing.T) {
	s := serverTest(t)
	scopes := []string{"create:component-status"}

	scopedRealClientTests(t, scopes, func(ctx context.Context, authToken string, respCode int, expectError bool) error {
		s.Client.SetToken(authToken)
		componentStatus := &fleetdbapi.ComponentStatus{
			ServerComponentID: uuid.MustParse(dbtools.FixtureNemoRightFin.ID),
			Health:            "healthy",
			State:             "running",
			Info:              "all systems normal",
		}
		// update existing component status
		resp, err := s.Client.SetComponentStatus(ctx, componentStatus)
		if !expectError {
			require.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, "resource created", resp.Message)
		}

		return err
	})

	// inserts new component status
	componentStatus := &fleetdbapi.ComponentStatus{
		ServerComponentID: uuid.MustParse(dbtools.FixtureMarlinLeftFin.ID),
		Health:            "error",
		State:             "on-fire",
		Info:              "this is fine",
	}

	_, err := s.Client.SetComponentStatus(context.Background(), componentStatus)
	assert.Nil(t, err)

	// insert record with non-existent server component returns error
	componentStatusFails := &fleetdbapi.ComponentStatus{
		ServerComponentID: uuid.New(),
		Health:            "error",
		State:             "on-fire",
		Info:              "this is fine",
	}

	_, err = s.Client.SetComponentStatus(context.Background(), componentStatusFails)
	assert.NotNil(t, err)
}

func TestIntegrationComponentStatusList(t *testing.T) {
	s := serverTest(t)

	scopes := []string{"read:component-status"}
	scopedRealClientTests(t, scopes, func(ctx context.Context, authToken string, respCode int, expectError bool) error {
		s.Client.SetToken(authToken)

		expectCount := len(dbtools.FixtureComponentStatuses)
		got, resp, err := s.Client.ListComponentStatus(ctx)
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

func TestIntegrationComponentStatusGet(t *testing.T) {
	s := serverTest(t)
	componentID := uuid.MustParse(dbtools.FixtureNemoLeftFin.ID)

	scopes := []string{"read:component-status"}
	scopedRealClientTests(t, scopes, func(ctx context.Context, authToken string, _ int, expectError bool) error {
		s.Client.SetToken(authToken)
		got, resp, err := s.Client.GetComponentStatus(ctx, componentID)
		if !expectError {
			require.NoError(t, err)
			assert.NotNil(t, resp.Record)
			assert.NotNil(t, got)
			assert.Equal(t, componentID, got.ServerComponentID)
		}

		return err
	})

	_, _, err := s.Client.GetComponentStatus(context.Background(), uuid.New())
	assert.ErrorContains(t, err, "404")
}

func TestIntegrationComponentStatusDelete(t *testing.T) {
	s := serverTest(t)
	scopes := []string{"delete:component-status"}
	scopedRealClientTests(t, scopes, func(ctx context.Context, authToken string, _ int, expectError bool) error {
		s.Client.SetToken(authToken)

		resp, err := s.Client.DeleteComponentStatus(ctx, uuid.MustParse(dbtools.FixtureNemoLeftFin.ID))
		if !expectError {
			require.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, "resource deleted", resp.Message)
		}
		return err
	})

	s.Client.SetToken(validToken([]string{"read:component-status"}))
	_, _, err := s.Client.GetComponentStatus(context.Background(), uuid.MustParse(dbtools.FixtureNemoLeftFin.ID))
	assert.ErrorContains(t, err, "404")
}

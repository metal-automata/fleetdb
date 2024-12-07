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

func TestCompareCapabilitySlices(t *testing.T) {
	// Helper function to create capability
	createCapability := func(name, description string, enabled bool) *fleetdbapi.ComponentCapability {
		return &fleetdbapi.ComponentCapability{
			ID:                uuid.New(),
			ServerComponentID: uuid.New(),
			ComponentName:     "test-component",
			Name:              name,
			Description:       description,
			Enabled:           enabled,
		}
	}

	tests := []struct {
		name     string
		existing []*fleetdbapi.ComponentCapability
		incoming []*fleetdbapi.ComponentCapability
		want     struct {
			create int
			update int
			delete int
		}
	}{
		{
			name:     "empty slices",
			existing: nil,
			incoming: nil,
			want:     struct{ create, update, delete int }{0, 0, 0},
		},
		{
			name:     "all new capabilities",
			existing: nil,
			incoming: []*fleetdbapi.ComponentCapability{
				createCapability("cap1", "desc1", true),
				createCapability("cap2", "desc2", false),
			},
			want: struct{ create, update, delete int }{2, 0, 0},
		},
		{
			name: "one capability deleted",
			existing: []*fleetdbapi.ComponentCapability{
				createCapability("cap1", "desc1", true),
				createCapability("cap2", "desc2", false),
			},
			incoming: []*fleetdbapi.ComponentCapability{
				createCapability("cap1", "desc1", true),
			},
			want: struct{ create, update, delete int }{0, 0, 1},
		},
		{
			name: "all capabilities deleted",
			existing: []*fleetdbapi.ComponentCapability{
				createCapability("cap1", "desc1", true),
				createCapability("cap2", "desc2", false),
			},
			incoming: nil,
			want:     struct{ create, update, delete int }{0, 0, 2},
		},
		{
			name: "mixed operations",
			existing: []*fleetdbapi.ComponentCapability{
				createCapability("cap1", "desc1", true),
				createCapability("cap2", "desc2", true),
				createCapability("cap3", "desc3", true),
			},
			incoming: []*fleetdbapi.ComponentCapability{
				createCapability("cap1", "desc1", true),     // no change
				createCapability("cap2", "new desc2", true), // description changed
				createCapability("cap3", "desc3", false),    // enabled changed
				createCapability("cap4", "desc4", true),     // new capability
			},
			want: struct{ create, update, delete int }{1, 2, 0},
		},
		{
			name: "handles nil and empty names",
			existing: []*fleetdbapi.ComponentCapability{
				nil,
				createCapability("cap1", "desc1", true),
				createCapability("", "desc2", true),
			},
			incoming: []*fleetdbapi.ComponentCapability{
				createCapability("cap1", "desc1", true),
				nil,
				createCapability("", "desc3", true),
				createCapability("cap2", "desc2", true),
			},
			want: struct{ create, update, delete int }{1, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			create, update, deletes := fleetdbapi.CompareCapabilitySlices(tt.existing, tt.incoming)
			assert.Equal(t, tt.want.create, len(create), "creates slice length mismatch")
			assert.Equal(t, tt.want.update, len(update), "updates slice length mismatch")
			assert.Equal(t, tt.want.delete, len(deletes), "deletes slice length mismatch")
		})
	}
}

func TestIntegrationComponentCapabilitySet(t *testing.T) {
	s := serverTest(t)

	scopes := []string{"create:component-capability"}
	scopedRealClientTests(t, scopes, func(ctx context.Context, authToken string, respCode int, expectError bool) error {
		s.Client.SetToken(authToken)
		capability := &fleetdbapi.ComponentCapability{
			ID:                uuid.New(),
			ServerComponentID: uuid.MustParse(dbtools.FixtureMarlinLeftFin.ID),
			Name:              "power_control",
			Description:       "Allows power control operations",
			Enabled:           true,
		}

		resp, err := s.Client.SetComponentCapability(ctx, []*fleetdbapi.ComponentCapability{capability})
		if !expectError {
			require.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, "resource created", resp.Message)
		}

		return err
	})

	// insert record with non-existent server component returns error
	componentCapabiliyFails := &fleetdbapi.ComponentCapability{
		ServerComponentID: uuid.New(),
		Name:              "foo",
		Description:       "bar",
		Enabled:           true,
	}

	_, err := s.Client.SetComponentCapability(context.Background(), []*fleetdbapi.ComponentCapability{componentCapabiliyFails})
	assert.NotNil(t, err)

}

func TestIntegrationComponentCapabilityGet(t *testing.T) {
	s := serverTest(t)
	componentID := uuid.MustParse(dbtools.FixtureNemoLeftFin.ID)

	scopes := []string{"read:component-capability"}
	scopedRealClientTests(t, scopes, func(ctx context.Context, authToken string, _ int, expectError bool) error {
		s.Client.SetToken(authToken)
		got, resp, err := s.Client.GetComponentCapability(ctx, componentID)
		if !expectError {
			require.NoError(t, err)
			assert.NotNil(t, resp.Record)
			assert.NotNil(t, got)
			assert.Equal(t, componentID, got.ServerComponentID)
		}

		return err
	})

	_, _, err := s.Client.GetComponentCapability(context.Background(), uuid.New())
	assert.ErrorContains(t, err, "404")
}

func TestIntegrationComponentCapabilityDelete(t *testing.T) {
	s := serverTest(t)
	scopes := []string{"delete:component-capability"}
	componentID := uuid.MustParse(dbtools.FixtureNemoLeftFin.ID)

	scopedRealClientTests(t, scopes, func(ctx context.Context, authToken string, _ int, expectError bool) error {
		s.Client.SetToken(authToken)

		resp, err := s.Client.DeleteComponentCapability(ctx, componentID)
		if !expectError {
			require.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, "resource deleted", resp.Message)
		}
		return err
	})

	s.Client.SetToken(validToken([]string{"read:component-capability"}))
	_, _, err := s.Client.GetComponentCapability(context.Background(), componentID)
	assert.ErrorContains(t, err, "404")
}

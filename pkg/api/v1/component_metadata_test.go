package fleetdbapi_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/google/uuid"
	"github.com/metal-automata/fleetdb/internal/dbtools"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/sqlboiler/v4/types"

	fleetdbapi "github.com/metal-automata/fleetdb/pkg/api/v1"
)

func TestCompareMetadataSlices(t *testing.T) {
	// Helper function to create metadata
	createMetadata := func(namespace string, data map[string]string) *fleetdbapi.ComponentMetadata {
		jsonData, _ := json.Marshal(data)
		return &fleetdbapi.ComponentMetadata{
			ID:                uuid.New(),
			ServerComponentID: uuid.New(),
			ComponentName:     "test-component",
			Namespace:         namespace,
			Data:              jsonData,
		}
	}

	tests := []struct {
		name     string
		existing []*fleetdbapi.ComponentMetadata
		incoming []*fleetdbapi.ComponentMetadata
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
			want: struct {
				create int
				update int
				delete int
			}{0, 0, 0},
		},
		{
			name:     "all new items",
			existing: nil,
			incoming: []*fleetdbapi.ComponentMetadata{
				createMetadata("ns1", map[string]string{"key": "value"}),
				createMetadata("ns2", map[string]string{"key": "value"}),
			},
			want: struct {
				create int
				update int
				delete int
			}{2, 0, 0},
		},
		{
			name: "all items deleted",
			existing: []*fleetdbapi.ComponentMetadata{
				createMetadata("ns1", map[string]string{"key": "value"}),
				createMetadata("ns2", map[string]string{"key": "value"}),
			},
			incoming: nil,
			want: struct {
				create int
				update int
				delete int
			}{0, 0, 2},
		},
		{
			name: "mixed operations",
			existing: []*fleetdbapi.ComponentMetadata{
				createMetadata("ns1", map[string]string{"key": "value1"}),
				createMetadata("ns2", map[string]string{"key": "value2"}),
				createMetadata("ns3", map[string]string{"key": "value3"}),
			},
			incoming: []*fleetdbapi.ComponentMetadata{
				createMetadata("ns1", map[string]string{"key": "value1"}),    // no change
				createMetadata("ns2", map[string]string{"key": "newvalue2"}), // update
				createMetadata("ns4", map[string]string{"key": "value4"}),    // create
			},
			want: struct {
				create int
				update int
				delete int
			}{1, 1, 1},
		},
		{
			name: "handles nil values in slices",
			existing: []*fleetdbapi.ComponentMetadata{
				nil,
				createMetadata("ns1", map[string]string{"key": "value1"}),
				nil,
			},
			incoming: []*fleetdbapi.ComponentMetadata{
				createMetadata("ns1", map[string]string{"key": "value1"}),
				nil,
				createMetadata("ns2", map[string]string{"key": "value2"}),
			},
			want: struct {
				create int
				update int
				delete int
			}{1, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			create, update, delete := fleetdbapi.CompareMetadataSlices(tt.existing, tt.incoming)

			assert.Equal(t, tt.want.create, len(create), "create slice length mismatch")
			assert.Equal(t, tt.want.update, len(update), "update slice length mismatch")
			assert.Equal(t, tt.want.delete, len(delete), "delete slice length mismatch")

			// Additional checks for the mixed operations test
			if tt.name == "mixed operations" {
				// Verify specific items
				assert.Equal(t, "ns4", create[0].Namespace, "expected ns4 to be created")
				assert.Equal(t, "ns2", update[0].Namespace, "expected ns2 to be updated")
				assert.Equal(t, "ns3", delete[0].Namespace, "expected ns3 to be deleted")

				// Verify the content of the updated item
				var updateData map[string]string
				err := json.Unmarshal(update[0].Data, &updateData)
				assert.NoError(t, err)
				assert.Equal(t, "newvalue2", updateData["key"])
			}
		})
	}
}

func TestIntegrationComponentMetadataSet(t *testing.T) {
	s := serverTest(t)

	scopes := []string{"create:component-metadata"}
	scopedRealClientTests(t, scopes, func(ctx context.Context, authToken string, respCode int, expectError bool) error {
		s.Client.SetToken(authToken)

		nicData1, _ := json.Marshal(map[string]string{
			"driver":   "bcrm",
			"duplex":   "full",
			"firmware": "999",
			"link":     "no",
		})

		metadata := &fleetdbapi.ComponentMetadata{
			ID:                uuid.New(),
			ServerComponentID: uuid.MustParse(dbtools.FixtureNemoLeftFin.ID),
			Namespace:         dbtools.FixtureComponentMetadataNS,
			Data:              types.JSON(nicData1),
		}

		resp, err := s.Client.SetComponentMetadata(ctx, []*fleetdbapi.ComponentMetadata{metadata})
		if !expectError {
			require.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, "resource created", resp.Message)
		}

		return err
	})

	// insert record with non-existent server component returns error
	componentMetadataFails := &fleetdbapi.ComponentMetadata{
		ServerComponentID: uuid.New(),
		Namespace:         dbtools.FixtureComponentMetadataNS,
		Data:              []byte(`{"hello": "world"}`),
	}

	_, err := s.Client.SetComponentMetadata(context.Background(), []*fleetdbapi.ComponentMetadata{componentMetadataFails})
	assert.NotNil(t, err)
}

func TestIntegrationComponentMetadataList(t *testing.T) {
	s := serverTest(t)
	componentID := uuid.MustParse(dbtools.FixtureNemoLeftFin.ID)
	namespace := fleetdbapi.ComponentMetadataGenericNS

	scopes := []string{"read:component-metadata"}
	scopedRealClientTests(t, scopes, func(ctx context.Context, authToken string, respCode int, expectError bool) error {
		s.Client.SetToken(authToken)

		expectCount := 1
		got, resp, err := s.Client.ListComponentMetadata(ctx, componentID, "")
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

	// Test filtering
	t.Run("filter by component", func(t *testing.T) {
		s.Client.SetToken(validToken(scopes))
		got, _, err := s.Client.ListComponentMetadata(context.Background(), componentID, "")
		require.NoError(t, err)
		for _, metadata := range got {
			assert.Equal(t, componentID, metadata.ServerComponentID)
		}
	})

	t.Run("filter by namespace", func(t *testing.T) {
		s.Client.SetToken(validToken(scopes))
		got, _, err := s.Client.ListComponentMetadata(context.Background(), componentID, namespace)
		require.NoError(t, err)
		for _, metadata := range got {
			assert.Equal(t, namespace, metadata.Namespace)
		}
	})
}

func TestIntegrationComponentMetadataGet(t *testing.T) {
	s := serverTest(t)
	componentID := uuid.MustParse(dbtools.FixtureNemoLeftFin.ID)
	namespace := fleetdbapi.ComponentMetadataGenericNS

	scopes := []string{"read:component-metadata"}
	scopedRealClientTests(t, scopes, func(ctx context.Context, authToken string, _ int, expectError bool) error {
		s.Client.SetToken(authToken)
		got, resp, err := s.Client.GetComponentMetadata(ctx, componentID, namespace)
		if !expectError {
			require.NoError(t, err)
			assert.NotNil(t, resp.Record)
			assert.NotNil(t, got)
			assert.Equal(t, componentID, got.ServerComponentID)
			assert.Equal(t, namespace, got.Namespace)
		}

		return err
	})

	_, _, err := s.Client.GetComponentMetadata(context.Background(), uuid.New(), "nonexistent")
	assert.ErrorContains(t, err, "404")
}

func TestIntegrationComponentMetadataDelete(t *testing.T) {
	s := serverTest(t)
	scopes := []string{"delete:component-metadata"}
	componentID := uuid.MustParse(dbtools.FixtureNemoLeftFin.ID)
	namespace := fleetdbapi.ComponentMetadataGenericNS

	scopedRealClientTests(t, scopes, func(ctx context.Context, authToken string, _ int, expectError bool) error {
		s.Client.SetToken(authToken)

		resp, err := s.Client.DeleteComponentMetadata(ctx, componentID, namespace)
		if !expectError {
			require.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, "resource deleted", resp.Message)
		}
		return err
	})

	s.Client.SetToken(validToken([]string{"read:component-metadata"}))
	_, _, err := s.Client.GetComponentMetadata(context.Background(), componentID, namespace)
	assert.ErrorContains(t, err, "404")
}

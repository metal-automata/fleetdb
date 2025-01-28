package fleetdbapi_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/metal-automata/fleetdb/internal/dbtools"
	"github.com/metal-automata/fleetdb/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	fleetdbapi "github.com/metal-automata/fleetdb/pkg/api/v1"
	"github.com/metal-automata/fleetdb/pkg/api/v1/fixtures"
)

func TestComponentSliceChanges(t *testing.T) {
	componentID1 := uuid.New()
	componentID2 := uuid.New()
	componentID3 := uuid.New()
	componentID4 := uuid.New()

	// Helper function to create a component
	createComponent := func(
		id uuid.UUID,
		slug,
		serial,
		vendor,
		model,
		firmwareVersion,
		status,
		capability string,
		metadata []byte,
	) *fleetdbapi.ServerComponent {
		c := &fleetdbapi.ServerComponent{
			UUID:              id,
			ServerUUID:        uuid.New(),
			Name:              slug,
			Serial:            serial,
			Vendor:            vendor,
			Model:             model,
			InstalledFirmware: &fleetdbapi.InstalledFirmware{Version: firmwareVersion},
			Status:            &fleetdbapi.ComponentStatus{State: status},
			Capabilities:      []*fleetdbapi.ComponentCapability{{Name: capability, Enabled: true}},
		}

		if len(metadata) > 0 {
			c.Metadata = []*fleetdbapi.ComponentMetadata{{Data: metadata}}
		}

		return c
	}

	tests := []struct {
		name        string
		currentSet  fleetdbapi.ServerComponentSlice
		incomingSet fleetdbapi.ServerComponentSlice
		want        struct {
			creates int
			updates int
			deletes int
		}
	}{
		{
			name:        "empty sets",
			currentSet:  nil,
			incomingSet: nil,
			want:        struct{ creates, updates, deletes int }{0, 0, 0},
		},
		{
			name:       "all creates",
			currentSet: nil,
			incomingSet: fleetdbapi.ServerComponentSlice{
				createComponent(componentID1, "cpu", "123", "Intel", "i7", "", "", "", []byte(``)),
				createComponent(componentID2, "ram", "456", "Kingston", "DDR4", "", "", "", []byte(``)),
			},
			want: struct{ creates, updates, deletes int }{2, 0, 0},
		},
		{
			name: "all deletes",
			currentSet: fleetdbapi.ServerComponentSlice{
				createComponent(componentID1, "cpu", "123", "Intel", "i7", "", "", "", []byte(``)),
				createComponent(componentID2, "ram", "456", "Kingston", "DDR4", "", "", "", []byte(``)),
			},
			incomingSet: nil,
			want:        struct{ creates, updates, deletes int }{0, 0, 2},
		},
		{
			name: "mixed operations",
			currentSet: fleetdbapi.ServerComponentSlice{
				createComponent(componentID1, "cpu", "123", "Intel", "i7", "", "", "", []byte(``)),
				createComponent(componentID2, "ram", "456", "Kingston", "DDR4", "", "", "", []byte(``)),
				createComponent(componentID3, "gpu", "789", "Nvidia", "3080", "", "", "", []byte(``)),
			},
			incomingSet: fleetdbapi.ServerComponentSlice{
				createComponent(componentID1, "cpu", "123", "Intel", "i7", "", "", "", []byte(``)),
				createComponent(componentID3, "gpu", "789", "Nvidia", "3081", "", "", "", []byte(``)), // updated component
				createComponent(componentID4, "ssd", "999", "Samsung", "970", "", "", "", []byte(``)), // new component
			},
			want: struct{ creates, updates, deletes int }{1, 1, 1},
		},
		{
			name: "same component different attributes",
			currentSet: fleetdbapi.ServerComponentSlice{
				createComponent(componentID1, "cpu", "123", "Intel", "i7", "", "", "", []byte(``)),
			},
			incomingSet: fleetdbapi.ServerComponentSlice{
				createComponent(componentID1, "cpu", "123", "Intel", "i9", "", "", "", []byte(``)), // model changed
			},
			want: struct{ creates, updates, deletes int }{0, 1, 0},
		},
		{
			name: "same slug different serial",
			currentSet: fleetdbapi.ServerComponentSlice{
				createComponent(componentID1, "cpu", "123", "Intel", "i7", "", "", "", []byte(``)),
			},
			incomingSet: fleetdbapi.ServerComponentSlice{
				createComponent(componentID2, "cpu", "456", "Intel", "i7", "", "", "", []byte(``)), // different serial = new component
			},
			want: struct{ creates, updates, deletes int }{1, 0, 1},
		},
		{
			name: "installed firmware different",
			currentSet: fleetdbapi.ServerComponentSlice{
				createComponent(componentID1, "", "", "", "", "v1.1", "", "", []byte(``)),
			},
			incomingSet: fleetdbapi.ServerComponentSlice{
				createComponent(componentID2, "", "", "", "", "v1.2", "", "", []byte(``)), // firmware version update
			},
			want: struct{ creates, updates, deletes int }{0, 1, 0},
		},
		{
			name: "status different",
			currentSet: fleetdbapi.ServerComponentSlice{
				createComponent(componentID1, "", "", "", "", "", "OK", "", []byte(``)),
			},
			incomingSet: fleetdbapi.ServerComponentSlice{
				createComponent(componentID2, "", "", "", "", "", "WARNING", "", []byte(``)), // firmware version update
			},
			want: struct{ creates, updates, deletes int }{0, 1, 0},
		},
		{
			name: "capability different",
			currentSet: fleetdbapi.ServerComponentSlice{
				createComponent(componentID1, "", "", "", "", "", "", "spins", []byte(``)),
			},
			incomingSet: fleetdbapi.ServerComponentSlice{
				createComponent(componentID2, "", "", "", "", "", "", "jumps", []byte(``)), // firmware version update
			},
			want: struct{ creates, updates, deletes int }{0, 1, 0},
		},
		{
			name: "metadata different",
			currentSet: fleetdbapi.ServerComponentSlice{
				createComponent(componentID1, "", "", "", "", "", "", "", []byte(`{"foo": "bar"}`)),
			},
			incomingSet: fleetdbapi.ServerComponentSlice{
				createComponent(componentID2, "", "", "", "", "", "", "", []byte(`{"bar": "bar"}`)), // firmware version update
			},
			want: struct{ creates, updates, deletes int }{0, 1, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			creates, updates, deletes := tt.currentSet.Compare(tt.incomingSet)

			assert.Equal(t, tt.want.creates, len(creates), "creates length mismatch")
			assert.Equal(t, tt.want.updates, len(updates), "updates length mismatch")
			assert.Equal(t, tt.want.deletes, len(deletes), "deletes length mismatch")
		})
	}
}

func TestIntegrationServerGetComponents(t *testing.T) {
	s := serverTest(t)
	realClientTests(t, func(ctx context.Context, authToken string, _ int, expectError bool) error {
		s.Client.SetToken(authToken)
		components, _, err := s.Client.GetComponents(ctx, uuid.MustParse(dbtools.FixtureNemo.ID), nil)
		if !expectError {
			require.NoError(t, err)
			assert.Len(t, components, len(dbtools.FixtureNemoComponents))
		}

		return err
	})

	// The Nemo tail component has no firmware or status records
	expectFirmwareAndStatus := func(component *fleetdbapi.ServerComponent) bool {
		return component.UUID.String() != dbtools.FixtureNemoTail.ID
	}

	var testCases = []struct {
		testName string
		params   fleetdbapi.ServerComponentGetParams
		errorMsg string
	}{
		{
			testName: "no includes",
			params:   fleetdbapi.ServerComponentGetParams{},
		},
		{
			testName: "include firmware",
			params: fleetdbapi.ServerComponentGetParams{
				InstalledFirmware: true,
			},
		},
		{
			testName: "include status",
			params: fleetdbapi.ServerComponentGetParams{
				Status: true,
			},
		},
		{
			testName: "include metadata",
			params: fleetdbapi.ServerComponentGetParams{
				Metadata: []string{dbtools.FixtureComponentMetadataNS},
			},
		},
		{
			testName: "include capabilities",
			params: fleetdbapi.ServerComponentGetParams{
				Capabilities: true,
			},
		},
	}

	var wantComponentUUIDs []string
	for _, fc := range dbtools.FixtureNemoComponents {
		wantComponentUUIDs = append(wantComponentUUIDs, fc.ID)
	}

	serverID := uuid.MustParse(dbtools.FixtureNemo.ID)
	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {

			got, _, err := s.Client.GetComponents(context.TODO(), serverID, &tt.params)
			if tt.errorMsg != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMsg)
				return
			}

			assert.Nil(t, err)

			assert.Equal(t, len(dbtools.FixtureNemoComponents), len(got))

			var gotComponentUUIDs []string
			for _, component := range got {
				assert.Equal(t, serverID, component.ServerUUID)
				gotComponentUUIDs = append(gotComponentUUIDs, component.UUID.String())
				if expectFirmwareAndStatus(component) && tt.params.InstalledFirmware {
					assert.NotNil(t, component.InstalledFirmware)
				} else {
					assert.Nil(t, component.InstalledFirmware)
				}

				if expectFirmwareAndStatus(component) && tt.params.Status {
					assert.NotNil(t, component.Status)
				} else {
					assert.Nil(t, component.Status)
				}

				if tt.params.Capabilities {
					assert.NotNil(t, component.Capabilities)
				} else {
					assert.Nil(t, component.Capabilities)
				}

				if len(tt.params.Metadata) > 0 {
					assert.NotNil(t, component.Metadata)
				} else {
					assert.Nil(t, component.Metadata)
				}
			}

			assert.ElementsMatch(t, wantComponentUUIDs, gotComponentUUIDs)
		})
	}
}

func TestIntegrationServerComponentsInit(t *testing.T) {
	s := serverTest(t)
	testDB := dbtools.TestDatastore(t)

	// prep components for insert
	serverID := uuid.MustParse(dbtools.FixturePufferfish.ID)
	conv := fleetdbapi.NewComponentConverter(fleetdbapi.Inband, componentSlugTypeMap(), false)
	inventory, err := conv.FromCommonDevice(serverID, fixtures.DellR6515)
	assert.Nil(t, err)

	purgePufferfishComponents := func(t *testing.T) {
		models.ServerComponents(
			qm.Where(models.ServerComponentColumns.ServerID+"=?", serverID),
		).DeleteAll(context.Background(), testDB)
		require.NoError(t, err)
	}

	// Test auth scenarios
	scopes := []string{"create:server", "create:server:component"}
	scopedRealClientTests(t, scopes, func(ctx context.Context, authToken string, respCode int, expectError bool) error {
		s.Client.SetToken(authToken)
		//s.Client.SetDumper(os.Stdout)
		_, err := s.Client.InitComponentCollection(ctx, serverID, inventory.Components, fleetdbapi.Inband)
		if !expectError {
			require.NoError(t, err)

			// verify

			// components inserted
			mods := []qm.QueryMod{
				qm.Where("server_id = ?", serverID),
			}
			count, err := models.ServerComponents(mods...).Count(context.Background(), testDB)
			assert.NoError(t, err)
			assert.Equal(t, len(inventory.Components), int(count))

			// server inventory ts refreshed
			server, err := models.FindServer(context.Background(), testDB, serverID.String())
			assert.NoError(t, err)
			assert.True(t, server.InventoryRefreshedAt.Valid)

			// cleanup
			purgePufferfishComponents(t)
		}

		return err
	})

	// nolint:typecheck // composite literal types are obvious
	testCases := []struct {
		name       string
		serverID   string
		components fleetdbapi.ServerComponentSlice
		setupFn    func(t *testing.T)
		verifyFn   func(t *testing.T)
		errorMsg   string
	}{
		{
			name:       "empty components",
			serverID:   serverID.String(),
			components: fleetdbapi.ServerComponentSlice{},
			errorMsg:   "ServerComponentSlice is empty",
		},
		{
			name:     "unknown component type",
			serverID: serverID.String(),
			components: fleetdbapi.ServerComponentSlice{
				{
					ServerUUID: serverID,
					Name:       "not_exists",
					Model:      "Model",
					Serial:     "Serial",
				},
			},
			errorMsg: fleetdbapi.ErrComponentType.Error(),
		},
		{
			name:     "server with existing components",
			serverID: serverID.String(),
			components: fleetdbapi.ServerComponentSlice{
				{
					Name:       "fins",
					Model:      "NewModel",
					Serial:     "NewSerial",
					ServerUUID: serverID,
				},
			},
			setupFn: func(t *testing.T) {
				// Add a component to make initialization fail
				comp := &models.ServerComponent{
					ServerID:              serverID.String(),
					ServerComponentTypeID: dbtools.FixtureComponentTypeSlugMap["fins"].ID,
					Model:                 null.StringFrom("Existing"),
					Serial:                "Existing",
				}
				err := comp.Insert(context.Background(), testDB, boil.Infer())
				require.NoError(t, err)
			},
			errorMsg: "use the components/update or component-changes/report instead",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setupFn != nil {
				tt.setupFn(t)
			}

			_, err := s.Client.InitComponentCollection(
				context.Background(),
				uuid.MustParse(tt.serverID),
				tt.components,
				fleetdbapi.Inband,
			)

			if tt.errorMsg != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMsg)
				return
			}

			assert.NoError(t, err)
			if tt.verifyFn != nil {
				tt.verifyFn(t)
			}
		})
	}
}

func TestIntegrationServerComponentsUpdate(t *testing.T) {
	s := serverTest(t)
	testDB := dbtools.TestDatastore(t)

	// prep server and components for update
	serverID := uuid.MustParse(dbtools.FixtureNemo.ID)
	updateComponent := &fleetdbapi.ServerComponent{
		Name:       "fins", // matches FixtureNemoLeftFin's type
		Model:      "UpdatedModel",
		Serial:     dbtools.FixtureNemoLeftFin.Serial, // match existing serial
		ServerUUID: serverID,
		UUID:       uuid.MustParse(dbtools.FixtureNemoLeftFin.ID),
	}

	// Test auth scenarios
	scopes := []string{"update:server", "update:server:component"}
	scopedRealClientTests(t, scopes, func(ctx context.Context, authToken string, respCode int, expectError bool) error {
		s.Client.SetToken(authToken)
		//s.Client.SetDumper(os.Stdout)
		_, err := s.Client.UpdateComponentCollection(ctx, serverID, fleetdbapi.ServerComponentSlice{updateComponent}, fleetdbapi.Inband)
		if !expectError {
			require.NoError(t, err)
			// verify update
			mods := []qm.QueryMod{
				qm.Where("server_id = ?", serverID),
				qm.Where("serial = ?", updateComponent.Serial),
			}
			comp, err := models.ServerComponents(mods...).One(context.Background(), testDB)
			assert.NoError(t, err)
			assert.Equal(t, updateComponent.Model, comp.Model.String)
		}
		return err
	})

	testCases := []struct {
		name       string
		serverID   string
		components fleetdbapi.ServerComponentSlice
		setupFn    func(t *testing.T)
		verifyFn   func(t *testing.T)
		errorMsg   string
	}{

		{
			name:       "empty components",
			serverID:   serverID.String(),
			components: fleetdbapi.ServerComponentSlice{},
			errorMsg:   "ServerComponentSlice is empty",
		},
		{
			name:     "non-existent serial",
			serverID: serverID.String(),
			components: fleetdbapi.ServerComponentSlice{
				{
					Name:       "fins",
					Model:      "NewModel",
					Serial:     "NonExistentSerial",
					ServerUUID: serverID,
				},
			},
			errorMsg: "unknown component",
		},
		{
			name:     "server not found",
			serverID: uuid.New().String(),
			components: fleetdbapi.ServerComponentSlice{
				updateComponent,
			},
			errorMsg: "resource not found",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setupFn != nil {
				tt.setupFn(t)
			}
			var err error
			_, err = s.Client.UpdateComponentCollection(
				context.Background(),
				uuid.MustParse(tt.serverID),
				tt.components,
				fleetdbapi.Inband,
			)
			if tt.errorMsg != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMsg)
				return
			}
			assert.NoError(t, err)
			if tt.verifyFn != nil {
				tt.verifyFn(t)
			}
		})
	}
}

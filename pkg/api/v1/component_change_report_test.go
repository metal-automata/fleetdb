package fleetdbapi_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/metal-automata/fleetdb/internal/dbtools"
	"github.com/metal-automata/fleetdb/internal/models"
	fleetdbapi "github.com/metal-automata/fleetdb/pkg/api/v1"
)

func TestIntegrationComponentChangeReport(t *testing.T) {
	s := serverTest(t)

	testDB := dbtools.TestDatastore(t)

	// Helper function to create a component
	createComponent := func(name, serial string) *fleetdbapi.ServerComponent {
		return &fleetdbapi.ServerComponent{
			ServerUUID: uuid.MustParse(dbtools.FixtureNemo.ID),
			Name:       name,
			Serial:     serial,
		}
	}

	scopes := []string{"create:component-change"}
	scopedRealClientTests(t, scopes, func(ctx context.Context, authToken string, respCode int, expectError bool) error {
		s.Client.SetToken(authToken)
		//s.Client.SetDumper(os.Stdout)
		change := &fleetdbapi.ComponentChangeReport{
			CollectionMethod: string(fleetdbapi.Inband),
			Creates: []*fleetdbapi.ServerComponent{
				createComponent("cpu", "123"),
				createComponent("physicalmemory", "456"),
			},
			Deletes: []*fleetdbapi.ServerComponent{
				fleetdbapi.ServerComponentFromModel(dbtools.FixtureNemoLeftFin),
			},
		}

		_, resp, err := s.Client.ReportComponentChanges(ctx, dbtools.FixtureNemo.ID, change)
		if !expectError {
			require.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, "resource created", resp.Message)
		}

		return err
	})

	tests := []struct {
		name          string
		serverID      string
		report        *fleetdbapi.ComponentChangeReport
		expectError   bool
		errorContains string
	}{
		{
			name:     "add new components",
			serverID: dbtools.FixtureNemo.ID,
			report: &fleetdbapi.ComponentChangeReport{
				CollectionMethod: "inband",
				Creates: []*fleetdbapi.ServerComponent{
					createComponent("cpu", "CPU123"),
					createComponent("physicalmemory", "MEM456"),
				},
			},
		},
		{
			name:     "remove existing components",
			serverID: dbtools.FixtureNemo.ID,
			report: &fleetdbapi.ComponentChangeReport{
				CollectionMethod: "inband",
				Deletes: []*fleetdbapi.ServerComponent{
					fleetdbapi.ServerComponentFromModel(dbtools.FixtureNemoLeftFin),
				},
			},
		},
		{
			name:     "both add and remove",
			serverID: dbtools.FixtureNemo.ID,
			report: &fleetdbapi.ComponentChangeReport{
				CollectionMethod: "inband",
				Creates: []*fleetdbapi.ServerComponent{
					createComponent("gpu", "GPU001"),
				},
				Deletes: []*fleetdbapi.ServerComponent{
					fleetdbapi.ServerComponentFromModel(dbtools.FixtureNemoRightFin),
				},
			},
		},
		{
			name:     "using outofband collection",
			serverID: dbtools.FixtureNemo.ID,
			report: &fleetdbapi.ComponentChangeReport{
				CollectionMethod: "outofband",
				Creates: []*fleetdbapi.ServerComponent{
					createComponent("nic", "NIC001"),
				},
			},
		},
		{
			name:     "invalid collection method",
			serverID: dbtools.FixtureNemo.ID,
			report: &fleetdbapi.ComponentChangeReport{
				CollectionMethod: "invalid",
				Creates: []*fleetdbapi.ServerComponent{
					createComponent("cpu", "CPU123"),
				},
			},
			expectError:   true,
			errorContains: "unexpected CollectionMethod",
		},
		{
			name:     "duplicate components - add",
			serverID: dbtools.FixtureNemo.ID,
			report: &fleetdbapi.ComponentChangeReport{
				CollectionMethod: "inband",
				Creates: []*fleetdbapi.ServerComponent{
					createComponent("cpu", "CPU123"),
					createComponent("cpu", "CPU123"),
				},
			},
			expectError:   true,
			errorContains: "duplicate component",
		},
		{
			name:     "duplicate components - delete",
			serverID: dbtools.FixtureNemo.ID,
			report: &fleetdbapi.ComponentChangeReport{
				CollectionMethod: "inband",
				Deletes: []*fleetdbapi.ServerComponent{
					fleetdbapi.ServerComponentFromModel(dbtools.FixtureNemoRightFin),
					fleetdbapi.ServerComponentFromModel(dbtools.FixtureNemoRightFin),
				},
			},
			expectError:   true,
			errorContains: "duplicate component",
		},
		{
			name:     "duplicate components - mix add, delete",
			serverID: dbtools.FixtureNemo.ID,
			report: &fleetdbapi.ComponentChangeReport{
				CollectionMethod: "inband",
				Creates: []*fleetdbapi.ServerComponent{
					createComponent("cpu", "CPU123"),
					createComponent("cpu", "CPU123"),
				},
				Deletes: []*fleetdbapi.ServerComponent{
					fleetdbapi.ServerComponentFromModel(dbtools.FixtureNemoRightFin),
					fleetdbapi.ServerComponentFromModel(dbtools.FixtureNemoRightFin),
				},
			},
			expectError:   true,
			errorContains: "duplicate component",
		},

		{
			name:     "empty changes list",
			serverID: dbtools.FixtureNemo.ID,
			report: &fleetdbapi.ComponentChangeReport{
				CollectionMethod: "inband",
			},
			expectError:   true,
			errorContains: "expected components to Creates/Deletes",
		},
		{
			name:     "server id mismatch",
			serverID: dbtools.FixtureNemo.ID,
			report: &fleetdbapi.ComponentChangeReport{
				CollectionMethod: "inband",
				Creates: []*fleetdbapi.ServerComponent{
					func() *fleetdbapi.ServerComponent {
						c := createComponent("cpu", "CPU123")
						c.ServerUUID = uuid.New() // Different server ID
						return c
					}(),
				},
			},
			expectError:   true,
			errorContains: "serverID mismatch",
		},
		{
			name:     "invalid server id",
			serverID: "not-a-uuid",
			report: &fleetdbapi.ComponentChangeReport{
				CollectionMethod: "inband",
				Creates: []*fleetdbapi.ServerComponent{
					createComponent("cpu", "CPU123"),
				},
			},
			expectError:   true,
			errorContains: "400",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scopes := []string{"create:component-change"}
			s.Client.SetToken(validToken(scopes))
			result, res, err := s.Client.ReportComponentChanges(context.Background(), tt.serverID, tt.report)

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorContains)
				return
			}

			require.NoError(t, err)
			assert.NotNil(t, res)
			assert.Equal(t, "resource created", res.Message)
			assert.Equal(t, len(tt.report.Creates), len(result.ChangeIDCreates))
			assert.Equal(t, len(tt.report.Deletes), len(result.ChangeIDDeletes))

			// verify adds
			if len(tt.report.Creates) > 0 {
				mods := []qm.QueryMod{
					qm.Where(models.ComponentChangeReportColumns.ReportID+"=?", result.ReportID),
					qm.And(models.ComponentChangeReportColumns.RemoveComponent+"=?", false),
				}

				got, err := models.ComponentChangeReports(mods...).All(context.Background(), testDB)
				assert.Nil(t, err)
				assert.Equal(t, len(tt.report.Creates), len(got))
				for _, change := range got {
					assert.False(t, change.RemoveComponent.Bool)
					assert.NotNil(t, change.Data)
					assert.NotNil(t, change.ServerComponentTypeID)
				}
			}

			// verify removes
			if len(tt.report.Deletes) > 0 {
				mods := []qm.QueryMod{
					qm.Where(models.ComponentChangeReportColumns.ReportID+"=?", result.ReportID),
					qm.And(models.ComponentChangeReportColumns.RemoveComponent+"=?", true),
				}

				//boil.DebugMode = true
				//defer func() { boil.DebugMode = false }()
				got, err := models.ComponentChangeReports(mods...).All(context.Background(), testDB)
				assert.Nil(t, err)
				assert.Equal(t, len(tt.report.Deletes), len(got))
				for _, change := range got {
					assert.True(t, change.RemoveComponent.Bool)
					assert.NotNil(t, change.ServerComponentTypeID)
				}
			}
		})
	}
}

func TestIntegrationComponentChangeAccept(t *testing.T) {
	s := serverTest(t)

	scopes := []string{"update:component-change"}
	scopedRealClientTests(t, scopes, func(ctx context.Context, authToken string, respCode int, expectError bool) error {
		s.Client.SetToken(authToken)
		resp, err := s.Client.AcceptComponentChanges(ctx, dbtools.FixtureNemo.ID, []string{dbtools.FixtureNemoComponentChangeReportAdd.ID})
		if !expectError {
			assert.Nil(t, err)
			assert.Equal(t, "resource created", resp.Message)
		}

		return err
	})

	// Helper to create a change report and get the report ID
	createChangeReport := func(t *testing.T, report *fleetdbapi.ComponentChangeReport) (createIDs, deleteIDs []string, reportID string) {
		s.Client.SetToken(validToken([]string{"create:component-change"}))
		result, resp, err := s.Client.ReportComponentChanges(context.Background(), dbtools.FixtureNemo.ID, report)
		require.NoError(t, err)
		require.NotNil(t, resp)
		return result.ChangeIDCreates, result.ChangeIDDeletes, result.ReportID
	}

	testDB := dbtools.TestDatastore(t)

	tests := []struct {
		name          string
		serverID      string
		setupFn       func(t *testing.T) (createIDs, deleteIDs []string, reportID string)
		expectError   bool
		errorContains string
		verifyFn      func(t *testing.T, reportID string)
	}{
		{
			name:     "nonexistent change IDs",
			serverID: dbtools.FixtureNemo.ID,
			setupFn: func(t *testing.T) (createIDs, deleteIDs []string, reportID string) {
				return []string{uuid.NewString()}, []string{uuid.NewString()}, uuid.NewString()
			},
			expectError:   true,
			errorContains: "no changes identified",
		},
		{
			name:     "invalid server ID",
			serverID: "not-a-uuid",
			setupFn: func(t *testing.T) (createIDs, deleteIDs []string, reportID string) {
				return []string{uuid.NewString()}, []string{uuid.NewString()}, uuid.NewString()
			},
			expectError:   true,
			errorContains: "400",
		},

		{
			name:     "accept component additions",
			serverID: dbtools.FixtureNemo.ID,
			setupFn: func(t *testing.T) (createIDs, deleteIDs []string, reportID string) {
				report := &fleetdbapi.ComponentChangeReport{
					CollectionMethod: "inband",
					Creates: []*fleetdbapi.ServerComponent{
						{
							ServerUUID: uuid.MustParse(dbtools.FixtureNemo.ID),
							Name:       "cpu",
							Serial:     "CPU9991",
						},
					},
				}

				return createChangeReport(t, report)
			},
			verifyFn: func(t *testing.T, reportID string) {
				// Verify component was added to server_components
				mods := []qm.QueryMod{
					qm.Where("serial = ?", "CPU9991"),
				}

				exists, err := models.ServerComponents(mods...).Exists(context.Background(), testDB)
				assert.NoError(t, err)
				assert.True(t, exists)

				// Verify change report was deleted
				mods = []qm.QueryMod{
					qm.Where("serial = ?", "CPU9991"),
				}
				exists, err = models.ComponentChangeReports(mods...).Exists(context.Background(), testDB)
				assert.NoError(t, err)
				assert.False(t, exists)
			},
		},
		{
			name:     "accept component removals",
			serverID: dbtools.FixtureNemo.ID,
			setupFn: func(t *testing.T) (createIDs, deleteIDs []string, reportID string) {
				report := &fleetdbapi.ComponentChangeReport{
					CollectionMethod: "inband",
					Deletes: []*fleetdbapi.ServerComponent{
						fleetdbapi.ServerComponentFromModel(dbtools.FixtureNemoRightFin),
					},
				}
				return createChangeReport(t, report)
			},
			verifyFn: func(t *testing.T, reportID string) {
				// Verify component was removed from server_components
				mods := []qm.QueryMod{
					qm.Where("id = ?", dbtools.FixtureNemoRightFin.ID),
				}
				exists, err := models.ServerComponents(mods...).Exists(context.Background(), testDB)
				assert.NoError(t, err)
				assert.False(t, exists)

				// Verify change report was deleted
				exists, err = models.ComponentChangeReports(mods...).Exists(context.Background(), testDB)
				assert.NoError(t, err)
				assert.False(t, exists)
			},
		},
		{
			name:     "accept multiple changes",
			serverID: dbtools.FixtureNemo.ID,
			setupFn: func(t *testing.T) (createIDs, deleteIDs []string, reportID string) {
				report := &fleetdbapi.ComponentChangeReport{
					CollectionMethod: "inband",
					Creates: []*fleetdbapi.ServerComponent{
						{
							ServerUUID: uuid.MustParse(dbtools.FixtureNemo.ID),
							Name:       "gpu",
							Serial:     "GPU999",
						},
					},
					Deletes: []*fleetdbapi.ServerComponent{
						fleetdbapi.ServerComponentFromModel(dbtools.FixtureNemoRightFin),
					},
				}
				return createChangeReport(t, report)
			},
			verifyFn: func(t *testing.T, reportID string) {
				// Verify new component was added
				mods := []qm.QueryMod{
					qm.Where("serial = ?", "GPU999"),
				}
				exists, err := models.ServerComponents(mods...).Exists(context.Background(), testDB)
				assert.NoError(t, err)
				assert.True(t, exists)

				// Verify old component was removed
				mods = []qm.QueryMod{
					qm.Where("id = ?", dbtools.FixtureNemoRightFin.ID),
				}
				exists, err = models.ServerComponents(mods...).Exists(context.Background(), testDB)
				assert.NoError(t, err)
				assert.False(t, exists)

				// Verify all change reports were deleted
				mods = []qm.QueryMod{
					qm.Where("report_id = ?", reportID),
				}
				exists, err = models.ComponentChangeReports(mods...).Exists(context.Background(), testDB)
				assert.NoError(t, err)
				assert.False(t, exists)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createIDs, deleteIDs, reportID := tt.setupFn(t)

			all := []string{}
			all = append(all, createIDs...)
			all = append(all, deleteIDs...)
			s.Client.SetToken(validToken(scopes))
			resp, err := s.Client.AcceptComponentChanges(context.Background(), tt.serverID, all)

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorContains)
				return
			}

			require.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, "resource created", resp.Message)

			if tt.verifyFn != nil {
				tt.verifyFn(t, reportID)
			}
		})
	}
}

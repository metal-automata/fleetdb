package fleetdbapi_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/metal-automata/fleetdb/internal/dbtools"
	fleetdbapi "github.com/metal-automata/fleetdb/pkg/api/v1"
)

// TODO for when server list is being implemented
func TestIntegrationServerListPagination(t *testing.T) {
	s := serverTest(t)
	s.Client.SetToken(validToken(adminScopes))
	p := &fleetdbapi.ServerQueryParams{PaginationParams: &fleetdbapi.PaginationParams{Limit: 2, Page: 1}}
	r, resp, err := s.Client.ListServers(context.TODO(), p)
	assert.NoError(t, err)
	assert.Len(t, r, 2)

	assert.EqualValues(t, 2, resp.PageCount)
	assert.EqualValues(t, 3, resp.TotalPages)
	assert.EqualValues(t, 5, resp.TotalRecordCount)
	// Since we have a next page let's make sure all the links are set
	assert.NotNil(t, resp.Links.Next)
	assert.Nil(t, resp.Links.Previous)
	assert.True(t, resp.HasNextPage())

	//
	// Get the next page and verify the results
	//
	resp, err = s.Client.NextPage(context.TODO(), *resp, &r)
	assert.NoError(t, err)
	assert.Len(t, r, 2)

	assert.EqualValues(t, 2, resp.PageCount)

	// get the last page
	resp, err = s.Client.NextPage(context.TODO(), *resp, &r)
	assert.NoError(t, err)
	assert.Len(t, r, 1)

	// we should have followed the cursor so first/previous/next/last links shouldn't be set
	// but there is another page so we should have a next cursor link. Total counts are not includes
	// cursor pages
	assert.EqualValues(t, 3, resp.TotalPages)
	assert.EqualValues(t, 5, resp.TotalRecordCount)
	assert.NotNil(t, resp.Links.First)
	assert.NotNil(t, resp.Links.Previous)
	assert.Nil(t, resp.Links.Next)
	assert.NotNil(t, resp.Links.Last)
	assert.False(t, resp.HasNextPage())
}

func TestIntegrationServerListFilter(t *testing.T) {
	testcases := []struct {
		name   string
		filter []fleetdbapi.Filter
		verify func(s []fleetdbapi.Server, r *fleetdbapi.ServerResponse)
		errMsg string
	}{
		{
			name: "Compare equals",
			filter: []fleetdbapi.Filter{
				{
					Attribute:          "name",
					ComparisonOperator: fleetdbapi.ComparisonOpEqual,
					Value:              "Pufferfish",
				},
			},
			verify: func(s []fleetdbapi.Server, r *fleetdbapi.ServerResponse) {
				assert.Len(t, s, 1)
				assert.Equal(t, "Pufferfish", s[0].Name)
			},
		},
		{
			name: "Compare not equals",
			filter: []fleetdbapi.Filter{
				{
					Attribute:          "name",
					ComparisonOperator: fleetdbapi.ComparisonOpNotEqual,
					Value:              "Pufferfish",
				},
			},
			verify: func(s []fleetdbapi.Server, r *fleetdbapi.ServerResponse) {
				assert.Len(t, s, 1)
				assert.NotEqual(t, "Pufferfish", s[0].Name)
			},
		},
		{
			name: "Starts with",
			filter: []fleetdbapi.Filter{
				{
					Attribute:          "name",
					ComparisonOperator: fleetdbapi.ComparisonOpStartsWith,
					Value:              "Puffer",
				},
			},
			verify: func(s []fleetdbapi.Server, r *fleetdbapi.ServerResponse) {
				assert.Len(t, s, 1)
				assert.Equal(t, "Pufferfish", s[0].Name)
			},
		},
		{
			name: "Starts with, case insensitive",
			filter: []fleetdbapi.Filter{
				{
					Attribute:          "name",
					ComparisonOperator: fleetdbapi.ComparisonOpStartsWith,
					Modifier:           fleetdbapi.ModifierCaseInsensitive,
					Value:              "pUffeR",
				},
			},
			verify: func(s []fleetdbapi.Server, r *fleetdbapi.ServerResponse) {
				assert.Len(t, s, 1)
				assert.Equal(t, "Pufferfish", s[0].Name)
			},
		},
		{
			name: "Ends with",
			filter: []fleetdbapi.Filter{
				{
					Attribute:          "name",
					ComparisonOperator: fleetdbapi.ComparisonOpEndsWith,
					Value:              "fish",
				},
			},
			verify: func(s []fleetdbapi.Server, r *fleetdbapi.ServerResponse) {
				assert.Len(t, s, 1)
				assert.Equal(t, "Pufferfish", s[0].Name)
			},
		},
		{
			name: "Ends with, case insensitive",
			filter: []fleetdbapi.Filter{
				{
					Attribute:          "name",
					ComparisonOperator: fleetdbapi.ComparisonOpEndsWith,
					Modifier:           fleetdbapi.ModifierCaseInsensitive,
					Value:              "fIsH",
				},
			},
			verify: func(s []fleetdbapi.Server, r *fleetdbapi.ServerResponse) {
				assert.Len(t, s, 1)
				assert.Equal(t, "Pufferfish", s[0].Name)
			},
		},
	}

	s := serverTest(t)
	s.Client.SetToken(validToken(adminScopes))

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {

			params := &fleetdbapi.ServerQueryParams{
				PaginationParams: &fleetdbapi.PaginationParams{
					Limit: 1,
					Page:  1,
				},
				FilterParams: &fleetdbapi.FilterParams{
					Target:  &fleetdbapi.Server{},
					Filters: tt.filter,
				},
			}

			r, resp, err := s.Client.ListServers(context.TODO(), params)
			if tt.errMsg != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errMsg)
				return
			}

			tt.verify(r, resp)
			assert.Nil(t, err)
		})
	}

}

func TestIntegrationServerListFilterWithLogicalOp(t *testing.T) {
	testcases := []struct {
		name      string
		filters   []fleetdbapi.Filter
		limit     int
		logicalOp []string
		verify    func(s []fleetdbapi.Server, r *fleetdbapi.ServerResponse)
		errMsg    string
	}{
		{
			name: "Logical operation AND",
			filters: []fleetdbapi.Filter{
				{
					Attribute:          "name",
					ComparisonOperator: fleetdbapi.ComparisonOpStartsWith,
					Value:              "Puffer",
				},
				{
					Attribute:          "facility_code",
					ComparisonOperator: fleetdbapi.ComparisonOpEqual,
					Value:              "EastAustralianCurrent",
				},
			},
			limit:     1,
			logicalOp: []string{"name", "and", "facility_code"},
			verify: func(s []fleetdbapi.Server, r *fleetdbapi.ServerResponse) {
				assert.Len(t, s, 1)
				assert.Equal(t, "Pufferfish", s[0].Name)
			},
		},
		{
			name: "Logical operation OR - same attribute",
			filters: []fleetdbapi.Filter{
				{
					Attribute:          "name",
					ComparisonOperator: fleetdbapi.ComparisonOpStartsWith,
					Value:              "Puffer",
				},
				{
					Attribute:          "name",
					ComparisonOperator: fleetdbapi.ComparisonOpEqual,
					Value:              "Dory",
				},
			},
			limit:     2,
			logicalOp: []string{"name", "or", "name"},
			verify: func(s []fleetdbapi.Server, r *fleetdbapi.ServerResponse) {
				assert.Len(t, s, 2)
			},
		},
		{
			name: "Logical operation OR - different attributes",
			filters: []fleetdbapi.Filter{
				{
					Attribute:          "name",
					ComparisonOperator: fleetdbapi.ComparisonOpEndsWith,
					Value:              "mo",
				},
				{
					Attribute:          "facility_code",
					ComparisonOperator: fleetdbapi.ComparisonOpEqual,
					Value:              "Sydney",
				},
			},
			limit:     2,
			logicalOp: []string{"name", "or", "name"},
			verify: func(s []fleetdbapi.Server, r *fleetdbapi.ServerResponse) {
				assert.Len(t, s, 1)
			},
		},
	}

	s := serverTest(t)
	s.Client.SetToken(validToken(adminScopes))

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {

			params := &fleetdbapi.ServerQueryParams{
				PaginationParams: &fleetdbapi.PaginationParams{
					Limit: tt.limit,
					Page:  1,
				},
				FilterParams: &fleetdbapi.FilterParams{
					Target:           &fleetdbapi.Server{},
					Filters:          tt.filters,
					LogicalOperation: tt.logicalOp,
				},
			}

			r, resp, err := s.Client.ListServers(context.TODO(), params)
			if tt.errMsg != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errMsg)
				return
			}

			tt.verify(r, resp)
			assert.Nil(t, err)
		})
	}

}

func TestIntegrationServerGetPreload(t *testing.T) {
	s := serverTest(t)
	s.Client.SetToken(validToken(adminScopes))
	r, _, err := s.Client.GetServer(context.TODO(), uuid.MustParse(dbtools.FixtureNemo.ID), &fleetdbapi.ServerQueryParams{IncludeComponents: true})
	assert.NoError(t, err)
	assert.Len(t, r.Components, 2, "server components")
	assert.Nil(t, r.DeletedAt, "DeletedAt should be nil for non deleted server")
}

func TestIntegrationServerGetDeleted(t *testing.T) {
	s := serverTest(t)

	realClientTests(t, func(ctx context.Context, authToken string, _ int, expectError bool) error {
		s.Client.SetToken(authToken)

		r, _, err := s.Client.GetServer(ctx, uuid.MustParse(dbtools.FixtureChuckles.ID), nil)
		if !expectError {
			require.NoError(t, err)
			assert.Equal(t, r.UUID, uuid.MustParse(dbtools.FixtureChuckles.ID), "Expected UUID %s, got %s", dbtools.FixtureChuckles.ID, r.UUID.String())
			assert.Equal(t, r.Name, dbtools.FixtureChuckles.Name.String)
			assert.NotNil(t, r.DeletedAt)
		}

		return err
	})
}

func TestIntegrationServerCreate(t *testing.T) {
	s := serverTest(t)

	realClientTests(t, func(ctx context.Context, authToken string, _ int, expectError bool) error {
		s.Client.SetToken(authToken)
		testServer := fleetdbapi.Server{
			UUID:         uuid.New(),
			Name:         "test-server",
			FacilityCode: "int",
			Vendor:       "foo",
			Model:        "123",
		}

		id, resp, err := s.Client.Create(ctx, testServer)
		if !expectError {
			require.NoError(t, err)
			assert.NotNil(t, id)
			assert.Equal(t, testServer.UUID.String(), id.String())
			assert.NotNil(t, resp.Links.Self)
			assert.Equal(t, fmt.Sprintf("http://test.hollow.com/api/v1/servers/%s", id), resp.Links.Self.Href)
		}

		return err
	})

	serverWithBMCID := uuid.New()
	var testCases = []struct {
		testName string
		srv      *fleetdbapi.Server
		errorMsg string
	}{
		{
			"fails on a duplicate uuid",
			&fleetdbapi.Server{
				UUID:         uuid.MustParse(dbtools.FixtureNemo.ID),
				FacilityCode: "int-test",
				Vendor:       "foo",
			},
			"duplicate key",
		},
		{
			"fails on unknown vendor",
			&fleetdbapi.Server{
				UUID:         uuid.New(),
				Name:         "test-server2",
				FacilityCode: "int",
				Vendor:       "unknown",
			},
			"resource not found",
		},
		{
			"fails on unknown model",
			&fleetdbapi.Server{
				UUID:         uuid.New(),
				Name:         "test-server3",
				FacilityCode: "int",
				Vendor:       "foo",
				Model:        "unknown",
			},
			"resource not found",
		},
		{
			"create with BMC attributes",
			&fleetdbapi.Server{
				UUID:         serverWithBMCID,
				Name:         "test-server",
				FacilityCode: "int",
				Vendor:       "foo",
				Model:        "123",
				BMC: &fleetdbapi.ServerBMC{
					ServerID:           serverWithBMCID,
					IPAddress:          "127.0.0.1",
					Username:           "foo",
					Password:           "baz",
					HardwareVendorName: "foo",
					HardwareModelName:  "123",
					MacAddress:         "00:00:de:ad:be:ef",
				},
			},
			"",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			_, _, err := s.Client.Create(context.TODO(), *tt.srv)
			if tt.errorMsg != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMsg)
				return
			}

			assert.Nil(t, err)
		})
	}
}

func TestIntegrationServerDelete(t *testing.T) {
	s := serverTest(t)

	realClientTests(t, func(ctx context.Context, authToken string, _ int, _ bool) error {
		s.Client.SetToken(authToken)
		_, err := s.Client.Delete(ctx, fleetdbapi.Server{UUID: uuid.MustParse(dbtools.FixtureNemo.ID)})

		return err
	})

	var testCases = []struct {
		testName  string
		uuid      uuid.UUID
		errorMsg  string
		expectErr bool
	}{
		{
			"fails on unknown uuid",
			uuid.New(),
			"resource not found",
			false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			_, err := s.Client.Delete(context.TODO(), fleetdbapi.Server{UUID: tt.uuid})
			if tt.expectErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMsg)
			}
		})
	}
}

func TestIntegrationServerUpdate(t *testing.T) {
	s := serverTest(t)

	realClientTests(t, func(ctx context.Context, authToken string, _ int, expectError bool) error {
		s.Client.SetToken(authToken)

		resp, err := s.Client.Update(ctx, uuid.MustParse(dbtools.FixtureDory.ID), fleetdbapi.Server{Name: "The New Dory", Vendor: "ocean", FacilityCode: "pacific"})
		if !expectError {
			require.NoError(t, err)
			assert.NotNil(t, resp.Links.Self)
			assert.Equal(t, fmt.Sprintf("http://test.hollow.com/api/v1/servers/%s", dbtools.FixtureDory.ID), resp.Links.Self.Href)
		}

		return err
	})
}

func TestIntegrationServerGet(t *testing.T) {
	s := serverTest(t)

	// Test auth scenarios
	scopes := []string{"read:server"}
	scopedRealClientTests(t, scopes, func(ctx context.Context, authToken string, respCode int, expectError bool) error {
		s.Client.SetToken(authToken)
		//s.Client.SetDumper(os.Stdout)

		srv, _, err := s.Client.GetServer(ctx, uuid.MustParse(dbtools.FixtureNemo.ID), nil)
		if !expectError {
			require.NoError(t, err)
			assert.Equal(t, dbtools.FixtureNemo.ID, srv.UUID.String())
			assert.Equal(t, dbtools.FixtureNemo.Name.String, srv.Name)
			assert.Equal(t, dbtools.FixtureNemo.FacilityCode.String, srv.FacilityCode)
			assert.Equal(t, dbtools.FixtureHardwareVendorBar.Name, srv.Vendor)
			assert.Equal(t, dbtools.FixtureHardwareModelBar123.Name, srv.Model)
		}
		return err
	})

	testCases := []struct {
		name     string
		serverID string
		params   *fleetdbapi.ServerQueryParams
		setupFn  func(t *testing.T)
		verifyFn func(t *testing.T, srv *fleetdbapi.Server)
		errorMsg string
	}{
		{
			name:     "server not found",
			serverID: uuid.New().String(),
			errorMsg: "resource not found",
		},
		{
			name:     "get with components",
			serverID: dbtools.FixtureNemo.ID,
			params: &fleetdbapi.ServerQueryParams{
				IncludeComponents: true,
			},
			verifyFn: func(t *testing.T, srv *fleetdbapi.Server) {
				assert.NotEmpty(t, srv.Components)
				foundComponents := make(map[string]bool)
				for _, comp := range srv.Components {
					foundComponents[comp.UUID.String()] = true
				}
				assert.True(t, foundComponents[dbtools.FixtureNemoLeftFin.ID])
				assert.True(t, foundComponents[dbtools.FixtureNemoRightFin.ID])

				assert.Equal(t, dbtools.FixtureHardwareVendorBar.Name, srv.Vendor)
				assert.Equal(t, dbtools.FixtureHardwareModelBar123.Name, srv.Model)
			},
		},
		{
			name:     "get with components and thier attributes",
			serverID: dbtools.FixtureNemo.ID,
			params: &fleetdbapi.ServerQueryParams{
				IncludeComponents: true,
				ComponentParams: &fleetdbapi.ServerComponentGetParams{
					InstalledFirmware: true,
					Status:            true,
					Capabilities:      true,
				},
			},
			verifyFn: func(t *testing.T, srv *fleetdbapi.Server) {
				assert.NotEmpty(t, srv.Components)
				foundComponents := make(map[string]bool)
				for _, comp := range srv.Components {
					assert.NotNil(t, comp.InstalledFirmware)
					assert.NotEmpty(t, comp.InstalledFirmware.Version)

					assert.NotNil(t, comp.Status)
					assert.NotEmpty(t, comp.Status.State)

					assert.NotNil(t, comp.Capabilities)
					assert.NotEmpty(t, comp.Capabilities[0].Description)

					foundComponents[comp.UUID.String()] = true
				}

				assert.True(t, foundComponents[dbtools.FixtureNemoLeftFin.ID])
				assert.True(t, foundComponents[dbtools.FixtureNemoRightFin.ID])

				assert.Equal(t, dbtools.FixtureHardwareVendorBar.Name, srv.Vendor)
				assert.Equal(t, dbtools.FixtureHardwareModelBar123.Name, srv.Model)
			},
		},
		{ // fixture Puffer fish has no components - at minimum the server object is returned
			name:     "get with components (no components in fixture)",
			serverID: dbtools.FixturePufferfish.ID,
			params: &fleetdbapi.ServerQueryParams{
				IncludeComponents: true,
				ComponentParams:   &fleetdbapi.ServerComponentGetParams{},
			},
			verifyFn: func(t *testing.T, srv *fleetdbapi.Server) {
				assert.Equal(t, srv.Name, dbtools.FixturePufferfish.Name.String)
			},
		},
		{
			name:     "get with BMC",
			serverID: dbtools.FixtureNemo.ID,
			params: &fleetdbapi.ServerQueryParams{
				IncludeBMC: true,
			},
			verifyFn: func(t *testing.T, srv *fleetdbapi.Server) {
				assert.NotNil(t, srv.BMC)
				assert.Equal(t, "127.0.0.1", srv.BMC.IPAddress)
				assert.Equal(t, "de:ad:be:ef:ca:fe", srv.BMC.MacAddress)
				assert.Equal(t, "user", srv.BMC.Username)
				assert.Equal(t, "super-secret-bmc-password", srv.BMC.Password)
			},
		},
		{
			name:     "get with status",
			serverID: dbtools.FixtureNemo.ID,
			params: &fleetdbapi.ServerQueryParams{
				IncludeStatus: true,
			},
			verifyFn: func(t *testing.T, srv *fleetdbapi.Server) {
				assert.NotNil(t, srv.Status)
				assert.Equal(t, "healthy", srv.Status.Health)
				assert.Equal(t, "running", srv.Status.State)
				assert.Equal(t, "normal operation", srv.Status.Info)
			},
		},
		{
			name:     "get with all includes",
			serverID: dbtools.FixtureNemo.ID,
			params: &fleetdbapi.ServerQueryParams{
				IncludeComponents: true,
				IncludeBMC:        true,
				IncludeStatus:     true,
			},
			verifyFn: func(t *testing.T, srv *fleetdbapi.Server) {
				assert.NotEmpty(t, srv.Components)
				assert.NotNil(t, srv.BMC)
				assert.NotNil(t, srv.Status)
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setupFn != nil {
				tt.setupFn(t)
			}

			var err error
			var srv *fleetdbapi.Server
			if tt.serverID == "not-a-uuid" {
				srv, _, err = s.Client.GetServer(context.Background(), uuid.Nil, tt.params)
			} else {
				srv, _, err = s.Client.GetServer(context.Background(), uuid.MustParse(tt.serverID), tt.params)
			}

			if tt.errorMsg != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMsg)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, srv)
			if tt.verifyFn != nil {
				tt.verifyFn(t, srv)
			}
		})
	}
}

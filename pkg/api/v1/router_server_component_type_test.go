package fleetdbapi_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/metal-automata/fleetdb/internal/dbtools"
	fleetdbapi "github.com/metal-automata/fleetdb/pkg/api/v1"
)

func TestIntegrationCreateServerComponentType(t *testing.T) {
	s := serverTest(t)

	realClientTests(t, func(ctx context.Context, authToken string, _ int, expectError bool) error {
		s.Client.SetToken(authToken)

		hct := fleetdbapi.ServerComponentType{Name: "integration-test"}

		resp, err := s.Client.CreateServerComponentType(ctx, hct)
		if !expectError {
			require.NoError(t, err)
			assert.Equal(t, "integration-test", resp.Slug)
			assert.NotNil(t, resp)
			assert.NotNil(t, resp.Links.Self)
			assert.Equal(t, fmt.Sprintf("http://test.hollow.com/api/v1/server-component-types/%s", resp.Slug), resp.Links.Self.Href)
		}

		return err
	})
}

func TestIntegrationListServerComponentTypes(t *testing.T) {
	s := serverTest(t)

	realClientTests(t, func(ctx context.Context, authToken string, _ int, expectError bool) error {
		s.Client.SetToken(authToken)

		r, resp, err := s.Client.ListServerComponentTypes(ctx, nil)
		if !expectError {
			require.NoError(t, err)
			assert.GreaterOrEqual(t, len(r), 1)
			assert.NotNil(t, r.ByName(dbtools.FixtureComponentTypeSlugMap["fins"].Name))
			assert.NotNil(t, resp)
			assert.NotNil(t, resp.Links.Self)
		}

		return err
	})
}

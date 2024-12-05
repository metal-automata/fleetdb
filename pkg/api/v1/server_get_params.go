package fleetdbapi

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/metal-automata/fleetdb/internal/models"
)

const (
	// server get query string include params
	includeServerBMC    = "s.bmc"
	includeServerStatus = "s.status"
	includeComponents   = "s.components"
)

// ServerGetParams allows you to filter server results and specify what related data to include.
type ServerGetParams struct {
	// include BMC attributes, including credentials
	IncludeBMC bool
	// include server status attributes
	IncludeStatus bool
	// include components
	IncludeComponents bool
	// Component include parameters
	//
	// requires Components set to true
	ComponentParams *ServerComponentGetParams
}

// encode converts the params into URL query parameters
func (p *ServerGetParams) encode() string {
	if p == nil {
		return ""
	}

	values := url.Values{}
	includes := []string{}

	// Handle BMC include
	if p.IncludeBMC {
		includes = append(includes, includeServerBMC)
	}

	// Handle Status include
	if p.IncludeStatus {
		includes = append(includes, includeServerStatus)
	}

	if p.IncludeComponents {
		includes = append(includes, includeComponents)

		// Add component includes if they exist
		if p.ComponentParams != nil {
			p.ComponentParams.encode()
		}
	}

	// Set includes if we have any
	if len(includes) > 0 {
		values.Set("include", strings.Join(includes, ","))
	}

	return values.Encode()
}

// Decode parses URL query parameters into ServerGetParams
func (p *ServerGetParams) decode(values url.Values) {
	// Parse includes
	includes := values.Get("include")
	if includes == "" {
		return
	}

	// Split includes and process each one
	for _, include := range splitAndTrim(includes) {
		switch {
		case include == includeServerBMC:
			p.IncludeBMC = true

		case include == includeServerStatus:
			p.IncludeStatus = true

		case include == includeComponents:
			p.IncludeComponents = true
			if p.ComponentParams == nil {
				p.ComponentParams = &ServerComponentGetParams{}
			}

		case strings.HasPrefix(include, "c."):
			// Initialize ComponentParams if needed
			if p.ComponentParams == nil {
				p.ComponentParams = &ServerComponentGetParams{}
			}

			p.ComponentParams.decodeIncludePart(include)
		}
	}
}

// Helper function to split comma-separated values and trim whitespace
func splitAndTrim(s string) []string {
	parts := strings.Split(s, ",")
	result := make([]string, 0, len(parts))

	for _, p := range parts {
		if trimmed := strings.TrimSpace(p); trimmed != "" {
			result = append(result, trimmed)
		}
	}

	return result
}

// setQuery sets URL query parameters based on the params
// setQuery implements the queryParams interface
func (p *ServerGetParams) setQuery(q url.Values) {
	if p == nil {
		return
	}

	if p.encode() == "" {
		return
	}

	// Parse the encoded string into values
	values, err := url.ParseQuery(p.encode())
	if err != nil {
		return
	}

	// Copy all values to the provided query object
	for key, vals := range values {
		for _, val := range vals {
			q.Add(key, val)
		}
	}
}

// queryMods returns query modifiers for SQL queries
func (p *ServerGetParams) queryMods(serverID string) []qm.QueryMod {
	mods := []qm.QueryMod{
		qm.Where(models.ServerTableColumns.ID+"=?", serverID),
		qm.WithDeleted(),
	}

	// Add server status query mods if requested
	if p.IncludeStatus {
		mods = append(mods,
			// join server status
			qm.InnerJoin(
				fmt.Sprintf(
					"%s on %s = %s",
					models.TableNames.ServerStatus,
					models.ServerTableColumns.ID,
					models.ServerStatusTableColumns.ServerID,
				),
			),
			// Load relationship in db model struct field R
			qm.Load(models.ServerRels.ServerStatus),
		)
	}

	// Add server BMC query mods if requested
	if p.IncludeBMC {
		mods = append(mods,
			// join server BMC
			qm.InnerJoin(
				fmt.Sprintf(
					"%s on %s = %s",
					models.TableNames.ServerBMCS,
					models.ServerTableColumns.ID,
					models.ServerBMCTableColumns.ServerID,
				),
			),
			qm.Load(models.ServerRels.ServerBMC),

			// join credentials to server
			qm.InnerJoin(
				fmt.Sprintf(
					"%s on %s = %s",
					models.TableNames.ServerCredentials,
					models.ServerCredentialTableColumns.ServerID,
					models.ServerTableColumns.ID,
				),
			),
			qm.Load(models.ServerRels.ServerCredentials),

			// join credentials with type
			qm.InnerJoin(
				fmt.Sprintf("%s as t on %s = %s",
					models.TableNames.ServerCredentialTypes,
					"t.id",
					models.ServerCredentialTableColumns.ServerCredentialTypeID,
				),
			),
			qm.Where(fmt.Sprintf("t.%s=?", models.ServerCredentialTypeColumns.Slug), "bmc"),
			// Load relationship in db model struct field R
			// qm.Load(models.ServerCredentialRels.ServerCredentialType),
		)
	}

	if p.IncludeComponents {
		mods = append(mods,
			qm.InnerJoin(
				fmt.Sprintf(
					"%s on %s = %s",
					models.TableNames.ServerComponents,
					models.ServerTableColumns.ID,
					models.ServerComponentTableColumns.ServerID,
				),
			),
			qm.Load(models.ServerRels.ServerComponents),
		)

		// Add component query mods if component params are present
		if p.ComponentParams != nil {
			mods = append(mods, p.ComponentParams.queryMods(false)...)
		}
	}

	return mods
}

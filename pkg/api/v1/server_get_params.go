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
//
// TODO: move include params into its own type?
// Generalize IncludeParams into its own type
// Have Servers implement a Includable() method that returns the various include feilds - similar to the FilterTarget interface
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

	// server attribute include parameters
	serverIncludes := []string{}

	// BMC information to be included
	if p.IncludeBMC {
		serverIncludes = append(serverIncludes, includeServerBMC)
	}

	// Server status to be included
	if p.IncludeStatus {
		serverIncludes = append(serverIncludes, includeServerStatus)
	}

	// url values to be encoded as a query string
	urlValues := url.Values{}

	// Component information to be included
	if p.IncludeComponents {
		serverIncludes = append(serverIncludes, includeComponents)

		// Component attributes to be included
		if p.ComponentParams != nil {
			p.ComponentParams.setQuery(urlValues)
		}
	}

	if len(serverIncludes) > 0 {
		// ComponentParams.setQuery may have set "include", prefix server includes for readability
		includeVals, exists := urlValues["include"]
		if exists {
			finalVals := strings.Join(serverIncludes, ",") + "," + includeVals[0]
			urlValues.Set("include", finalVals)
		} else {
			urlValues.Set("include", strings.Join(serverIncludes, ","))
		}
	}

	return urlValues.Encode()
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

	encoded := p.encode()
	if encoded == "" {
		return
	}

	// Parse the encoded string into values
	values, err := url.ParseQuery(encoded)
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
		qm.InnerJoin(
			fmt.Sprintf(
				"%s on %s = %s",
				models.TableNames.HardwareVendors,
				models.ServerTableColumns.VendorID,
				models.HardwareVendorTableColumns.ID,
			),
		),

		qm.Load(models.ServerRels.Vendor),

		qm.InnerJoin(
			fmt.Sprintf(
				"%s on %s = %s",
				models.TableNames.HardwareModels,
				models.ServerTableColumns.ModelID,
				models.HardwareModelTableColumns.ID,
			),
		),

		qm.Load(models.ServerRels.Model),
	}

	// Add server components if required
	//
	// Note: for server component attributes to be included, use r.componentsByServer()
	if p.IncludeComponents {
		mods = append(mods, qm.Load(models.ServerRels.ServerComponents))
	}

	// Add server status query mods if requested
	if p.IncludeStatus {
		mods = append(mods,
			// left join server status
			qm.LeftOuterJoin(
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
			// inner join server BMC - at the moment we expect all servers to have a bmc
			qm.InnerJoin(
				fmt.Sprintf(
					"%s on %s = %s",
					models.TableNames.ServerBMCS,
					models.ServerTableColumns.ID,
					models.ServerBMCTableColumns.ServerID,
				),
			),
			qm.Load(models.ServerRels.ServerBMC),

			// inner join credentials to server - expect credentials
			qm.InnerJoin(
				fmt.Sprintf(
					"%s on %s = %s",
					models.TableNames.ServerCredentials,
					models.ServerCredentialTableColumns.ServerID,
					models.ServerTableColumns.ID,
				),
			),
			qm.Load(models.ServerRels.ServerCredentials),

			// inner join credentials with type - expect cred type
			qm.InnerJoin(
				fmt.Sprintf("%s as t on %s = %s",
					models.TableNames.ServerCredentialTypes,
					"t.id",
					models.ServerCredentialTableColumns.ServerCredentialTypeID,
				),
			),
			qm.Where(fmt.Sprintf("t.%s=?", models.ServerCredentialTypeColumns.Slug), "bmc"),
		)
	}

	return mods
}

package fleetdbapi

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/metal-automata/fleetdb/internal/models"
)

const (
	serverInclude = "server_include"
	// server get query string serverInclude params
	includeServerBMC    = "bmc"
	includeServerStatus = "status"
	includeComponents   = "components"
)

// ServerQueryParams allows you to filter server results and specify what related data to serverInclude.
//
// TODO: move serverInclude params into its own type?
// Generalize IncludeParams into its own type
// Have Servers implement a Includable() method that returns the various serverInclude feilds - similar to the FilterTarget interface
type ServerQueryParams struct {
	// serverInclude BMC attributes, including credentials
	IncludeBMC bool
	// serverInclude server status attributes
	IncludeStatus bool
	// serverInclude components
	IncludeComponents bool
	// Component serverInclude parameters
	//
	// requires Components set to true
	ComponentParams *ServerComponentGetParams
	// FilterParams applies to queries listing more than one server
	FilterParams *FilterParams
	// ParginationParams applies to queries listing more than one server
	PaginationParams *PaginationParams
}

func (p *ServerQueryParams) toURLValues() url.Values {
	// url values from query params
	urlValues := url.Values{}

	if p == nil {
		return urlValues
	}

	addIncludeValue := func(v string) {
		_, exists := urlValues[serverInclude]
		if !exists {
			urlValues.Add(serverInclude, "")
		}

		// this method initializes url.Values and we expect there be only one value
		existing := urlValues[serverInclude][0]
		if existing == "" {
			urlValues[serverInclude][0] = v
		} else {
			join := []string{existing, v}
			urlValues[serverInclude][0] = strings.Join(join, ",")
		}
	}

	// BMC information to be included
	if p.IncludeBMC {
		addIncludeValue(includeServerBMC)
	}

	// Server status to be included
	if p.IncludeStatus {
		addIncludeValue(includeServerStatus)
	}

	// Component information to be included
	if p.IncludeComponents {
		addIncludeValue(includeComponents)

		// Component attributes to be included
		if p.ComponentParams != nil {
			componentUrlValues := p.ComponentParams.toURLValues()
			if componentUrlValues.Has(componentInclude) {
				urlValues.Set(componentInclude, componentUrlValues.Get(componentInclude))
			}
		}
	}

	return urlValues
}

func (p *ServerQueryParams) fromURLValues(values url.Values) error {
	if p == nil {
		return nil
	}

	// decode filter params if any
	if containsFilterParams(values) {
		p.FilterParams = &FilterParams{Target: &Server{}}
		p.FilterParams.fromURLValues(values)
	}

	// decode pagination params if any
	if containsPaginationParams(values) {
		paginationParams, err := parsePaginationURLQuery(values)
		if err != nil {
			return err
		}

		p.PaginationParams = &paginationParams
	}

	// Parse includes
	if !values.Has(serverInclude) {
		return nil
	}

	// Decode serverInclude params
	//
	// Split includes and process each one
	for _, include := range splitAndTrim(values.Get(serverInclude)) {
		switch {
		case include == includeServerBMC:
			p.IncludeBMC = true

		case include == includeServerStatus:
			p.IncludeStatus = true

		case include == includeComponents:
			p.IncludeComponents = true
			if p.ComponentParams == nil {
				p.ComponentParams = &ServerComponentGetParams{}
				p.ComponentParams.fromURLValues(values)
			}
		}
	}

	return nil
}

func containsFilterParams(values url.Values) bool {
	// TODO: add a server:: component:: prefix for these filter query keys
	// to differentiate between server and component query keys going ahead
	// and this match can be unique
	srv := &Server{}

	for _, s := range srv.FilterableColumnNames() {
		for key := range values {
			if strings.HasPrefix(key, s) {
				return true
			}
		}
	}

	return false
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
// it updates the given url.Values object with ServerQueryParameters
//
// setQuery implements the queryParams interface
func (p *ServerQueryParams) setQuery(q url.Values) {
	if p == nil {
		return
	}

	// Copy all values to the provided query object
	for key, vals := range p.toURLValues() {
		for _, val := range vals {
			q.Add(key, val)
		}
	}

	// Set filter params if any
	if p.FilterParams != nil {
		p.FilterParams.setQuery(q)
	}

	// Set pagination params if any
	if p.PaginationParams != nil {
		p.PaginationParams.setQuery(q)
	}
}

// queryMods returns query modifiers for SQL queries
func (p *ServerQueryParams) queryMods() []qm.QueryMod {
	mods := []qm.QueryMod{

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

	// Include filter mods for when its a server listing
	if p.FilterParams != nil {
		mods = append(mods, p.FilterParams.queryMods("servers")...)
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

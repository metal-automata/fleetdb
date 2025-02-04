package fleetdbapi

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/metal-automata/fleetdb/internal/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

const (
	componentInclude = "component_include"
	// component get query string include params
	includeComponentCaps             = "capabilities"
	includeComponentInstalledFw      = "installed_firmware"
	includeComponentStatus           = "status"
	includeComponentMetadataNSPrefix = "metadata_ns__"
)

// Set fields in this struct to for additional data to be included in the response
type ServerComponentGetParams struct {
	InstalledFirmware bool     // include installed firmware
	Status            bool     // include status, health
	Capabilities      bool     // include capabilities
	Metadata          []string // include metadata identified by the namespace
	Pagination        *PaginationParams
}

func (p *ServerComponentGetParams) setQuery(q url.Values) {
	if p == nil {
		return
	}

	// Copy all values to the provided query object
	for key, vals := range p.toURLValues() {
		for _, val := range vals {
			q.Add(key, val)
		}
	}

	// Set pagination params if any
	if p.Pagination != nil {
		p.Pagination.setQuery(q)
	}
}

func (p *ServerComponentGetParams) fromURLValues(values url.Values) {
	includes := values.Get(componentInclude)
	if includes == "" {
		return
	}

	for _, value := range splitAndTrim(values.Get(componentInclude)) {
		switch {
		case value == includeComponentCaps:
			p.Capabilities = true
		case value == includeComponentInstalledFw:
			p.InstalledFirmware = true
		case value == includeComponentStatus:
			p.Status = true
		case strings.HasPrefix(value, includeComponentMetadataNSPrefix):
			if p.Metadata == nil {
				p.Metadata = []string{}
			}

			// split metadata_ns__foo.bar
			parts := strings.Split(value, includeComponentMetadataNSPrefix)
			if len(parts) == 2 {
				p.Metadata = append(p.Metadata, strings.TrimSpace(parts[1]))
			}
		}
	}
}

// Returns url.Values based on the parameters
func (p *ServerComponentGetParams) toURLValues() url.Values {
	// url values from query params
	urlValues := url.Values{}
	if p == nil {
		return urlValues
	}

	addIncludeValue := func(v string) {
		_, exists := urlValues[componentInclude]
		if !exists {
			urlValues.Add(componentInclude, "")
		}

		// this method initializes url.Values and we expect there be only one value
		existing := urlValues[componentInclude][0]
		if existing == "" {
			urlValues[componentInclude][0] = v
		} else {
			join := []string{existing, v}
			urlValues[componentInclude][0] = strings.Join(join, ",")
		}
	}

	if p.Capabilities {
		addIncludeValue(includeComponentCaps)
	}

	if p.InstalledFirmware {
		addIncludeValue(includeComponentInstalledFw)
	}

	if p.Status {
		addIncludeValue(includeComponentStatus)
	}

	// Add metadata with namespaces if present
	if len(p.Metadata) > 0 {
		for _, ns := range p.Metadata {
			ns = strings.TrimSpace(ns)
			if ns != "" {
				addIncludeValue(includeComponentMetadataNSPrefix + ns)
			}
		}
	}

	return urlValues
}

// returns query mods based get parameters
func (p *ServerComponentGetParams) queryMods(joinComponentTypeIDs bool) []qm.QueryMod {
	//	TODO: make mods reusable
	scqm := &serverComponentQueryMods{}

	mods := []qm.QueryMod{}

	// join server component on server component types
	if joinComponentTypeIDs {
		mods = append(mods, scqm.types()...)
	}

	// join server components on installed firmware
	if p.InstalledFirmware {
		mods = append(mods, scqm.installedFirmware()...)
	}

	// join server components on status
	if p.Status {
		mods = append(mods, scqm.status()...)
	}

	// join server components on capabilities
	if p.Capabilities {
		mods = append(mods, scqm.capabilities()...)
	}

	// join server components on metadata
	if len(p.Metadata) > 0 {
		mods = append(mods, scqm.metadata(p.Metadata)...)
	}

	return mods
}

// struct holds methods to return JOIN query mods on the server components table
type serverComponentQueryMods struct{}

func (s *serverComponentQueryMods) types() []qm.QueryMod {
	return []qm.QueryMod{
		qm.InnerJoin(
			fmt.Sprintf(
				"%s on %s = %s",
				models.TableNames.ServerComponentTypes,
				models.ServerComponentTableColumns.ServerComponentTypeID,
				models.ServerComponentTypeTableColumns.ID,
			),
		),
		// Load relationship in db model struct field R
		qm.Load(models.ServerComponentRels.ServerComponentType),
	}
}

func (s *serverComponentQueryMods) installedFirmware() []qm.QueryMod {
	return []qm.QueryMod{
		qm.LeftOuterJoin(
			fmt.Sprintf(
				"%s on %s = %s",
				models.TableNames.InstalledFirmware,
				models.ServerComponentTableColumns.ID,
				models.InstalledFirmwareTableColumns.ServerComponentID,
			),
		),
		// Load relationship in db model struct field R
		qm.Load(models.ServerComponentRels.InstalledFirmware),
	}
}

func (s *serverComponentQueryMods) status() []qm.QueryMod {
	return []qm.QueryMod{
		// join server components on status
		qm.LeftOuterJoin(
			fmt.Sprintf(
				"%s on %s = %s",
				models.TableNames.ComponentStatus,
				models.ServerComponentTableColumns.ID,
				models.ComponentStatusTableColumns.ServerComponentID,
			),
		),
		// Load relationship in db model struct field R
		qm.Load(models.ServerComponentRels.ComponentStatus),
	}
}

func (s *serverComponentQueryMods) metadata(namespaces []string) []qm.QueryMod {
	whereMods := []qm.QueryMod{}
	for _, ns := range namespaces {
		whereMods = append(whereMods,
			qm.Where(models.ComponentMetadatumTableColumns.Namespace+" =?", ns),
		)
	}

	// This OR clause is included so as to allow all components to be listed,
	// without this clause we end up with the below query that limits the components returned
	// since not all of them would have metadata
	//
	// LEFT JOIN component_metadata on server_components.id = component_metadata.server_component_id
	// WHERE (component_metadata.namespace ='metadata.generic') AND ("server_components"."server_id"='c75dc8ae...');
	//
	// The other option would be to have nested queries but that doesn't mix well with query mods.
	// https://github.com/volatiletech/sqlboiler/issues/581
	whereMods = append(whereMods, qm.Or(fmt.Sprintf("%s is NULL", models.ComponentMetadatumTableColumns.Namespace)))

	cmods := []qm.QueryMod{
		// join server components on metadata
		qm.LeftOuterJoin(
			fmt.Sprintf(
				"%s on %s = %s",
				models.TableNames.ComponentMetadata,
				models.ServerComponentTableColumns.ID,
				models.ComponentMetadatumTableColumns.ServerComponentID,
			),
		),
		// Load relationship in db model struct field R
		qm.Load(models.ServerComponentRels.ComponentMetadata),
	}

	cmods = append(cmods, whereMods...)
	return cmods
}

func (s *serverComponentQueryMods) capabilities() []qm.QueryMod {
	return []qm.QueryMod{
		// join server components on capabilities
		qm.LeftOuterJoin(
			fmt.Sprintf(
				"%s on %s = %s",
				models.TableNames.ComponentCapabilities,
				models.ServerComponentTableColumns.ID,
				models.ComponentCapabilityTableColumns.ServerComponentID,
			),
		),
		// Load relationship in db model struct field R
		qm.Load(models.ServerComponentRels.ComponentCapabilities),
	}
}

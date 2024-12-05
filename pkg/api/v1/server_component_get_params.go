package fleetdbapi

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"

	"github.com/metal-automata/fleetdb/internal/models"
)

const (
	// component get query string include params
	includeComponentCaps        = "c.capabilities"
	includeComponentInstalledFw = "c.installed_firmware"
	includeComponentStatus      = "c.status"
	includeComponentMetadata    = "c.metadata"
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

	parsedValues, err := url.ParseQuery(p.encode())
	if err != nil {
		zap.L().Error("query params parse error", zap.Error(err))
		return
	}

	for k, v := range parsedValues {
		if len(v) == 0 {
			q.Set(k, "")
			continue
		}
		if len(v) > 1 {
			zap.L().Error(
				"expected single value for query param, got multiple",
				zap.Int("val", len(v)),
			)
			return
		}

		q.Set(k, v[0])
	}

	// url.ParseQuery()
	if p.Pagination != nil {
		p.Pagination.setQuery(q)
	}
}

// decode parses the include query parameter for server components
// Handles formats like: ?include=capabilities,firmware,status,metadata{namespace1,namespace2}
func (p *ServerComponentGetParams) decode(values url.Values) {
	includeParam := values.Get("include")
	if includeParam == "" {
		return
	}

	// Split by comma, but handle the special case of metadata{...}
	var parts []string
	var currentPart strings.Builder
	inBraces := false

	for _, char := range includeParam {
		switch char {
		case '{':
			inBraces = true
			currentPart.WriteRune(char)
		case '}':
			inBraces = false
			currentPart.WriteRune(char)
		case ',':
			if inBraces {
				currentPart.WriteRune(char)
			} else {
				parts = append(parts, strings.TrimSpace(currentPart.String()))
				currentPart.Reset()
			}
		default:
			currentPart.WriteRune(char)
		}
	}

	// Add the last part if not empty
	if currentPart.Len() > 0 {
		parts = append(parts, strings.TrimSpace(currentPart.String()))
	}

	// Process each part
	for _, part := range parts {
		p.decodeIncludePart(part)
	}
}

func (p *ServerComponentGetParams) decodeIncludePart(part string) {
	switch {
	case part == includeComponentCaps:
		p.Capabilities = true
	case part == includeComponentInstalledFw:
		p.InstalledFirmware = true
	case part == includeComponentStatus:
		p.Status = true
	case strings.HasPrefix(part, includeComponentMetadata+"{"):
		// Extract namespaces from metadata{ns1,ns2}
		ns := strings.TrimPrefix(part, includeComponentMetadata+"{")
		ns = strings.TrimSuffix(ns, "}")
		if ns != "" {
			namespaces := strings.Split(ns, ",")
			p.Metadata = make([]string, len(namespaces))
			for i, namespace := range namespaces {
				p.Metadata[i] = strings.TrimSpace(namespace)
			}
		}
	}
}

// encode converts ServerComponentGetParams into URL query parameters
func (p *ServerComponentGetParams) encode() string {
	if p == nil {
		return ""
	}

	var includes []string

	// Add simple boolean flags
	if p.Capabilities {
		includes = append(includes, includeComponentCaps)
	}
	if p.InstalledFirmware {
		includes = append(includes, includeComponentInstalledFw)
	}
	if p.Status {
		includes = append(includes, includeComponentStatus)
	}

	// Add metadata with namespaces if present
	if len(p.Metadata) > 0 {
		var namespaces []string
		for _, ns := range p.Metadata {
			if ns != "" {
				namespaces = append(namespaces, ns)
			}
		}
		if len(namespaces) > 0 {
			includes = append(includes, fmt.Sprintf("%s{%s}", includeComponentMetadata, strings.Join(namespaces, ",")))
		}
	}

	// If no includes, return empty string
	if len(includes) == 0 {
		return ""
	}

	// Join all includes with commas
	return "include=" + strings.Join(includes, ",")
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

	// join server components on metadaa
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
		qm.InnerJoin(
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
		qm.InnerJoin(
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

	cmods := []qm.QueryMod{
		// join server components on metadata
		qm.InnerJoin(
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
		qm.InnerJoin(
			fmt.Sprintf(
				"%s on %s = %s",
				models.TableNames.ComponentCapabilities,
				models.ServerComponentTableColumns.ID,
				models.ComponentCapabilityColumns.ServerComponentID,
			),
		),
		// Load relationship in db model struct field R
		qm.Load(models.ServerComponentRels.ComponentCapabilities),
	}
}

package fleetdbapi

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type ServerListParams struct {
	// TODO: add include params
	FilterParams     *FilterParams
	PaginationParams *PaginationParams
}

// implements the queryParams interface for the client
//
// TODO: update queryParams interface method to return error
func (p *ServerListParams) setQuery(q url.Values) {
	p.FilterParams.setQuery(q)
	p.PaginationParams.setQuery(q)
}

// parse server list parameters from the incoming request into the ServerListParams struct
func (p *ServerListParams) decode(values url.Values) error {
	p.FilterParams = &FilterParams{Target: &Server{}}
	p.FilterParams.decode(values)

	paginationParams, err := parsePaginationURLQuery(values)
	if err != nil {
		return err
	}

	p.PaginationParams = &paginationParams

	return nil
}

// returns queryMods based on FilterParams
func (p *ServerListParams) queryMods() []qm.QueryMod {
	mods := p.PaginationParams.queryMods()

	tableName := "servers"

	asAnySlice := func(s string) []interface{} {
		return []interface{}{s}
	}

	var currAttribute, prevAttribute string
	for idx, filter := range p.FilterParams.Filters {
		// vars to keep track of the filter being processed
		// for logical operations
		prevAttribute = currAttribute
		currAttribute = filter.Attribute

		var clause string
		args := []interface{}{}

		switch filter.ComparisonOperator {
		case ComparisonOpEqual:
			clause = fmt.Sprintf("%s.%s = ?", tableName, filter.Attribute)
			args = asAnySlice(filter.Value)

		case ComparisonOpNotEqual:
			clause = fmt.Sprintf("%s.%s != ?", tableName, filter.Attribute)
			args = asAnySlice(filter.Value)

		case ComparisonOpStartsWith:
			if filter.Modifier == ModifierCaseInsensitive {
				clause = fmt.Sprintf("LOWER(%s.%s) LIKE ?", tableName, filter.Attribute)
				args = asAnySlice(strings.ToLower(filter.Value + "%"))
			} else {
				clause = fmt.Sprintf("%s.%s LIKE ?", tableName, filter.Attribute)
				args = asAnySlice(filter.Value + "%")
			}

		case ComparisonOpEndsWith:
			if filter.Modifier == ModifierCaseInsensitive {
				clause = fmt.Sprintf("LOWER(%s.%s) LIKE ?", tableName, filter.Attribute)
				args = asAnySlice(strings.ToLower("%" + filter.Value))
			} else {
				clause = fmt.Sprintf("%s.%s LIKE ?", tableName, filter.Attribute)
				args = asAnySlice("%" + filter.Value)
			}

		case ComparisonOpContains:
			if filter.Modifier == ModifierCaseInsensitive {
				clause = fmt.Sprintf("LOWER(%s.%s) LIKE ?", tableName, filter.Attribute)
				args = asAnySlice(strings.ToLower("%" + filter.Value + "%"))
			} else {
				clause = fmt.Sprintf("%s.%s LIKE ?", tableName, filter.Attribute)
				args = asAnySlice("%" + filter.Value + "%")
			}
		}

		var mod qm.QueryMod

		// Apply logical operator on filters on the second until the second last filter
		if idx > 0 && idx <= len(p.FilterParams.Filters)-1 {
			op := p.FilterParams.LogicalOperatorFor(prevAttribute, currAttribute)
			switch op {
			case LogicalOpOr:
				mod = qm.Or(clause, args...)
			default:
				mod = qm.Where(clause, args...)
			}
		} else {
			mod = qm.Where(clause, args...)
		}

		mods = append(mods, mod)
	}

	// Include query mods
	// Filter query mods
	// Pagination query mods
	return mods
}

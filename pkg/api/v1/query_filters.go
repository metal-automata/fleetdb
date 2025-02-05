package fleetdbapi

import (
	"fmt"
	"net/url"
	"slices"
	"strings"

	"github.com/pkg/errors"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type ComparisonOperator string
type LogicalOperator string
type Modifier string

// A type that implements the FilterTarget is required to be passed into FilterParams
type FilterTarget interface {
	// FilterableColumnNames is the list of column names a client is allowed to create filters on
	FilterableColumnNames() []string
}

const (
	// comparison operators
	ComparisonOpEqual       ComparisonOperator = "eq"
	ComparisonOpNotEqual    ComparisonOperator = "ne"
	ComparisonOpGreaterThan ComparisonOperator = "gt"
	ComparisonOpLessThan    ComparisonOperator = "lt"
	ComparisonOpStartsWith  ComparisonOperator = "sw"
	ComparisonOpEndsWith    ComparisonOperator = "ew"
	ComparisonOpContains    ComparisonOperator = "contains"

	// logical operators
	LogicalOperation                 = "op"
	LogicalOpAnd     LogicalOperator = "and"
	LogicalOpNot     LogicalOperator = "not"
	LogicalOpOr      LogicalOperator = "or"

	// separator between ComparisonOperator, modifier and logical operator
	separator = "__"

	// modifiers
	ModifierCaseInsensitive Modifier = "cin"
)

type FilterParams struct {
	// type that implements the FilterTarget interface
	Target FilterTarget
	// Filters to apply
	Filters []Filter
	// Apply a logical operation on the given Filters
	LogicalOperation []string // odd elements are FilterableColumns, even are operators
}

// Helper method that returns the logical operator specified for the given two attributes
// method assumes the LogicalOperation is validated
func (s *FilterParams) LogicalOperatorFor(a, b string) LogicalOperator {
	if len(s.LogicalOperation) == 0 {
		return ""
	}

	var op LogicalOperator
	found := []string{}
	for idx, item := range s.LogicalOperation {
		if a == item || b == item {
			found = append(found, item)
		}

		if idx%2 != 0 {
			op = LogicalOperator(item)
		}

		if len(found) == 2 {
			break
		}
	}

	return op
}

// Filter represents a query parameter
type Filter struct {
	Attribute          string             // the left hand side parameter
	ComparisonOperator ComparisonOperator // The comparison operator
	Modifier           Modifier           // The string match modifier
	Value              string             // The value
}

// implements the client queryParams interface
func (s *FilterParams) setQuery(existing url.Values) {
	if s == nil {
		return
	}

	// update current with filter parameter url Values
	for key, values := range s.toURLValues() {
		for _, val := range values {
			_, exists := existing[key]
			if !exists {
				existing.Set(key, val)
			} else {
				existing.Add(key, val)
			}
		}
	}
}

// Returns FilterParams encoded as url.Values
func (s *FilterParams) toURLValues() url.Values {
	urlValues := url.Values{}

	if s == nil {
		return urlValues
	}

	if s.Target == nil {
		return urlValues
	}

	if len(s.Filters) == 0 {
		return urlValues
	}

	valid := s.Target.FilterableColumnNames()
	attributeIsKnown := func(s string) bool {
		// verify its a known attribute
		k := strings.ToLower(s)
		return slices.Contains(valid, k)
	}

	for _, qp := range s.Filters {
		if !attributeIsKnown(qp.Attribute) {
			continue
		}

		key := fmt.Sprintf("%s__%s", qp.Attribute, qp.ComparisonOperator)
		if qp.Modifier != "" {
			key += fmt.Sprintf("__%s", qp.Modifier)
		}

		_, exists := urlValues[key]
		if exists {
			urlValues.Add(key, qp.Value)
		} else {
			urlValues.Set(key, qp.Value)
		}
	}

	// include logical operation
	if len(s.LogicalOperation) > 0 && len(s.LogicalOperation)%2 != 0 {
		key := "op"
		value := strings.Join(s.LogicalOperation, "__")
		urlValues.Set(key, value)
	}

	return urlValues
}

// Decode given url.Values into FilterParams
//
// nolint:gocyclo // cyclomatic complexity is high to keep the context in one place
func (s *FilterParams) fromURLValues(values url.Values) {
	if len(values) == 0 {
		return
	}

	if s.Target == nil {
		return
	}

	s.Filters = []Filter{}

	valid := s.Target.FilterableColumnNames()
	attributeIsKnown := func(s string) bool {
		// verify its a known attribute
		k := strings.ToLower(s)
		return slices.Contains(valid, k)
	}

	// slice of attributes in query parameters
	// used to prepare the logical operator further below
	inQueryAttributes := []string{}

	for k, vals := range values {
		if k == LogicalOperation {
			continue
		}

		for _, val := range vals {
			qp := Filter{}

			// split <attribute>__<operator>__<modifier> by separator
			parts := strings.Split(k, separator)
			if !attributeIsKnown(parts[0]) {
				// unknown attribute gets query parameter ignored
				goto next
			}

			if !slices.Contains(inQueryAttributes, parts[0]) {
				inQueryAttributes = append(inQueryAttributes, parts[0])
			}

			// no operator defaults to equals
			if len(parts) == 1 {
				qp.Attribute = k
				qp.ComparisonOperator = ComparisonOpEqual
			} else {
				for idx, part := range parts {
					switch idx {
					case 0:
						qp.Attribute = part
					case 1:
						if op, err := asCompOperator(part); err == nil {
							qp.ComparisonOperator = op
						}
					case 2:
						if mod, err := asModifier(part); err == nil {
							qp.Modifier = mod
						}
					}
				}
			}

			qp.Value = val
			s.Filters = append(s.Filters, qp)

		next:
		}
	}

	// Odd parts are attributes, even parts are operators.
	//
	// - If any part of the logical operator parameter is incorrect, the whole is ignored:
	//
	//  Logical operations will also be ignored if:
	//
	// - An operator trails at the end of the op string
	// - The operation includes attributes not part of the query params
	// - An even count of (attributes + operators) is specified
	//
	//	<attribute>__<operator>__<attribute>
	logicalOpVals := values.Get(LogicalOperation)

	operation := []string{}
	parts := strings.Split(logicalOpVals, separator)
	if len(parts)%2 == 0 {
		return
	}

	for idx, part := range parts {
		if idx%2 == 0 {
			if !attributeIsKnown(part) || !slices.Contains(inQueryAttributes, part) {
				return
			}

			operation = append(operation, part)
		} else {
			op, err := asLogiOperator(part)
			if err != nil {
				return
			}

			// trailing operator is ignored
			if idx == len(parts)-1 {
				continue
			}

			operation = append(operation, string(op))
		}
	}

	s.LogicalOperation = operation
}

func asCompOperator(o string) (ComparisonOperator, error) {
	err := errors.New("unknown comparison operator")

	all := []ComparisonOperator{
		ComparisonOpEqual,
		ComparisonOpNotEqual,
		ComparisonOpGreaterThan,
		ComparisonOpLessThan,
		ComparisonOpStartsWith,
		ComparisonOpEndsWith,
		ComparisonOpContains,
	}

	op := ComparisonOperator(o)
	if !slices.Contains(all, op) {
		return "", errors.Wrap(err, o)
	}

	return op, nil
}

func asModifier(m string) (Modifier, error) {
	err := errors.New("unknown modifier")

	mod := Modifier(m)
	// nolint:gocritic // switch stays
	switch mod {
	case ModifierCaseInsensitive:
		return ModifierCaseInsensitive, nil
	}

	return "", errors.Wrap(err, m)
}

func asLogiOperator(o string) (LogicalOperator, error) {
	err := errors.New("unknown logical operator")

	all := []LogicalOperator{
		LogicalOperation,
		LogicalOpAnd,
		LogicalOpNot,
		LogicalOpOr,
	}

	op := LogicalOperator(o)
	if !slices.Contains(all, op) {
		return "", errors.Wrap(err, o)
	}

	return op, nil
}

// Returns query mods based on the FilterParams
func (s *FilterParams) queryMods(tableName string) []qm.QueryMod {
	mods := []qm.QueryMod{}

	asAnySlice := func(s string) []interface{} { // Include query mods
		return []interface{}{s}
	}

	var currAttribute, prevAttribute string
	for idx, filter := range s.Filters {
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
		if idx > 0 && idx <= len(s.Filters)-1 {
			op := s.LogicalOperatorFor(prevAttribute, currAttribute)
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

	return mods
}

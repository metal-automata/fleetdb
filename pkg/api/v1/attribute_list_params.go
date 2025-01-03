package fleetdbapi

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// AttributeListParams allow you to filter the results based on attributes
type AttributeListParams struct {
	Namespace string
	Keys      []string
	Operator  OperatorComparitorType
	Value     string
	// OperatorLogicalType is used to define how this AttributeListParam value should be SQL queried
	// this value defaults to OperatorLogicalAND.
	AttributeOperator OperatorLogicalType
}

const pairSize = 2

func encodeAttributesListParams(alp []AttributeListParams, key string, q url.Values) {
	for _, ap := range alp {
		value := ap.Namespace

		if len(ap.Keys) != 0 && value != "" {
			value = fmt.Sprintf("%s~%s", value, strings.Join(ap.Keys, "."))

			if ap.Operator != "" && ap.Value != "" {
				value = fmt.Sprintf("%s~%s~%s", value, ap.Operator, ap.Value)
			}
		}

		if ap.AttributeOperator != "" {
			value += "~" + string(ap.AttributeOperator)
		}

		q.Add(key, value)
	}
}

func appendToQueryFirmwareSetsParams(params []AttributeListParams, keys, op, value string) []AttributeListParams {
	return append(params, AttributeListParams{
		Namespace: FirmwareSetAttributeNS,
		Keys:      []string{keys},
		Operator:  OperatorComparitorType(op),
		Value:     strings.ToLower(value),
	})
}

// Function to parse labels parameter into a map.
// TODO: may support or/and operators in the future.
// TODO: may want to check SQL injection?
func parseQueryFirmwareSetsLabels(labels string) map[string]string {
	result := make(map[string]string)
	pairs := strings.Split(labels, ",")

	for _, pair := range pairs {
		kv := strings.SplitN(pair, "=", pairSize)
		if len(kv) == pairSize {
			key := kv[0]
			value := kv[1]
			result[key] = value
		}
	}

	return result
}

func parseQueryAttributesListParams(c *gin.Context, key string) []AttributeListParams {
	alp := []AttributeListParams{}

	attrQueryParams := c.QueryArray(key)

	for _, p := range attrQueryParams {
		// format accepted
		// "ns~keys.dot.seperated~operation~value"
		// With attr OR operator: "ns~keys.dot.seperated~operation~value~or"
		// With attr AND operator: "ns~keys.dot.seperated~operation~value~and"
		parts := strings.Split(p, "~")

		param := AttributeListParams{
			Namespace: parts[0],
		}

		if len(parts) == 1 {
			alp = append(alp, param)
			continue
		}

		param.Keys = strings.Split(parts[1], ".")

		if len(parts) == 4 || len(parts) == 5 { // nolint
			switch o := (*OperatorComparitorType)(&parts[2]); *o {
			case OperatorComparitorEqual, OperatorComparitorLike, OperatorComparitorGreaterThan, OperatorComparitorLessThan:
				param.Operator = *o
				param.Value = parts[3]
			}

			// An attribute operator is only applicable when,
			// - Theres 5 parts in the attr param string when split on `~`.
			// - Theres multiple attribute query parameters defined.
			if len(parts) == 5 && len(attrQueryParams) > 1 {
				switch o := (*OperatorLogicalType)(&parts[4]); *o {
				case OperatorLogicalAND, OperatorLogicalOR:
					param.AttributeOperator = *o
				}
			}

			// if the like search doesn't contain any % add one at the end
			if param.Operator == OperatorComparitorLike && !strings.Contains(param.Value, "%") {
				param.Value += "%"
			}
		}

		alp = append(alp, param)
	}

	return alp
}

// queryMods converts the list params into sql conditions that can be added to
// sql queries
func (p *AttributeListParams) queryMods(tblName string) qm.QueryMod {
	nsMod := qm.Where(fmt.Sprintf("%s.namespace = ?", tblName), p.Namespace)

	values := []interface{}{}
	jsonPath := ""

	// If we only have a namespace and no keys we are limiting by namespace only
	if len(p.Keys) == 0 {
		return nsMod
	}

	for i, k := range p.Keys {
		if i > 0 {
			jsonPath += " , "
		}
		// the actual key is represented as a "?" this helps protect against SQL
		// injection since these strings are passed in by the user.
		jsonPath += "?"

		values = append(values, k)
	}

	where, values := p.setJSONBWhereClause(tblName, jsonPath, values)

	// namespace AND JSONB query as a query mod
	queryMods := []qm.QueryMod{nsMod, qm.And(where, values...)}

	// OR ( namespace AND JSONB query )
	if p.AttributeOperator == OperatorLogicalOR {
		return qm.Or2(qm.Expr(queryMods...))
	}

	// AND ( namespace AND JSONB query )
	return qm.Expr(queryMods...)
}

func (p *AttributeListParams) setJSONBWhereClause(tblName, jsonPath string, values []interface{}) (where string, items []interface{}) {
	switch p.Operator {
	case OperatorComparitorLessThan:
		values = append(values, p.Value)
		where = fmt.Sprintf("jsonb_extract_path_text(%s.data::JSONB, %s)::int < ?", tblName, jsonPath)
	case OperatorComparitorGreaterThan:
		values = append(values, p.Value)
		where = fmt.Sprintf("jsonb_extract_path_text(%s.data::JSONB, %s)::int > ?", tblName, jsonPath)
	case OperatorComparitorLike:
		values = append(values, p.Value)
		where = fmt.Sprintf("jsonb_extract_path_text(%s.data::JSONB, %s) LIKE ?", tblName, jsonPath)
	case OperatorComparitorEqual:
		values = append(values, p.Value)
		where = fmt.Sprintf("jsonb_extract_path_text(%s.data::JSONB, %s) = ?", tblName, jsonPath)
	default:
		// we only have keys so we just want to ensure the key is there
		where = fmt.Sprintf("%s.data::JSONB", tblName)

		if len(p.Keys) != 0 {
			for range p.Keys[0 : len(p.Keys)-1] {
				where += " -> ?"
			}

			// query is existing_where ? key
			where += " \\? ?"
		}
	}

	return where, values
}

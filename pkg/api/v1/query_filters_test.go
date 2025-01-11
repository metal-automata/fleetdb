package fleetdbapi

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFilterParamsLogicalOperationFor(t *testing.T) {
	testcases := []struct {
		name   string
		op     []string
		params [2]string
		expect LogicalOperator
	}{
		{
			name:   "one op",
			op:     []string{"serial", "and", "name"},
			params: [2]string{"serial", "name"},
			expect: LogicalOpAnd,
		},
		{
			name:   "two ops",
			op:     []string{"serial", "and", "name", "or", "vendor"},
			params: [2]string{"name", "vendor"},
			expect: LogicalOpOr,
		},
		{
			name:   "three ops",
			op:     []string{"serial", "and", "name", "or", "vendor", "not", "model"},
			params: [2]string{"vendor", "model"},
			expect: LogicalOpNot,
		},
		{
			name:   "no op",
			op:     []string{"serial"},
			expect: "",
		},
		{
			name:   "attributes with the same name - OR",
			op:     []string{"name", "or", "name"},
			params: [2]string{"name", "name"},
			expect: LogicalOpOr,
		},
	}
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			f := FilterParams{LogicalOperation: tt.op}
			got := f.LogicalOperatorFor(tt.params[0], tt.params[1])
			assert.Equal(t, tt.expect, got)
		})
	}
}

func TestSetQuery(t *testing.T) {
	filterTarget := &Server{}
	tests := []struct {
		name    string
		expect  string
		params  *FilterParams
		current url.Values
	}{
		{
			name: "Does not overwrite existing",
			current: url.Values{
				"existing": []string{"val1", "val2"},
			},
			expect: "existing=val1&existing=val2&op=serial__or__serial&serial__eq=123&serial__eq=456",
			params: &FilterParams{
				Target: filterTarget,
				Filters: []Filter{
					{
						Attribute:          "serial",
						ComparisonOperator: ComparisonOpEqual,
						Value:              "123",
					},
					{
						Attribute:          "serial",
						ComparisonOperator: ComparisonOpEqual,
						Value:              "456",
					},
				},
				LogicalOperation: []string{"serial", "or", "serial"},
			},
		},
		{
			name:    "No existing params",
			current: url.Values{},
			expect:  "serial__eq=123",
			params: &FilterParams{
				Target: filterTarget,
				Filters: []Filter{
					{
						Attribute:          "serial",
						ComparisonOperator: ComparisonOpEqual,
						Value:              "123",
					},
				},
			},
		},
		{
			name: "Different filter attributes",
			current: url.Values{
				"existing": []string{"val1"},
			},
			expect: "existing=val1&serial__eq=123&name__eq=foo",
			params: &FilterParams{
				Target: filterTarget,
				Filters: []Filter{
					{
						Attribute:          "serial",
						ComparisonOperator: ComparisonOpEqual,
						Value:              "123",
					},
					{
						Attribute:          "name",
						ComparisonOperator: ComparisonOpEqual,
						Value:              "foo",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.params.setQuery(tt.current)

			want, _ := url.ParseQuery(tt.expect)
			assert.Equal(t, tt.current, want)
		})
	}
}

func TestFilterParamsEncode(t *testing.T) {
	filterTarget := &Server{}
	tests := []struct {
		name   string
		expect string
		params *FilterParams
	}{
		{
			name:   "default to equals",
			expect: "serial__eq=123",
			params: &FilterParams{
				Target: filterTarget,
				Filters: []Filter{
					{
						Attribute:          "serial",
						ComparisonOperator: ComparisonOpEqual,
						Value:              "123",
					},
				},
			},
		},
		{
			name:   "nil filter Target returns nothing",
			expect: "",
			params: &FilterParams{
				Filters: []Filter{
					{
						Attribute:          "serial",
						ComparisonOperator: ComparisonOpEqual,
						Value:              "123",
					},
				},
			},
		},
		{
			name:   "unknown Attribute is ignored",
			expect: "serial__eq=123",
			params: &FilterParams{
				Target: filterTarget,
				Filters: []Filter{
					{
						Attribute:          "serial",
						ComparisonOperator: ComparisonOpEqual,
						Value:              "123",
					},
					{
						Attribute:          "baz",
						ComparisonOperator: ComparisonOpEqual,
						Value:              "999",
					},
				},
			},
		},
		{
			name:   "filter on same attribute - or",
			expect: "serial__eq=123&serial__eq=456&op=serial__or__serial",
			params: &FilterParams{
				Target: filterTarget,
				Filters: []Filter{
					{
						Attribute:          "serial",
						ComparisonOperator: ComparisonOpEqual,
						Value:              "123",
					},
					{
						Attribute:          "serial",
						ComparisonOperator: ComparisonOpEqual,
						Value:              "456",
					},
				},
				LogicalOperation: []string{"serial", "or", "serial"},
			},
		},
		{
			name:   "operator equals with logical op - and",
			expect: "serial__eq=123&name__eq=foobar&op=serial__and__name",
			params: &FilterParams{
				Target: filterTarget,
				Filters: []Filter{
					{
						Attribute:          "serial",
						ComparisonOperator: ComparisonOpEqual,
						Value:              "123",
					},
					{
						Attribute:          "name",
						ComparisonOperator: ComparisonOpEqual,
						Value:              "foobar",
					},
				},
				LogicalOperation: []string{"serial", "and", "name"},
			},
		},
		{
			name:   "operator equals with logical op - or",
			expect: "name__eq=foo&serial__eq=123&op=name__or__serial",
			params: &FilterParams{
				Target: filterTarget,
				Filters: []Filter{
					{
						Attribute:          "name",
						ComparisonOperator: ComparisonOpEqual,
						Value:              "foo",
					},
					{
						Attribute:          "serial",
						ComparisonOperator: ComparisonOpEqual,
						Value:              "123",
					},
				},
				LogicalOperation: []string{"name", "or", "serial"},
			},
		},
		{
			name:   "incorrect logical operation is ignored",
			expect: "name__eq=foo&serial__eq=123",
			params: &FilterParams{
				Target: filterTarget,
				Filters: []Filter{
					{
						Attribute:          "name",
						ComparisonOperator: ComparisonOpEqual,
						Value:              "foo",
					},
					{
						Attribute:          "serial",
						ComparisonOperator: ComparisonOpEqual,
						Value:              "123",
					},
				},
			},
		},
		{
			name:   "operator with modifier and logical op",
			expect: "name__eq__cin=foo&serial__sw=123&vendor__eq=foobar&op=name__and__serial__or__vendor",
			params: &FilterParams{
				Target: filterTarget,
				Filters: []Filter{
					{
						Attribute:          "name",
						ComparisonOperator: ComparisonOpEqual,
						Modifier:           ModifierCaseInsensitive,
						Value:              "foo",
					},
					{
						Attribute:          "serial",
						ComparisonOperator: ComparisonOpStartsWith,
						Value:              "123",
					},
					{
						Attribute:          "vendor",
						ComparisonOperator: ComparisonOpEqual,
						Value:              "foobar",
					},
				},
				LogicalOperation: []string{"name", "and", "serial", "or", "vendor"},
			},
		},

		{
			name:   "operator endswith",
			expect: "serial__ew=123",
			params: &FilterParams{
				Target: filterTarget,
				Filters: []Filter{
					{
						Attribute:          "serial",
						ComparisonOperator: ComparisonOpEndsWith,
						Value:              "123",
					},
				},
			},
		},
		{
			name:   "operator contains",
			expect: "serial__contains=123",
			params: &FilterParams{
				Target: filterTarget,
				Filters: []Filter{
					{
						Attribute:          "serial",
						ComparisonOperator: ComparisonOpContains,
						Value:              "123",
					},
				},
			},
		},
		{
			name:   "operator greater than",
			expect: "serial__gt=123",
			params: &FilterParams{
				Target: filterTarget,
				Filters: []Filter{
					{
						Attribute:          "serial",
						ComparisonOperator: ComparisonOpGreaterThan,
						Value:              "123",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.params.encode()
			assert.Equal(t, tt.expect, got)
		})
	}
}

func TestFilterParamsDecode(t *testing.T) {
	tests := []struct {
		name        string
		queryString string
		expect      *FilterParams
	}{
		{
			name:        "empty query",
			queryString: "",
			expect:      &FilterParams{},
		},
		{
			name:        "default to equals",
			queryString: "serial=123",
			expect: &FilterParams{
				Filters: []Filter{
					{
						Attribute:          "serial",
						ComparisonOperator: ComparisonOpEqual,
						Value:              "123",
					},
				},
			},
		},
		{
			name:        "operator equals",
			queryString: "serial__eq=123",
			expect: &FilterParams{
				Filters: []Filter{
					{
						Attribute:          "serial",
						ComparisonOperator: ComparisonOpEqual,
						Value:              "123",
					},
				},
			},
		},
		{
			name:        "unknown Attribute is ignored",
			queryString: "baz__eq=foo&serial__eq=123",
			expect: &FilterParams{
				Filters: []Filter{
					{
						Attribute:          "serial",
						ComparisonOperator: ComparisonOpEqual,
						Value:              "123",
					},
				},
			},
		},
		{
			name:        "operator equals with logical op - and",
			queryString: "serial__eq=123&name__eq=foobar&op=serial__and__name",
			expect: &FilterParams{
				Filters: []Filter{
					{
						Attribute:          "serial",
						ComparisonOperator: ComparisonOpEqual,
						Value:              "123",
					},
					{
						Attribute:          "name",
						ComparisonOperator: ComparisonOpEqual,
						Value:              "foobar",
					},
				},
				LogicalOperation: []string{"serial", "and", "name"},
			},
		},
		{
			name:        "operator equals with logical op - or",
			queryString: "name__eq=foo&serial__eq=123&op=name__or__serial",
			expect: &FilterParams{
				Filters: []Filter{
					{
						Attribute:          "name",
						ComparisonOperator: ComparisonOpEqual,
						Value:              "foo",
					},
					{
						Attribute:          "serial",
						ComparisonOperator: ComparisonOpEqual,
						Value:              "123",
					},
				},
				LogicalOperation: []string{"name", "or", "serial"},
			},
		},
		{
			name:        "filter on same attribute - or",
			queryString: "serial__eq=123&serial__eq=456&op=serial__or__serial",
			expect: &FilterParams{
				Filters: []Filter{
					{
						Attribute:          "serial",
						ComparisonOperator: ComparisonOpEqual,
						Value:              "123",
					},
					{
						Attribute:          "serial",
						ComparisonOperator: ComparisonOpEqual,
						Value:              "456",
					},
				},
				LogicalOperation: []string{"serial", "or", "serial"},
			},
		},
		{
			name:        "trailing logical operator is ignored",
			queryString: "name__eq=foo&serial__eq=123&op=name__or__serial__and",
			expect: &FilterParams{
				Filters: []Filter{
					{
						Attribute:          "name",
						ComparisonOperator: ComparisonOpEqual,
						Value:              "foo",
					},
					{
						Attribute:          "serial",
						ComparisonOperator: ComparisonOpEqual,
						Value:              "123",
					},
				},
			},
		},
		{
			name:        "incorrect Attribute in logical operation is ignored",
			queryString: "name__eq=foo&serial__eq=123&op=name__or__vendor",
			expect: &FilterParams{
				Filters: []Filter{
					{
						Attribute:          "name",
						ComparisonOperator: ComparisonOpEqual,
						Value:              "foo",
					},
					{
						Attribute:          "serial",
						ComparisonOperator: ComparisonOpEqual,
						Value:              "123",
					},
				},
			},
		},
		{
			name:        "odd count of parameters in the logical op param",
			queryString: "name__eq=foo&serial__eq=123&op=name__or",
			expect: &FilterParams{
				Filters: []Filter{
					{
						Attribute:          "name",
						ComparisonOperator: ComparisonOpEqual,
						Value:              "foo",
					},
					{
						Attribute:          "serial",
						ComparisonOperator: ComparisonOpEqual,
						Value:              "123",
					},
				},
			},
		},
		{
			name:        "operator with modifier and logical op",
			queryString: "name__eq__cin=foo&serial__sw=123&vendor__eq=foobar&op=name__and__serial__or__vendor",
			expect: &FilterParams{
				Filters: []Filter{
					{
						Attribute:          "name",
						ComparisonOperator: ComparisonOpEqual,
						Modifier:           ModifierCaseInsensitive,
						Value:              "foo",
					},
					{
						Attribute:          "serial",
						ComparisonOperator: ComparisonOpStartsWith,
						Value:              "123",
					},
					{
						Attribute:          "vendor",
						ComparisonOperator: ComparisonOpEqual,
						Value:              "foobar",
					},
				},
				LogicalOperation: []string{"name", "and", "serial", "or", "vendor"},
			},
		},
		{
			name:        "operator endswith",
			queryString: "serial__ew=123",
			expect: &FilterParams{
				Filters: []Filter{
					{
						Attribute:          "serial",
						ComparisonOperator: ComparisonOpEndsWith,
						Value:              "123",
					},
				},
			},
		},
		{
			name:        "operator contains",
			queryString: "serial__contains=123",
			expect: &FilterParams{
				Filters: []Filter{
					{
						Attribute:          "serial",
						ComparisonOperator: ComparisonOpContains,
						Value:              "123",
					},
				},
			},
		},
		{
			name:        "operator greater than",
			queryString: "serial__gt=123",
			expect: &FilterParams{
				Filters: []Filter{
					{
						Attribute:          "serial",
						ComparisonOperator: ComparisonOpGreaterThan,
						Value:              "123",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Values, err := url.ParseQuery(tt.queryString)
			require.NoError(t, err)

			var target *Server
			if tt.name != "nil target" {
				target = &Server{}
			}
			params := &FilterParams{Target: target}
			params.decode(Values)
			assert.ElementsMatch(t, tt.expect.Filters, params.Filters)
			assert.ElementsMatch(t, tt.expect.LogicalOperation, params.LogicalOperation)
		})
	}

}

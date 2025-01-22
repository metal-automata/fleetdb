package fleetdbapi

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerQueryParamsSetQuery(t *testing.T) {
	tests := []struct {
		name string
		// input gets converted to url values before being passed into setQuery
		input   string
		current *ServerQueryParams
		expect  string
	}{
		{
			name:    "empty params",
			input:   "",
			current: &ServerQueryParams{},
			expect:  "",
		},
		{
			name:  "pagingation params are included",
			input: "page=1&limit=2",
			current: &ServerQueryParams{
				IncludeBMC: true,
			},
			expect: "server_include=bmc&page=1&limit=2",
		},
		{
			name:  "pagingation and filter params are included",
			input: "page=1&limit=2&serial__eq=123",
			current: &ServerQueryParams{
				IncludeBMC: true,
			},
			expect: "server_include=bmc&serial__eq=123&page=1&limit=2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			urlValues, err := url.ParseQuery(tt.input)
			assert.Nil(t, err)

			// update urlValues with
			tt.current.setQuery(urlValues)

			expect, err := url.ParseQuery(tt.expect)
			assert.Nil(t, err)

			assert.Equal(t, expect, urlValues)
		})
	}
}

func TestServerQueryParamsToURLValues(t *testing.T) {
	tests := []struct {
		name   string
		input  *ServerQueryParams
		expect string
	}{
		{
			name:   "empty params",
			input:  &ServerQueryParams{},
			expect: "",
		},
		{
			name: "bmc only",
			input: &ServerQueryParams{
				IncludeBMC: true,
			},
			expect: "server_include=bmc",
		},
		{
			name: "status only",
			input: &ServerQueryParams{
				IncludeStatus: true,
			},
			expect: "server_include=status",
		},
		{
			name: "components only",
			input: &ServerQueryParams{
				IncludeComponents: true,
			},
			expect: "server_include=components",
		},
		{
			name: "server include with components include",
			input: &ServerQueryParams{
				IncludeComponents: true,
				ComponentParams: &ServerComponentGetParams{
					InstalledFirmware: true,
					Capabilities:      true,
					Status:            true,
					Metadata:          []string{"namespace1"},
				},
			},
			expect: "server_include=components&component_include=capabilities,installed_firmware,status,metadata_ns__namespace1",
		},
		{
			name: "bmc and status",
			input: &ServerQueryParams{
				IncludeBMC:    true,
				IncludeStatus: true,
			},
			expect: "server_include=bmc,status",
		},
		{
			name: "empty component params",
			input: &ServerQueryParams{
				IncludeBMC:      true,
				ComponentParams: &ServerComponentGetParams{},
			},
			expect: "server_include=bmc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expect, err := url.ParseQuery(tt.expect)
			assert.Nil(t, err)
			assert.Equal(t, expect, tt.input.toURLValues())
		})
	}
}

func TestServerQueryParamsFromURLValues(t *testing.T) {
	tests := []struct {
		name     string
		params   *ServerQueryParams
		expected string
	}{
		{
			name:     "empty params",
			params:   &ServerQueryParams{},
			expected: "",
		},
		{
			name: "bmc only",
			params: &ServerQueryParams{
				IncludeBMC: true,
			},
			expected: "server_include=bmc",
		},
		{
			name: "status only",
			params: &ServerQueryParams{
				IncludeStatus: true,
			},
			expected: "server_include=status",
		},
		{
			name: "components only",
			params: &ServerQueryParams{
				IncludeComponents: true,
			},
			expected: "server_include=components",
		},
		{
			name: "server include with components include",
			params: &ServerQueryParams{
				IncludeComponents: true,
				ComponentParams: &ServerComponentGetParams{
					InstalledFirmware: true,
					Capabilities:      true,
					Status:            true,
					Metadata:          []string{"namespace1"},
				},
			},
			expected: "server_include=components&component_include=capabilities,installed_firmware,status,metadata_ns__namespace1",
		},
		{
			name: "bmc and status",
			params: &ServerQueryParams{
				IncludeBMC:    true,
				IncludeStatus: true,
			},
			expected: "server_include=bmc,status",
		},
		{
			name: "empty component params",
			params: &ServerQueryParams{
				IncludeBMC:      true,
				ComponentParams: &ServerComponentGetParams{},
			},
			expected: "server_include=bmc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.params.toURLValues()

			// If we have an expected result, verify we can decode it back
			if tt.expected != "" {
				expect, err := url.ParseQuery(tt.expected)
				assert.Nil(t, err)
				assert.Equal(t, expect, result)
			} else {
				assert.Empty(t, result)
			}
		})
	}
}

func TestSplitAndTrim(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: []string{},
		},
		{
			name:     "single value",
			input:    "test",
			expected: []string{"test"},
		},
		{
			name:     "multiple values",
			input:    "one,two,three",
			expected: []string{"one", "two", "three"},
		},
		{
			name:     "with spaces",
			input:    " one , two , three ",
			expected: []string{"one", "two", "three"},
		},
		{
			name:     "with empty parts",
			input:    "one,,two,,three",
			expected: []string{"one", "two", "three"},
		},
		{
			name:     "with only spaces",
			input:    " , , ",
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := splitAndTrim(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

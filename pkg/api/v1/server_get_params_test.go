package fleetdbapi

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServerGetParamsEncode(t *testing.T) {
	tests := []struct {
		name     string
		params   *ServerGetParams
		expected string
	}{
		{
			name:     "nil params",
			params:   nil,
			expected: "",
		},
		{
			name:     "empty params",
			params:   &ServerGetParams{},
			expected: "",
		},
		{
			name: "bmc only",
			params: &ServerGetParams{
				IncludeBMC: true,
			},
			expected: "include=s.bmc",
		},
		{
			name: "status only",
			params: &ServerGetParams{
				IncludeStatus: true,
			},
			expected: "include=s.status",
		},
		{
			name: "components only",
			params: &ServerGetParams{
				IncludeComponents: true,
			},
			expected: "include=s.components",
		},
		{
			name: "components with attributes",
			params: &ServerGetParams{
				IncludeComponents: true,
				ComponentParams: &ServerComponentGetParams{
					InstalledFirmware: true,
					Capabilities:      true,
					Status:            true,
				},
			},
			expected: "include=s.components,c.capabilities,c.installed_firmware,c.status",
		},

		{
			name: "bmc and status",
			params: &ServerGetParams{
				IncludeBMC:    true,
				IncludeStatus: true,
			},
			expected: "include=s.bmc,s.status",
		},
		{
			name: "empty component params",
			params: &ServerGetParams{
				IncludeBMC:      true,
				ComponentParams: &ServerComponentGetParams{},
			},
			expected: "include=s.bmc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.params.encode()
			got, err := url.QueryUnescape(result)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, got)

			// If we have an expected result, verify we can decode it back
			if tt.expected != "" {
				values, err := url.ParseQuery(result)
				assert.NoError(t, err)

				decoded := &ServerGetParams{}
				decoded.decode(values)

				assert.Equal(t, tt.params.IncludeBMC, decoded.IncludeBMC)
				assert.Equal(t, tt.params.IncludeStatus, decoded.IncludeStatus)
			}
		})
	}
}

func TestServerGetParamsDecode(t *testing.T) {
	tests := []struct {
		name        string
		queryString string
		validate    func(*testing.T, *ServerGetParams)
	}{
		{
			name:        "empty query",
			queryString: "",
			validate: func(t *testing.T, params *ServerGetParams) {
				assert.False(t, params.IncludeBMC)
				assert.False(t, params.IncludeStatus)
				assert.Nil(t, params.ComponentParams)
			},
		},
		{
			name:        "include bmc",
			queryString: "include=s.bmc",
			validate: func(t *testing.T, params *ServerGetParams) {
				assert.True(t, params.IncludeBMC)
				assert.False(t, params.IncludeStatus)
				assert.Nil(t, params.ComponentParams)
			},
		},
		{
			name:        "include status",
			queryString: "include=s.status",
			validate: func(t *testing.T, params *ServerGetParams) {
				assert.False(t, params.IncludeBMC)
				assert.True(t, params.IncludeStatus)
				assert.Nil(t, params.ComponentParams)
			},
		},
		{
			name:        "include components",
			queryString: "include=s.components",
			validate: func(t *testing.T, params *ServerGetParams) {
				assert.False(t, params.IncludeBMC)
				assert.False(t, params.IncludeStatus)
				assert.True(t, params.IncludeComponents)
				assert.NotNil(t, params.ComponentParams)
			},
		},

		{
			name:        "include both bmc and status",
			queryString: "include=s.bmc,s.status",
			validate: func(t *testing.T, params *ServerGetParams) {
				assert.True(t, params.IncludeBMC)
				assert.True(t, params.IncludeStatus)
				assert.Nil(t, params.ComponentParams)
			},
		},
		{
			name:        "include components with status",
			queryString: "include=s.components,c.status",
			validate: func(t *testing.T, params *ServerGetParams) {
				assert.False(t, params.IncludeBMC)
				assert.False(t, params.IncludeStatus)
				assert.True(t, params.IncludeComponents)
				assert.NotNil(t, params.ComponentParams)
			},
		},
		{
			name:        "include everything",
			queryString: "include=s.bmc,s.status,c.status,s.components,c.firmware,",
			validate: func(t *testing.T, params *ServerGetParams) {
				assert.True(t, params.IncludeBMC)
				assert.True(t, params.IncludeStatus)
				assert.True(t, params.IncludeComponents)
				assert.NotNil(t, params.ComponentParams)
			},
		},
		{
			name:        "handles whitespace",
			queryString: "include= s.bmc , s.status, s.components , c.firmware ",
			validate: func(t *testing.T, params *ServerGetParams) {
				assert.True(t, params.IncludeBMC)
				assert.True(t, params.IncludeStatus)
				assert.True(t, params.IncludeComponents)
				assert.NotNil(t, params.ComponentParams)
			},
		},
		{
			name:        "ignores unknown includes",
			queryString: "include=s.bmc,unknown,s.status",
			validate: func(t *testing.T, params *ServerGetParams) {
				assert.True(t, params.IncludeBMC)
				assert.True(t, params.IncludeStatus)
				assert.Nil(t, params.ComponentParams)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			values, err := url.ParseQuery(tt.queryString)
			require.NoError(t, err)

			params := &ServerGetParams{}
			params.decode(values)
			tt.validate(t, params)

			// Verify encoding/decoding roundtrip if we have includes
			if includes := values.Get("include"); includes != "" {
				encoded := params.encode()
				newValues, err := url.ParseQuery(encoded)
				require.NoError(t, err)

				reDecoded := &ServerGetParams{}
				reDecoded.decode(newValues)

				assert.Equal(t, params.IncludeBMC, reDecoded.IncludeBMC)
				assert.Equal(t, params.IncludeStatus, reDecoded.IncludeStatus)
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

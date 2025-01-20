package fleetdbapi

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServerComponentQueryParamsSetQuery(t *testing.T) {
	tests := []struct {
		name string
		// input gets converted to url values before being passed into setQuery
		input   string
		current *ServerComponentGetParams
		expect  string
	}{
		{
			name:    "empty params",
			input:   "",
			current: &ServerComponentGetParams{},
			expect:  "",
		},
		{
			name:  "pagination params are included",
			input: "page=1&limit=2",
			current: &ServerComponentGetParams{
				InstalledFirmware: true,
			},
			expect: "component_include=installed_firmware&page=1&limit=2",
		},
		{
			name:  "pagination and filter params are included",
			input: "page=1&limit=2&serial__eq=123",
			current: &ServerComponentGetParams{
				Status: true,
			},
			expect: "component_include=status&serial__eq=123&page=1&limit=2",
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

func TestServerComponentGetParamsFromURLValues(t *testing.T) {
	tests := []struct {
		name           string
		queryString    string
		expectedParams *ServerComponentGetParams
	}{
		{
			name:           "empty query",
			queryString:    "",
			expectedParams: &ServerComponentGetParams{},
		},
		{
			name:        "single include",
			queryString: "component_include=capabilities",
			expectedParams: &ServerComponentGetParams{
				Capabilities: true,
			},
		},
		{
			name:        "multiple includes",
			queryString: "component_include=capabilities,installed_firmware,status",
			expectedParams: &ServerComponentGetParams{
				InstalledFirmware: true,
				Status:            true,
				Capabilities:      true,
			},
		},
		{
			name:        "multiple includes with spaces",
			queryString: "component_include=capabilities ,installed_firmware, status",
			expectedParams: &ServerComponentGetParams{
				InstalledFirmware: true,
				Status:            true,
				Capabilities:      true,
			},
		},
		{
			name:        "metadata with single namespace",
			queryString: "component_include=metadata_ns__namespace1",
			expectedParams: &ServerComponentGetParams{
				Metadata: []string{"namespace1"},
			},
		},
		{
			name:        "metadata with multiple namespaces",
			queryString: "component_include=metadata_ns__namespace1,metadata_ns__namespace2",
			expectedParams: &ServerComponentGetParams{
				Metadata: []string{"namespace1", "namespace2"},
			},
		},
		{
			name:        "everything combined",
			queryString: "component_include=capabilities,installed_firmware,status,metadata_ns__namespace1,metadata_ns__namespace2",
			expectedParams: &ServerComponentGetParams{
				InstalledFirmware: true,
				Status:            true,
				Capabilities:      true,
				Metadata:          []string{"namespace1", "namespace2"},
			},
		},
		{
			name:        "metadata with spaces",
			queryString: "component_include=metadata_ns__namespace1, metadata_ns__namespace2",
			expectedParams: &ServerComponentGetParams{
				Metadata: []string{"namespace1", "namespace2"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			values := url.Values{}
			if tt.queryString != "" {
				parsedURL, err := url.ParseQuery(tt.queryString)
				require.NoError(t, err)
				values = parsedURL
			}

			params := &ServerComponentGetParams{}
			params.fromURLValues(values)
			assert.Equal(t, tt.expectedParams, params)
		})
	}
}

func TestServerComponentGetParamsToURLValues(t *testing.T) {
	tests := []struct {
		name           string
		params         *ServerComponentGetParams
		expectedOutput string
	}{
		{
			name:           "empty params",
			params:         &ServerComponentGetParams{},
			expectedOutput: "",
		},
		{
			name: "capabilities",
			params: &ServerComponentGetParams{
				Capabilities: true,
			},
			expectedOutput: "component_include=capabilities",
		},
		{
			name: "multiple flags",
			params: &ServerComponentGetParams{
				Capabilities:      true,
				InstalledFirmware: true,
				Status:            true,
			},
			expectedOutput: "component_include=capabilities,installed_firmware,status",
		},
		{
			name: "single metadata namespace",
			params: &ServerComponentGetParams{
				Metadata: []string{"namespace1"},
			},
			expectedOutput: "component_include=metadata_ns__namespace1",
		},
		{
			name: "multiple metadata namespaces",
			params: &ServerComponentGetParams{
				Metadata: []string{"namespace1", "namespace2"},
			},
			expectedOutput: "component_include=metadata_ns__namespace1,metadata_ns__namespace2",
		},
		{
			name: "everything combined",
			params: &ServerComponentGetParams{
				Capabilities:      true,
				InstalledFirmware: true,
				Status:            true,
				Metadata:          []string{"namespace1", "namespace2"},
			},
			expectedOutput: "component_include=capabilities,installed_firmware,status,metadata_ns__namespace1,metadata_ns__namespace2",
		},
		{
			name: "metadata with empty namespaces",
			params: &ServerComponentGetParams{
				Metadata: []string{"", "namespace1", ""},
			},
			expectedOutput: "component_include=metadata_ns__namespace1",
		},
		{
			name: "only empty metadata",
			params: &ServerComponentGetParams{
				Metadata: []string{""},
			},
			expectedOutput: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.params.toURLValues()
			// If we have an expected result, verify we can decode it back
			if tt.expectedOutput != "" {
				expect, err := url.ParseQuery(tt.expectedOutput)
				assert.Nil(t, err)
				assert.Equal(t, expect, result)
			} else {
				assert.Empty(t, result)
			}
		})
	}
}

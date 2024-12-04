package fleetdbapi

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDecodeServerComponentGetParams(t *testing.T) {
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
			queryString: "include=c.capabilities",
			expectedParams: &ServerComponentGetParams{
				Capabilities: true,
			},
		},
		{
			name:        "multiple includes",
			queryString: "include=c.capabilities,c.installed_firmware,c.status",
			expectedParams: &ServerComponentGetParams{
				InstalledFirmware: true,
				Status:            true,
				Capabilities:      true,
			},
		},
		{
			name:        "multiple includes with spaces",
			queryString: "include=c.capabilities ,c.installed_firmware, c.status",
			expectedParams: &ServerComponentGetParams{
				InstalledFirmware: true,
				Status:            true,
				Capabilities:      true,
			},
		},

		{
			name:        "metadata with single namespace",
			queryString: "include=c.metadata{namespace1}",
			expectedParams: &ServerComponentGetParams{
				Metadata: []string{"namespace1"},
			},
		},
		{
			name:        "metadata with multiple namespaces",
			queryString: "include=c.metadata{namespace1,namespace2}",
			expectedParams: &ServerComponentGetParams{
				Metadata: []string{"namespace1", "namespace2"},
			},
		},
		{
			name:        "everything combined",
			queryString: "include=c.capabilities,c.installed_firmware,c.status,c.metadata{namespace1,namespace2}",
			expectedParams: &ServerComponentGetParams{
				InstalledFirmware: true,
				Status:            true,
				Capabilities:      true,
				Metadata:          []string{"namespace1", "namespace2"},
			},
		},
		{
			name:        "metadata with spaces",
			queryString: "include=c.metadata{namespace1, namespace2}",
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
			params.decode(values)
			assert.Equal(t, tt.expectedParams, params)
		})
	}
}

func TestEncodeServerComponentGetParams(t *testing.T) {
	tests := []struct {
		name           string
		params         *ServerComponentGetParams
		expectedOutput string
	}{
		{
			name:           "nil params",
			params:         nil,
			expectedOutput: "",
		},
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
			expectedOutput: "include=c.capabilities",
		},
		{
			name: "multiple flags",
			params: &ServerComponentGetParams{
				Capabilities:      true,
				InstalledFirmware: true,
				Status:            true,
			},
			expectedOutput: "include=c.capabilities,c.installed_firmware,c.status",
		},
		{
			name: "single metadata namespace",
			params: &ServerComponentGetParams{
				Metadata: []string{"namespace1"},
			},
			expectedOutput: "include=c.metadata{namespace1}",
		},
		{
			name: "multiple metadata namespaces",
			params: &ServerComponentGetParams{
				Metadata: []string{"namespace1", "namespace2"},
			},
			expectedOutput: "include=c.metadata{namespace1,namespace2}",
		},
		{
			name: "everything combined",
			params: &ServerComponentGetParams{
				Capabilities:      true,
				InstalledFirmware: true,
				Status:            true,
				Metadata:          []string{"namespace1", "namespace2"},
			},
			expectedOutput: "include=c.capabilities,c.installed_firmware,c.status,c.metadata{namespace1,namespace2}",
		},
		{
			name: "metadata with empty namespaces",
			params: &ServerComponentGetParams{
				Metadata: []string{"", "namespace1", ""},
			},
			expectedOutput: "include=c.metadata{namespace1}",
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
			result := tt.params.encode()
			assert.Equal(t, tt.expectedOutput, result)
		})
	}
}

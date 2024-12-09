package fleetdbapi

import (
	"encoding/json"
	"time"
)

// Attributes provide the ability to apply namespaced settings to an entity.
// For example servers could have attributes in the `com.equinixmetal.api` namespace
// that represents equinix metal specific attributes that are stored in the API.
// The namespace is meant to define who owns the schema and values.
type Attributes struct {
	Namespace string          `json:"namespace"`
	Data      json.RawMessage `json:"data"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

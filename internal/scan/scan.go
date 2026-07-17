package scan

import (
	"context"
	"encoding/json"
)

// Resources hold all what's exist on provider servers, volumes, DNS records etc.
// i want comment all fields they are self explonatory keep all simple.
type Resource struct {
	Provider string
	Kind     string
	ID       string
	Name     string
	Region   string
	Labels   map[string]string
	// exept this one it will store all raw API payload from provider we are going
	// store it in jsonb format in DB ( i will use postgres for it )
	Raw json.RawMessage
}

// Key return  the identity of resources acros snapshots.
func (r Resource) Key() string {
	return r.Provider + "/" + r.Kind + "/" + r.ID
}

// Scanner scans one provider for all resources visable to its credentials.
type Scanner interface {
	Name() string
	Scan(ctx context.Context) (<-chan Resource, <-chan error)
}

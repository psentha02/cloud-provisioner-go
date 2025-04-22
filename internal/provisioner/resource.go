package provisioner

import (
	"fmt"
	"time"
)

// Resource represents a cloud resource (like a VM or bucket)
type Resource struct {
	ID     string
	Name   string
	Status string
}

// NewResource creates a new cloud resource (mocked)
func NewResource(name string) *Resource {
	return &Resource{
		ID:     generateID(),
		Name:   name,
		Status: "created",
	}
}

// generateID generates a fake unique ID for the resource (mocked)
func generateID() string {
	return "resource-" + fmt.Sprintf("%d", time.Now().UnixNano())
}

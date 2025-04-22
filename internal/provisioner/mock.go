package provisioner

import (
	"sync"
)

var (
	resources = make(map[string]*Resource) // Mock database for resources
	mu        sync.Mutex                   // Mutex for concurrent safety
)

// CreateResource simulates creating a new cloud resource
func CreateResource(name string) *Resource {
	mu.Lock()
	defer mu.Unlock()
	resource := NewResource(name)
	resources[resource.ID] = resource
	return resource
}

// ListResources simulates listing all cloud resources
func ListResources() []*Resource {
	mu.Lock()
	defer mu.Unlock()
	var allResources []*Resource
	for _, r := range resources {
		allResources = append(allResources, r)
	}
	return allResources
}

// DeleteResource simulates deleting a cloud resource
func DeleteResource(id string) bool {
	mu.Lock()
	defer mu.Unlock()
	if _, exists := resources[id]; exists {
		delete(resources, id)
		return true
	}
	return false
}

package main

import (
	"fmt"
	"os"

	"github.com/psentha/cloud-provisioner-go/internal/provisioner"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cloud-provisioner-go <command>")
		fmt.Println("Available commands: create, list, delete")
		return
	}

	command := os.Args[1]

	switch command {
	case "create":
		if len(os.Args) < 3 {
			fmt.Println("Usage: cloud-provisioner-go create <resource_name>")
			return
		}
		name := os.Args[2]
		resource := provisioner.CreateResource(name)
		fmt.Printf("Created resource: %s (ID: %s, Status: %s)\n", resource.Name, resource.ID, resource.Status)

	case "list":
		resources := provisioner.ListResources()
		if len(resources) == 0 {
			fmt.Println("No resources found.")
			return
		}
		fmt.Println("Listing all resources:")
		for _, resource := range resources {
			fmt.Printf("ID: %s, Name: %s, Status: %s\n", resource.ID, resource.Name, resource.Status)
		}

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: cloud-provisioner-go delete <resource_id>")
			return
		}
		id := os.Args[2]
		success := provisioner.DeleteResource(id)
		if success {
			fmt.Printf("Resource with ID %s deleted.\n", id)
		} else {
			fmt.Printf("Resource with ID %s not found.\n", id)
		}

	default:
		fmt.Printf("Unknown command: %s\n", command)
	}
}

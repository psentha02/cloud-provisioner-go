package main

import (
	"fmt"
	"os"

	"github.com/psentha02/cloud-provisioner-go/internal/azure"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: create-rg <resourceGroupName> <location>")
		return
	}

	command := os.Args[1]

	switch command {

	case "create-rg":
		if len(os.Args) < 4 {
			fmt.Println("Usage: create-rg <resourceGroupName> <location>")
			return
		}
		name := os.Args[2]
		location := os.Args[3]
		err := azure.CreateResourceGroup(name, location)
		if err != nil {
			fmt.Printf("Error creating resource group: %v\n", err)
		}

	case "list-rg":
		err := azure.ListResourceGroups()
		if err != nil {
			fmt.Printf("Error listing resource groups: %v\n", err)
		}

	case "delete-rg":
		if len(os.Args) < 3 {
			fmt.Println("Usage: delete-rg <resourceGroupName>")
			return
		}
		name := os.Args[2]
		err := azure.DeleteResourceGroup(name)
		if err != nil {
			fmt.Printf("Error deleting resource group: %v\n", err)
		}

	default:
		fmt.Println("Unknown command.")
	}

}

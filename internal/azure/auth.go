package azure

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
)

type AzureCredentials struct {
	ClientID       string `json:"clientId"`
	ClientSecret   string `json:"clientSecret"`
	TenantID       string `json:"tenantId"`
	SubscriptionID string `json:"subscriptionId"`
}

func LoadCredentials() (*AzureCredentials, error) {
	authFile := os.Getenv("AZURE_AUTH_LOCATION")
	data, err := os.ReadFile(authFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read Azure credentials file: %v", err)
	}

	var creds AzureCredentials
	if err := json.Unmarshal(data, &creds); err != nil {
		return nil, fmt.Errorf("failed to unmarshal credentials: %v", err)
	}

	return &creds, nil
}

func CreateResourceGroup(resourceGroupName, location string) error {
	creds, err := LoadCredentials()
	if err != nil {
		return err
	}

	cred, err := azidentity.NewClientSecretCredential(
		creds.TenantID,
		creds.ClientID,
		creds.ClientSecret,
		nil,
	)
	if err != nil {
		return fmt.Errorf("authentication failed: %v", err)
	}

	client, err := armresources.NewResourceGroupsClient(creds.SubscriptionID, cred, nil)
	if err != nil {
		return fmt.Errorf("failed to create resource group client: %v", err)
	}

	ctx := context.Background()

	_, err = client.CreateOrUpdate(
		ctx,
		resourceGroupName,
		armresources.ResourceGroup{
			Location: &location,
		},
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to create resource group: %v", err)
	}

	fmt.Printf("✅ Resource Group '%s' created in '%s'\n", resourceGroupName, location)
	return nil
}

func ListResourceGroups() error {
	creds, err := LoadCredentials()
	if err != nil {
		return err
	}

	cred, err := azidentity.NewClientSecretCredential(
		creds.TenantID,
		creds.ClientID,
		creds.ClientSecret,
		nil,
	)
	if err != nil {
		return fmt.Errorf("authentication failed: %v", err)
	}

	client, err := armresources.NewResourceGroupsClient(creds.SubscriptionID, cred, nil)
	if err != nil {
		return fmt.Errorf("failed to create resource group client: %v", err)
	}

	ctx := context.Background()
	pager := client.NewListPager(nil)

	fmt.Println("Available Resource Groups:")

	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return fmt.Errorf("failed to list resource groups: %v", err)
		}
		for _, rg := range page.Value {
			fmt.Printf("- %s (%s)\n", *rg.Name, *rg.Location)
		}
	}

	return nil
}

func DeleteResourceGroup(resourceGroupName string) error {
	creds, err := LoadCredentials()
	if err != nil {
		return err
	}

	cred, err := azidentity.NewClientSecretCredential(
		creds.TenantID,
		creds.ClientID,
		creds.ClientSecret,
		nil,
	)
	if err != nil {
		return fmt.Errorf("authentication failed: %v", err)
	}

	client, err := armresources.NewResourceGroupsClient(creds.SubscriptionID, cred, nil)
	if err != nil {
		return fmt.Errorf("failed to create resource group client: %v", err)
	}

	ctx := context.Background()
	poller, err := client.BeginDelete(ctx, resourceGroupName, nil)
	if err != nil {
		return fmt.Errorf("failed to start deleting resource group: %v", err)
	}

	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to delete resource group: %v", err)
	}

	fmt.Printf("🗑️ Resource Group '%s' deleted successfully.\n", resourceGroupName)
	return nil
}

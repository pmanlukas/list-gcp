package main

import (
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"google.golang.org/api/cloudbilling/v1"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()

	// Load environment variables from .env file.
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}

	// Use the GOOGLE_APPLICATION_CREDENTIALS environment variable to authenticate with GCP.
	// Alternatively, you can specify a service account key file path using option.WithCredentialsFile().
	client, err := option.Client(ctx)
	if err != nil {
		log.Fatalf("Failed to create HTTP client: %v", err)
	}

	// Create a cloudbilling.Service client to interact with the Cloud Billing API.
	cbService, err := cloudbilling.New(client)
	if err != nil {
		log.Fatalf("Failed to create Cloud Billing client: %v", err)
	}

	// Prompt the user to enter a project name.
	fmt.Print("Enter a project name: ")
	var projectName string
	fmt.Scanln(&projectName)

	// Call the cloudbilling.Projects.GetBillingInfo method to retrieve the billing information for the project.
	billingInfo, err := cbService.Projects.GetBillingInfo(fmt.Sprintf("projects/%s", projectName)).Context(ctx).Do()
	if err != nil {
		log.Fatalf("Failed to retrieve billing info: %v", err)
	}

	// Print the billing information for the project.
	fmt.Printf("Project ID: %s\nBilling Account Name: %s\nBilling Enabled: %t\n",
		billingInfo.ProjectId, billingInfo.BillingAccountName, billingInfo.BillingEnabled)
}

package main

import (
	"context"
	"fmt"
	"log"


	"google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/option"
)

func main() {
	var path string
	fmt.Print("Enter the path to your GCP credentials file: ")
	fmt.Scan(&path)

	// Use the provided credentials file to create a new client
	

	// Create a new Cloud Resource Manager client using the provided credentials
	crmClient, err := cloudresourcemanager.NewService(context.Background(), option.WithCredentialsFile(path))
	if err != nil {
		log.Fatal(err)
	}

	// List the projects associated with the provided credentials
	projects, err := crmClient.Projects.List().Do()
	if err != nil {
		log.Fatal(err)
	}

	// Print the names of the projects
	fmt.Println("Projects:")
	for _, project := range projects.Projects {
		fmt.Printf("- %s\n", project.Name)
	}
}

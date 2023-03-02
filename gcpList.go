package main

import (
    "context"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"

    "cloud.google.com/go/billing/v1"
    "cloud.google.com/go/storage"
)

func main() {
    // Create a new billing client.
    ctx := context.Background()
    client, err := billing.NewClient(ctx, "https://cloud.google.com/billing/v1", os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
    if err != nil {
        log.Fatal(err)
    }

    // Prompt the user for the project name.
    projectName := promptUserForProjectName()

    // Create a new billing report.
    report, err := client.CreateReport(ctx, "projects/"+projectName, billing.Month("2023-01"))
    if err != nil {
        log.Fatal(err)
    }

    // Export the billing report.
    data, err := report.ExportToCsv()
    if err != nil {
        log.Fatal(err)
    }

    // Save the billing report locally.
    err = ioutil.WriteFile("report.csv", data, 0644)
    if err != nil {
        log.Fatal(err)
    }

    // Print a message.
    fmt.Println("Report saved to report.csv.")
}

// promptUserForProjectName prompts the user for the project name.
func promptUserForProjectName() string {
    fmt.Println("Enter the name of the project: ")
    projectName, err := os.ReadString(os.Stdin)
    if err != nil {
        log.Fatal(err)
    }
    return projectName
}
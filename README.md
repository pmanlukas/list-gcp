# list-gcp

##GCP Resource Viewer
This is a simple Go script that authenticates with Google Cloud Platform (GCP) and retrieves the list of projects, folders, and org policy constraints for the selected project.

##Prerequisites
- Go 1.16 or higher
- A GCP project with the appropriate permissions
- A service account key file with appropriate permissions
- (Optional) The github.com/joho/godotenv package if you want to load environment variables from an .env file

## Getting started
- Clone this repository to your local machine.
- Set the GOOGLE_APPLICATION_CREDENTIALS environment variable to the path of your service account key file. Alternatively, you can create an .env file in the same directory as the Go script and add the following line to it: GOOGLE_APPLICATION_CREDENTIALS=/path/to/your/private-key-file.json.
- (Optional) If you're using an .env file, make sure to install the github.com/joho/godotenv package by running the command go get github.com/joho/godotenv.
- Build and run the script by running the command go run main.go.

## Usage
When you run the script, it will retrieve the list of projects for your GCP account and ask you to select a project by entering the corresponding number. Then, it will retrieve the list of folders and org policy constraints for the selected project and print them to the console.

## Contributing
If you find any issues or have suggestions for improvements, please feel free to open an issue or a pull request.

### License
This project is licensed under the MIT License - see the LICENSE file for details.

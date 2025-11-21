package main

import (
	"log"
	authGoogle "oreonproject/basalt/cmd/auth"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
)

func main() {
	clientSecret, err := os.ReadFile("credentials.json") // Reads the Credentials File
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// Gets the Client Secret from the JSON file and makes a Read Request on the user's Drive
	config, err := google.ConfigFromJSON(clientSecret, drive.DriveReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	client := authGoogle.GetClient(config)
	// This uses a Deprecated Method which will be replaced in the future
	drvService, err := drive.New(client) // Initialises a New Drive Service
	if err != nil {
		log.Fatalf("Unable to retreive Drive client: %v", err)
	}

	FileList, err := drvService.Files.List().PageSize(10).Fields("files(id, name)").Do() // From the service it requests 10 files by ID and name fields
	if err != nil {
		log.Fatalf("Unable to retreive files: %v", err)
	}

	if len(FileList.Files) == 0 {
		log.Println("No lines were found.")
	} else {
		log.Println("Files:")
		for _, i := range FileList.Files {
			log.Printf("%s (%s)\n", i.Name, i.Id)
		}
	}

}

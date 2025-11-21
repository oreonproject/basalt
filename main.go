package main

import (
	"log"
	authGoogle "oreonproject/basalt/cmd/auth"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
)

func main() {
	clientSecret, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(clientSecret, drive.DriveReadonlyScope)

	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	client := authGoogle.GetClient(config)
	drvService, err := drive.New(client)
	if err != nil {
		log.Fatalf("Unable to retreive Drive client: %v", err)
	}

	r, err := drvService.Files.List().PageSize(10).Fields("files(id, name)").Do()
	if err != nil {
		log.Fatalf("Unable to retreive files: %v", err)
	}

	if len(r.Files) == 0 {
		log.Println("No lines were found.")
	} else {
		log.Println("Files:")
		for _, i := range r.Files {
			log.Printf("%s (%s)\n", i.Name, i.Id)
		}
	}

}

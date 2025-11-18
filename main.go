package main

import (
	"log"
	"os"
	"google.golang.org/api/drive/v3"
	"golang.org/x/oauth2/google"
        authGoogle "oreonproject/basalt/auth/google"
)

func main() {
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}
	config, err := google.ConfigFromJSON(b, drive.DriveReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := authGoogle.GetClient(config)
	srv, err := drive.New(client)
	if err != nil {
		log.Fatalf("Unable to retreive Drive client: %v", err)
	}
	r, err := srv.Files.List().PageSize(10).Fields("files(id, name)").Do()
	if err != nil {
		log.Fatalf("Unable to retreive files: %v", err)
	}
	if len(r.Files) == 0 {
		log.Println("No lines were found.")
	} else {
		log.Println("Files:")
		for _, i := range r.Files  {
			log.Printf("%s (%s)\n", i.Name, i.Id)
		}
	}
	
}

package google


import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"golang.org/x/oauth2"
)

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err 
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err 
}

func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to %s\n", path)
	f, err := os.Create(path)
	if err != nil {
		log.Fatalf("Unable to cache OAuth token: %v", err)
		
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline, oauth2.ApprovalForce)
	fmt.Printf("Go to the following link in your browser then type the code: \n%v\n", authURL)
	    var code string 
	    if _, err := fmt.Scan(&code); err != nil {
		    log.Fatalf("Unable to read code: %v", err)
	    }
	    tok, err := config.Exchange(context.Background(), code)
	    if err != nil {
		    log.Fatalf("Unable to retreive token from web: %v", err)
	    }
	    return tok
}


func GetClient(config *oauth2.Config) *http.Client {
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	} else if !tok.Valid() {
		ts := config.TokenSource(context.Background(), tok)
		tok, err := ts.Token()
		if err != nil {
			tok = getTokenFromWeb(config)
		}
		saveToken(tokFile, tok)
	}


	return config.Client(context.Background(), tok)
}

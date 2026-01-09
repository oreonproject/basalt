package authGoogle

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/zalando/go-keyring"
)


func oauthHandler(w http.ResponseWriter, r *http.Request) {
 	code := r.URL.Query().Get("code")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	r_URI := "http://localhost:8080/"
	code_verifier, err := keyring.Get("basalt", "code_verifier")
	if err != nil {
		log.Fatalf("Failed to read code verifier from keyring: %v", err)
	}
	form := url.Values{}
	form.Set("code", code)
	form.Set("client_id", clientID)
	form.Set("client_secret", clientSecret)
	form.Set("redirect_uri", r_URI)
	form.Set("grant_type", "authorization_code")
	form.Set("code_verifier", code_verifier)
	resp, err := http.PostForm("https://oauth2.googleapis.com/token", form)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)

	var tokenData map[string]any
	if err := json.Unmarshal(respBody, &tokenData); err != nil {
		log.Fatalf("Failed to parse response: %v", err)
	}
	refresh_token, ok := tokenData["refresh_token"].(string)
	if !ok || refresh_token == "" {
		log.Println("No refresh token was sent.")
	} else {
		err := keyring.Set("basalt", "refresh_token", refresh_token)
		if err != nil {
			log.Printf("Failed to write refresh token to keyring: %v", err)
		}
		fmt.Println("Refresh Token:", refresh_token)
	}
	fmt.Fprintf(w, "Authentication success,  you may close this window.")
}

func RunServer() {
	http.HandleFunc("/", oauthHandler)
	http.ListenAndServe(":8080", nil)
}

package authGoogle

import (
	"fmt"
	"net/url"
	"oreonproject/basalt/oauth"
	"oreonproject/basalt/utils"
	"strings"
	"os"
	"github.com/zalando/go-keyring"
	"github.com/joho/godotenv"
)

func CraftAuthURI() string {
	log := utils.LogInit("authReq.log")
	err := godotenv.Load()
	if err != nil {
		log.Printf("Failed to load envrionment file: %v", err)
	}
	clientID := os.Getenv("CLIENT_ID")
	if clientID == ""{
		log.Fatal("CLIENT_ID is not set.")
	}

	authServer := "https://accounts.google.com/o/oauth2/v2/auth" // Defines the base Auth server to send the request to
	// Creates a Values struct to hold the Key and Value pairs for OAuth2
	params := url.Values{}
	// Creates a URL struct for the Redirect URL
	redirect := &url.URL{
		Scheme: "http",
		Host:   "localhost:8080",
		Path:   "/",
	}
	verifier := oauth.CodeVerifierKeyGen()
	if err := keyring.Set("basalt", "code_verifier", verifier); err != nil {
		log.Fatalf("Failed to save code verifier to keyring: %v", err)
	}

	params.Add("client_id", clientID)
	params.Add("access_type", "offline") // Allows us to generate refresh tokens without the client going into their browsers
	params.Add("redirect_uri", redirect.String())
	params.Add("include_granted_scopes", "true")
	params.Add("response_type", "code")                                                                                                             // Response type is an AuthCode
	params.Add("scope", strings.Join([]string{"openid", "https://www.googleapis.com/auth/calendar", "https://www.googleapis.com/auth/drive"}, " ")) // our scopes
	params.Add("code_challenge", oauth.CodeChallengeGen(verifier))                                                                                          // Generates a Code Challenge
	params.Add("code_challenge_method", "S256")                                                                                                     // Code Challenge method is a base64 RAWURLENCODED SHA256 hash of the CodeVerifier
	params.Add("state", oauth.StateTokGen())                                                                                                        // Generates a State Token to Prevent Cross Site Request Forgery

	authURL := fmt.Sprintf("%s?%s", authServer, params.Encode()) // Formats the Parameters with the AuthURL
	authURL = strings.ReplaceAll(authURL, "+", "%20")            // Replaces all +'s with %20 as google OAuth2 servers crash out without them???
	log.Print("Published Auth URL")
	return authURL

}

package authGoogle

import (
	"fmt"
	"net/url"
	"strings"
)

func CraftAuthURI() string {
	authServer := "https://accounts.google.com/o/oauth2/v2/auth"
	params := url.Values{}
	redirect := &url.URL{
		Scheme: "http",
		Host:   "localhost:8080",
		Path:   "/",
	}

	params.Add("client_id", "844897693697-23j2de25lbf5fdmsh5lfs3hn0fr9kkh3.apps.googleusercontent.com")
	params.Add("access_type", "offline")
	params.Add("redirect_uri", redirect.String())
	params.Add("include_granted_scopes", "true")
	params.Add("response_type", "code")
	params.Add("scope", strings.Join([]string{"openid", "https://www.googleapis.com/auth/calendar", "https://www.googleapis.com/auth/drive"}, " "))
	params.Add("code_challenge", CodeChallengeGen())
	params.Add("code_challenge_method", "S256")
	params.Add("state", StateTokGen())

	authURL := fmt.Sprintf("%s?%s", authServer, params.Encode())
	authURL = strings.ReplaceAll(authURL, "+", "%20")
	fmt.Println(authURL)
	return authURL

}

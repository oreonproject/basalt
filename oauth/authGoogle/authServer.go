package authGoogle

import (
	"fmt"
	"log"
	"net/http"
	"github.com/zalando/go-keyring"
)


func oauthHandler(w http.ResponseWriter, r *http.Request) {
 	code := r.URL.Query().Get("code")
	err := keyring.Set("basalt", "user", code)
	if err != nil {
		log.Printf("Failed to save token to keyring: %v", err)
	}
	fmt.Fprintf(w, "Authentication success,  you may close this window.")
}

func RunServer() {
	http.HandleFunc("/", oauthHandler)
	http.ListenAndServe(":8080", nil)
}

package authGoogle

import (
	"fmt"
	"log"
	"net/http"
	"os"
)


func oauthHandler(w http.ResponseWriter, r *http.Request) {
 	code := r.URL.Query().Get("code")
	sec_file, err := os.Create(".secret")
	defer sec_file.Close()
	if err != nil {
		log.Printf("Error opening secret file: %v", err)

		return
	}
	_, err = sec_file.WriteString(code)
	if err != nil {
		log.Printf("Writing to secret file failed: %v", err)
	}
	fmt.Fprintf(w, "Authentication success,  you may close this window.")
}

func RunServer() {
	http.HandleFunc("/", oauthHandler)
	http.ListenAndServe(":8080", nil)
}

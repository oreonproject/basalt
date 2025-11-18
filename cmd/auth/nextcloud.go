package auth

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Credentials struct {
	URL      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var NextcloudCmd = &cobra.Command{
	Use:   "nextcloud",
	Short: "Login with Nextcloud",
	Run: func(cmd *cobra.Command, args []string) {
		var url string
		var username string
		var password string
		cmd.Println("Put in your nextcloud url: ")
		fmt.Scanln(&url)
		cmd.Println("Put in your nextcloud username: ")
		fmt.Scanln(&username)
		cmd.Println("Put in your nextcloud password: ")
		fmt.Scanln(&password)
		cmd.Println("Saving credentials...")

		homeDir, err := os.UserHomeDir()
		if err != nil {
			cmd.Println("Error getting home directory:", err)
			return
		}

		if err := os.MkdirAll(homeDir+"/.basalt", 0700); err != nil {
			cmd.Println("Error creating .basalt directory:", err)
			return
		}

		credsFilePath := homeDir + "/.basalt/credentials.json"
		credsFile, err := os.Create(credsFilePath)
		if err != nil {
			cmd.Println("Error creating credentials file:", err)
			return
		}
		defer credsFile.Close()

		credentials := Credentials{
			URL:      url,
			Username: username,
			Password: password,
		}

		if err := json.NewEncoder(credsFile).Encode(credentials); err != nil {
			cmd.Println("Error writing credentials to file:", err)
			return
		}

		cmd.Println("Credentials saved successfully!")
	},
}

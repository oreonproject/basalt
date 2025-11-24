package nextcloud

import (
	"encoding/json"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/zalando/go-keyring"
)

type Credentials struct {
	URL      string `json:"url"`
	Username string `json:"username"`
}

var NextcloudCmd = &cobra.Command{
	Use:   "nextcloud",
	Short: "Login with Nextcloud",
	Run: func(cmd *cobra.Command, args []string) {
		username, err := cmd.Flags().GetString("username")
		if err != nil {
			panic(err)
		}
		password, err := cmd.Flags().GetString("password")
		if err != nil {
			panic(err)
		}
		url, err := cmd.Flags().GetString("url")
		if err != nil {
			panic(err)
		}

		err = keyring.Set("basalt", username, password)
		if err != nil {
			log.Fatal(err)
		}

		// Store username and url in ~/basalt/credentials.json
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
		}

		if err := json.NewEncoder(credsFile).Encode(credentials); err != nil {
			cmd.Println("Error writing credentials to file:", err)
			return
		}

	},
}

func init() {
	NextcloudCmd.Flags().StringP("username", "u", "", "Nextcloud Username")
	NextcloudCmd.Flags().StringP("password", "p", "", "Nextcloud Password")
	NextcloudCmd.Flags().StringP("url", "l", "", "Nextcloud URL")

	if err := NextcloudCmd.MarkFlagRequired("username"); err != nil {
		panic(err)
	}
	if err := NextcloudCmd.MarkFlagRequired("password"); err != nil {
		panic(err)
	}
	if err := NextcloudCmd.MarkFlagRequired("url"); err != nil {
		panic(err)
	}
}

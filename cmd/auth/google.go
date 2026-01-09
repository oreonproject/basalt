package auth

import (
	"fmt"
	"log"
	"oreonproject/basalt/oauth/authGoogle"
	"os/exec"

	"github.com/spf13/cobra"
)

var GoogleCmd = &cobra.Command{
	Use:     "google",
	Short:   "Login with google",
	Long:    "Description:\nOpens the Authorisation URL in your default browser,\nThere you can grant the application permissions to access your resources on google's cloud services\nBasalt asks for modification access to your Google Drive and Google Calendar to provide a seamless integration with other features",
	Aliases: []string{"g", "ggl"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Authenticate in your browser.")
		url := authGoogle.CraftAuthURI()
		err := exec.Command("xdg-open", url).Start()
		if err != nil {
			cmd.Println("Error opening browser:", err)
			log.Print("Error opening browser:", err)
		}
		authGoogle.RunServer()
		
	},
}

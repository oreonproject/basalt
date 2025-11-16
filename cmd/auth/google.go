package auth

import (
	"github.com/spf13/cobra"
)

var GoogleCmd = &cobra.Command{
	Use:   "google",
	Short: "Login with google",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Logging in with google...")
	},
}

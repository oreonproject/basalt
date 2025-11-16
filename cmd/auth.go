package cmd

import (
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Login with cloud providers",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Auth command executed")
	},
}

package cmd

import (
	"log"
	"oreonproject/basalt/cmd/auth"
	"oreonproject/basalt/utils"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "basalt",
	Short: "Root command",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Hello world!")
		log.Print("Root Command Executed")
	},
}

func Execute() {
	utils.LogInit("logs/root.log")
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	utils.LogInit("logs/root.log")
	rootCmd.AddCommand(auth.AuthCmd)

	auth.AuthCmd.AddCommand(auth.GoogleCmd)
	auth.AuthCmd.AddCommand(auth.NextcloudCmd)
}

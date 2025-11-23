package main

import (
	"oreonproject/basalt/cmd"
	"oreonproject/basalt/oauth/authGoogle"
	"oreonproject/basalt/utils"
)

func main() {
	log := utils.LogInit("main.log")
	log.Print("Logger was Setup")
	cmd.Execute()
	log.Print("Executed root command")
	authGoogle.CraftAuthURI()
}

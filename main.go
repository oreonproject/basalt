package main

import (
	"log"
	"oreonproject/basalt/cmd"
	"oreonproject/basalt/utils"
)

func main() {
	utils.LogInit("logs/main")
	log.Print("Logger was Setup")
	cmd.Execute()
	log.Print("Executed root command")
}

/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"
	"oreonproject/basalt/cmd"
	"oreonproject/basalt/utils"
)

func main() {
	cmd.Execute()
	utils.LogSetup("logs/main.log")
	log.Print("cute test :3")
}

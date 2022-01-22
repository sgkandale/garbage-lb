package main

import (
	"fmt"
	"log"

	"garbagelb/serverLoad"
	"garbagelb/ui"
)

func main() {
	log.Println("Starting GarbageLB...")

	load, _ := serverLoad.GetServerLoad()
	fmt.Println("Server Load :", load)

	ui.ServeWebUI()
}

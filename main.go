package main

import (
	"fmt"
	"log"

	"garbagelb/adminServer"
	"garbagelb/serverLoad"
)

func main() {
	log.Println("Starting GarbageLB...")

	load, _ := serverLoad.GetServerLoad()
	fmt.Println("Server Load :", load)

	adminServer.ServeWebUI()
}

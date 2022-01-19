package main

import (
	"log"
	"simplelb/ui"
)

func main() {
	log.Println("Starting SimpleLB...")

	ui.ServeWebUI()
}

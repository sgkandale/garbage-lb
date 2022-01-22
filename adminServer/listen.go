package adminServer

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"garbagelb/config"

	rice "github.com/GeertJohan/go.rice"
)

func (webServer *AdminServer) Listen(wg *sync.WaitGroup) {

	appBox, err := rice.FindBox("../ui_src/build")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", serveAppHandler(appBox))
	http.HandleFunc("/serverLoad", GetServerLoad)

	log.Println(
		fmt.Sprintf(
			"Admin UI Server starting at port %s...",
			config.Config.Admin.Port,
		),
	)

	go func() {
		defer wg.Done()
		log.Fatal(webServer.ListenAndServe())
	}()
}

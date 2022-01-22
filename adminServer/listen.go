package adminServer

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"garbagelb/config"

	rice "github.com/GeertJohan/go.rice"
)

func (webServer *AdminServer) Listen(wg *sync.WaitGroup, listener *config.Listener) {

	webServer.Addr = fmt.Sprintf(":%d", listener.Port)

	appBox, err := rice.FindBox("../ui_src/build")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", serveAppHandler(appBox))
	http.HandleFunc("/serverLoad", GetServerLoad)

	log.Println(
		fmt.Sprintf(
			"Admin UI Server starting at port %d...",
			listener.Port,
		),
	)

	go func() {
		defer wg.Done()
		log.Fatal(webServer.ListenAndServe())
	}()
}

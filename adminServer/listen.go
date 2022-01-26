package adminServer

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"garbagelb/config"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
)

func (webServer *AdminServer) Listen(wg *sync.WaitGroup, listener *config.Listener) {

	appBox, err := rice.FindBox("../ui_src/build")
	if err != nil {
		log.Fatal(err)
	}

	muxRouter := mux.NewRouter()

	muxRouter.HandleFunc("/serverLoad", GetServerLoad).Methods(http.MethodGet, http.MethodOptions)
	muxRouter.HandleFunc("/cluster", GetClusters).Methods(http.MethodGet, http.MethodOptions)
	muxRouter.HandleFunc("/listener", GetListeners).Methods(http.MethodGet, http.MethodOptions)

	// define frontend routes at last to avoid 404
	muxRouter.HandleFunc("/", serveAppHandler(appBox))
	muxRouter.Handle("/static/{dir}/{file}", http.FileServer(appBox.HTTPBox()))
	muxRouter.Handle("/{file}", http.FileServer(appBox.HTTPBox()))

	webServer.Addr = fmt.Sprintf("localhost:%d", listener.Port)
	webServer.Handler = muxRouter

	log.Printf(
		"Admin UI Server starting at port %d...\n",
		listener.Port,
	)

	go func() {
		defer wg.Done()
		log.Fatal(webServer.ListenAndServe())
	}()
}

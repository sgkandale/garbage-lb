package ui

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"simplelb/config"

	rice "github.com/GeertJohan/go.rice"
)

func ServeWebUI() {

	appBox, err := rice.FindBox("../ui_src/build")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", serveAppHandler(appBox))

	log.Println(
		fmt.Sprintf(
			"UI Server starting at port %s...",
			config.Config.Admin.Port,
		),
	)

	log.Fatal(http.ListenAndServe(":"+config.Config.Admin.Port, nil))
}

func serveAppHandler(app *rice.Box) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		indexFile, err := app.Open("index.html")
		if err != nil {
			http.Error(
				w,
				"Internal Server Error",
				http.StatusInternalServerError,
			)
			return
		}

		http.ServeContent(w, r, "index.html", time.Time{}, indexFile)
	}
}

package adminServer

import (
	"net/http"
	"time"

	rice "github.com/GeertJohan/go.rice"
)

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

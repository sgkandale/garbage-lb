package http

import (
	"net/http"
)

func rejectionHandler(w *http.ResponseWriter, r *http.Request) {
	(*w).WriteHeader(http.StatusServiceUnavailable)
	(*w).Write([]byte("Service Unavailable"))
}

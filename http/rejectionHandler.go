package http

import (
	goHttp "net/http"
)

func rejectionHandler(w *goHttp.ResponseWriter, r *goHttp.Request) {
	(*w).WriteHeader(goHttp.StatusServiceUnavailable)
	(*w).Write([]byte("Service Unavailable"))
}

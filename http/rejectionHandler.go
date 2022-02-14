package http

import (
	goHttp "net/http"
)

func rejectUnavailable(w *goHttp.ResponseWriter, r *goHttp.Request) {
	(*w).WriteHeader(goHttp.StatusServiceUnavailable)
	(*w).Write([]byte("Service Unavailable"))
}

func rejectPayloadTooLarge(w *goHttp.ResponseWriter, r *goHttp.Request) {
	(*w).WriteHeader(goHttp.StatusRequestEntityTooLarge)
	(*w).Write([]byte("Request Entity Too Large"))
}

func rejectTooManyRequests(w *goHttp.ResponseWriter, r *goHttp.Request) {
	(*w).WriteHeader(goHttp.StatusTooManyRequests)
	(*w).Write([]byte("Too Many Requests"))
}

func rejectNotAcceptable(w *goHttp.ResponseWriter, r *goHttp.Request) {
	(*w).WriteHeader(goHttp.StatusNotAcceptable)
	(*w).Write([]byte("Not Acceptable"))
}

func rejectInternalProxyError(w *goHttp.ResponseWriter, r *goHttp.Request) {
	(*w).WriteHeader(goHttp.StatusInternalServerError)
	(*w).Write([]byte("Internal Proxy Error"))
}

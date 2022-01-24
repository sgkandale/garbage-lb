package http

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"garbagelb/config"
)

func forwardRequest(w *http.ResponseWriter, r *http.Request, endpoint *config.Endpoint) {

	reverseProxy := httputil.NewSingleHostReverseProxy(
		&url.URL{
			Host:   fmt.Sprintf("%s:%d", endpoint.Address, endpoint.Port),
			Scheme: "http",
		},
	)
	reverseProxy.ServeHTTP(*w, r)
}

package http

import (
	"fmt"
	goHttp "net/http"
	"net/http/httputil"
	"net/url"

	"garbagelb/config"
)

func forwardRequest(w *goHttp.ResponseWriter, r *goHttp.Request, endpoint *config.Endpoint) {

	// increment active connections counter
	endpoint.IncrementActiveConnectionsCounter()

	// forward request
	reverseProxy := httputil.NewSingleHostReverseProxy(
		&url.URL{
			Host:   fmt.Sprintf("%s:%d", endpoint.Address, endpoint.Port),
			Scheme: "http",
		},
	)
	reverseProxy.ServeHTTP(*w, r)

	// decrement active connections counter
	endpoint.DecrementActiveConnectionsCounter()
}

package http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	goHttp "net/http"

	"garbagelb/config"
)

func forwardRequest(w *goHttp.ResponseWriter, r *goHttp.Request, endpoint *config.Endpoint) {

	// increment active connections counter
	endpoint.IncrementActiveConnectionsCounter()

	// decrement active connections counter on exit
	defer endpoint.DecrementActiveConnectionsCounter()

	// construct url from endpoint
	endpointURL := fmt.Sprintf(
		"%s://%s:%d%s",
		endpoint.Protocol,
		endpoint.Address,
		endpoint.Port,
		r.URL.Path,
	)

	// copy body for new request
	toForwardBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		goHttp.Error(*w, err.Error(), goHttp.StatusInternalServerError)
		return
	}

	// request to forward
	endpointRequest, err := goHttp.NewRequest(
		r.Method,
		endpointURL,
		bytes.NewReader(toForwardBody),
	)
	if err != nil {
		goHttp.Error(*w, err.Error(), goHttp.StatusBadGateway)
		return
	}

	// copy headers from original request
	for header, values := range r.Header {
		for _, value := range values {
			endpointRequest.Header.Add(header, value)
		}
	}

	endpointRequest.URL.RawQuery = r.URL.Query().Encode()

	httpClient := goHttp.Client{}

	responseFromEndpoint, err := httpClient.Do(endpointRequest)
	if err != nil {
		goHttp.Error(*w, err.Error(), goHttp.StatusBadGateway)
		return
	}

	// copy headers from response
	for key, value := range responseFromEndpoint.Header {
		for _, eachValue := range value {
			(*w).Header().Add(key, eachValue)
		}
	}

	// copy body to http writer
	receivedBody, err := ioutil.ReadAll(responseFromEndpoint.Body)
	if err != nil {
		goHttp.Error(*w, err.Error(), goHttp.StatusInternalServerError)
		return
	}

	(*w).Write(receivedBody)

	responseFromEndpoint.Body.Close()

}

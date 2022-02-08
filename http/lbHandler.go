package http

import (
	"log"
	goHttp "net/http"
	"strings"
)

func (server *HTTPServer) LBHandler(w goHttp.ResponseWriter, r *goHttp.Request) {

	// get max connections count
	maxAllowedConnections := server.Listener.MaxConnections
	if maxAllowedConnections > 0 {
		// compare active connections count
		if server.Listener.ActiveConnections >= maxAllowedConnections {
			rejectionHandler(&w, r)
			return
		} else {
			// increment active connections count
			server.Listener.IncrementActiveConnections()
		}

		// decrement active connections on exit
		defer func() {
			server.Listener.DecrementActiveConnections()
		}()
	}

	if server.Listener.Filter != nil {
	rulesIterator:
		// iterate over the rules
		for _, eachRule := range server.Listener.Filter.Rules {
			// if the rule is enabled
			if eachRule.Enabled {
				// switch over the rule type
				switch eachRule.Type {
				case "path":
					// get provided path
					requiredPath := eachRule.Value
					// get the request path
					requestPath := r.URL.Path
					// check the incoming path length
					if len(requestPath) >= len(requiredPath) {
						// compare incoming path with provided path
						if requestPath[0:len(requiredPath)] == requiredPath {
							// if rule action is reject, return
							if eachRule.Action == "reject" {
								rejectionHandler(&w, r)
								return
							}
							// if rule action is forward
							if eachRule.Action == "forward" {
								// if match, then forward the request to target cluster
								if eachRule.TargetCluster != nil {
									clusterHadler(&w, r, eachRule.TargetCluster)
									return
								} else {
									rejectionHandler(&w, r)
									return
								}
							}
						}
					}
					// continue to next rule
					continue rulesIterator
				case "header":
					// get provided header name
					requiredHeaderName := eachRule.Key
					// get provided header value
					requiredHeaderValue := eachRule.Value
					// check the incoming header
					incomingHeaderValue := r.Header.Get(requiredHeaderName)
					// if header value matches required value
					if incomingHeaderValue == requiredHeaderValue {
						// if rule action is reject, return
						if eachRule.Action == "reject" {
							rejectionHandler(&w, r)
							return
						}
						// if rule action is forward
						if eachRule.Action == "forward" {
							// if match, then forward the request to target cluster
							if eachRule.TargetCluster != nil {
								clusterHadler(&w, r, eachRule.TargetCluster)
								return
							} else {
								rejectionHandler(&w, r)
								return
							}
						}
					}
					// continue to next rule
					continue rulesIterator
				case "cookie":
					// get provided cookie name
					requiredCookieName := eachRule.Key
					// get provided cookie value
					requiredCookieValue := eachRule.Value
					// check the incoming cookie
					incomingCookie, err := r.Cookie(requiredCookieName)
					if err != nil {
						if err == goHttp.ErrNoCookie {
							// if no cookie, then continue to next rule
							continue rulesIterator
						} else {
							log.Printf(
								"error reading cookie {%s} for listener {%s} : %s",
								requiredCookieName,
								server.Listener.Name,
								err.Error(),
							)
							rejectionHandler(&w, r)
							return
						}
					}
					// if cookie value matches required value
					if incomingCookie.Value == requiredCookieValue {
						// if rule action is reject, return
						if eachRule.Action == "reject" {
							rejectionHandler(&w, r)
							return
						}
						// if rule action is forward
						if eachRule.Action == "forward" {
							// if match, then forward the request to target cluster
							if eachRule.TargetCluster != nil {
								clusterHadler(&w, r, eachRule.TargetCluster)
								return
							} else {
								rejectionHandler(&w, r)
								return
							}
						}
					}
					// continue to next rule
					continue rulesIterator
				case "source_ip":
					// get provided source ip
					requiredSourceIP := eachRule.Value
					// get incoming source ip
					incomingSourceIP := r.RemoteAddr
					// split incoming source ip at :
					splittedSourceIP := strings.Split(incomingSourceIP, ":")
					if len(splittedSourceIP) > 1 {
						// get the ip
						incomingSourceIP = splittedSourceIP[0]
					}
					// check the incoming source ip
					if incomingSourceIP == requiredSourceIP {
						// if rule action is reject, return
						if eachRule.Action == "reject" {
							rejectionHandler(&w, r)
							return
						}
						// if rule action is forward
						if eachRule.Action == "forward" {
							// if match, then forward the request to target cluster
							if eachRule.TargetCluster != nil {
								clusterHadler(&w, r, eachRule.TargetCluster)
								return
							} else {
								rejectionHandler(&w, r)
								return
							}
						}
					}
					// continue to next rule
					continue rulesIterator
				case "source_port":
					// get provided source port
					requiredSourcePort := eachRule.Value
					// get incoming source ip
					incomingSourceIP := r.RemoteAddr
					incomingSourcePort := "0"
					// split incoming source ip at :
					splittedSourcePort := strings.Split(incomingSourceIP, ":")
					if len(splittedSourcePort) > 1 {
						// get the port
						incomingSourcePort = splittedSourcePort[1]
					}
					// check the incoming source port
					if incomingSourcePort == requiredSourcePort {
						// if rule action is reject, return
						if eachRule.Action == "reject" {
							rejectionHandler(&w, r)
							return
						}
						// if rule action is forward
						if eachRule.Action == "forward" {
							// if match, then forward the request to target cluster
							if eachRule.TargetCluster != nil {
								clusterHadler(&w, r, eachRule.TargetCluster)
								return
							} else {
								rejectionHandler(&w, r)
								return
							}
						}
					}
					// continue to next rule
					continue rulesIterator
				case "referrer", "referer":
					// get provided referrer
					requiredReferrer := eachRule.Value
					// get incoming referrer
					incomingReferrer := r.Referer()
					// check the incoming referrer
					if incomingReferrer == requiredReferrer {
						// if rule action is reject, return
						if eachRule.Action == "reject" {
							rejectionHandler(&w, r)
							return
						}
						// if rule action is forward
						if eachRule.Action == "forward" {
							// if match, then forward the request to target cluster
							if eachRule.TargetCluster != nil {
								clusterHadler(&w, r, eachRule.TargetCluster)
								return
							} else {
								rejectionHandler(&w, r)
								return
							}
						}
					}
					// continue to next rule
					continue rulesIterator
				case "method":
					// get provided method
					requiredMethod := eachRule.Value
					// get incoming method
					incomingMethod := r.Method
					// check the incoming method
					if strings.EqualFold(incomingMethod, requiredMethod) {
						// if rule action is reject, return
						if eachRule.Action == "reject" {
							rejectionHandler(&w, r)
							return
						}
						// if rule action is forward
						if eachRule.Action == "forward" {
							// if match, then forward the request to target cluster
							if eachRule.TargetCluster != nil {
								clusterHadler(&w, r, eachRule.TargetCluster)
								return
							} else {
								rejectionHandler(&w, r)
								return
							}
						}
					}
					// continue to next rule
					continue rulesIterator
				case "host":
					// get provided host
					requiredHost := eachRule.Value
					// get incoming host
					incomingHost := r.Host
					// check the incoming host
					if strings.EqualFold(incomingHost, requiredHost) {
						// if rule action is reject, return
						if eachRule.Action == "reject" {
							rejectionHandler(&w, r)
							return
						}
						// if rule action is forward
						if eachRule.Action == "forward" {
							// if match, then forward the request to target cluster
							if eachRule.TargetCluster != nil {
								clusterHadler(&w, r, eachRule.TargetCluster)
								return
							} else {
								rejectionHandler(&w, r)
								return
							}
						}
					}
					// continue to next rule
					continue rulesIterator
				default:
					rejectionHandler(&w, r)
					return
				}
			}
			// continue to next rule
			continue rulesIterator
		}
	}

	rejectionHandler(&w, r)
}

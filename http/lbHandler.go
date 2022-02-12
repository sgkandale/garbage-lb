package http

import (
	"log"
	goHttp "net/http"
	"strings"

	"garbagelb/internal/defaults"
)

func (server *HTTPServer) LBHandler(w goHttp.ResponseWriter, r *goHttp.Request) {

	// get max connections count
	// compare active connections count
	maxAllowedConnections := server.Listener.MaxConnections
	if maxAllowedConnections > 0 && server.Listener.ActiveConnections >= maxAllowedConnections {
		log.Println("max connections reached")
		rejectionHandler(&w, r)
		return
	}

	// increment active connections count
	server.Listener.IncrementActiveConnections()

	// decrement active connections on exit
	server.Listener.DecrementActiveConnections()

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
					// compare incoming path length
					if defaults.IsRuleComparisonValid(eachRule.Comparison, requiredPath, requestPath) {
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
				case "header":
					// get provided header name
					requiredHeaderName := eachRule.Key
					// get provided header value
					requiredHeaderValue := eachRule.Value
					// check the incoming header
					incomingHeaderValue := r.Header.Get(requiredHeaderName)
					// compare incoming header
					if defaults.IsRuleComparisonValid(eachRule.Comparison, requiredHeaderValue, incomingHeaderValue) {
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
					// compare incoming cookie
					if defaults.IsRuleComparisonValid(eachRule.Comparison, requiredCookieValue, incomingCookie.Value) {
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
					// compare incoming source ip
					if defaults.IsRuleComparisonValid(eachRule.Comparison, requiredSourceIP, incomingSourceIP) {
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
					// compare incoming source port
					if defaults.IsRuleComparisonValid(eachRule.Comparison, requiredSourcePort, incomingSourcePort) {
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
					// compare incoming referrer
					if defaults.IsRuleComparisonValid(eachRule.Comparison, requiredReferrer, incomingReferrer) {
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
					// compare incoming method
					if defaults.IsRuleComparisonValid(eachRule.Comparison, requiredMethod, incomingMethod) {
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
					// compare incoming host
					if defaults.IsRuleComparisonValid(eachRule.Comparison, requiredHost, incomingHost) {
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

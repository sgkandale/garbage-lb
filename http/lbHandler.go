package http

import (
	"net/http"
)

func (server *HTTPServer) LBHandler(w http.ResponseWriter, r *http.Request) {

	if server.Listener.Filter != nil {
	rulesIterator:
		// iterate over the rules
		for _, eachRule := range server.Listener.Filter.Rules {
			// if the rule is enabled
			if eachRule.Enabled {
				// switch over the rule type
				switch eachRule.Type {
				case "path":
					// get provided path length
					requiredPathLength := len(eachRule.Value)
					// check the incoming path length
					if len(r.URL.Path) >= requiredPathLength {
						// compare incoming path with provided path
						if r.URL.Path[0:requiredPathLength] == eachRule.Value {
							// if rule action is reject, return
							if eachRule.Action == "reject" {
								rejectionHandler(&w, r)
								return
							}
							// if rule action is allow
							if eachRule.Action == "allow" {
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

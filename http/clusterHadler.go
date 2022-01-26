package http

import (
	goHttp "net/http"

	"garbagelb/config"
)

func clusterHadler(w *goHttp.ResponseWriter, r *goHttp.Request, cluster *config.Cluster) {

	switch cluster.Policy {
	// if csssluster policy is round-robin
	case "round_robin":
		totalEndpoints := len(cluster.Endpoints)
		if totalEndpoints == 0 {
			rejectionHandler(w, r)
			return
		}

		currentEndpointIndex := getCurrentEndpointIndex(cluster)
		if currentEndpointIndex == -1 {
			// all endpoints are unhealthy
			rejectionHandler(w, r)
			return
		}

		// forward the request to current endpoint
		forwardRequest(w, r, cluster.Endpoints[currentEndpointIndex])
		return
	// if cluster policy is random
	case "random":
		totalEndpoints := len(cluster.Endpoints)
		if totalEndpoints == 0 {
			rejectionHandler(w, r)
			return
		}

		randomEndpointIndex := getRandomEndpointIndex(cluster)
		if randomEndpointIndex == -1 {
			// all endpoints are unhealthy
			rejectionHandler(w, r)
			return
		}

		// forward the request to current endpoint
		forwardRequest(w, r, cluster.Endpoints[randomEndpointIndex])
		return
	// if cluster policy is least_connections
	case "least_connections":
		totalEndpoints := len(cluster.Endpoints)
		if totalEndpoints == 0 {
			rejectionHandler(w, r)
			return
		}

		leastConnectionsEndpointIndex := getLeastConnectionsEndpointIndex(cluster)
		if leastConnectionsEndpointIndex == -1 {
			// all endpoints are unhealthy
			rejectionHandler(w, r)
			return
		}

		// forward the request to current endpoint
		forwardRequest(w, r, cluster.Endpoints[leastConnectionsEndpointIndex])
		return
	default:
		rejectionHandler(w, r)
		return
	}
}

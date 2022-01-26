package http

import (
	"net/http"

	"garbagelb/config"
)

func clusterHadler(w *http.ResponseWriter, r *http.Request, cluster *config.Cluster) {

	switch cluster.Policy {
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
	default:
		rejectionHandler(w, r)
		return
	}
}

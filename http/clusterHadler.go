package http

import (
	goHttp "net/http"

	"garbagelb/config"
)

func clusterHadler(w *goHttp.ResponseWriter, r *goHttp.Request, cluster *config.Cluster) {

	// no endpoints in cluster
	totalEndpoints := len(cluster.Endpoints)
	if totalEndpoints == 0 {
		rejectionHandler(w, r)
		return
	}

	// target endpoint
	targetEndpoint := -1

	switch cluster.Policy {

	// if csssluster policy is round-robin
	case "round_robin":
		targetEndpoint = cluster.GetCurrentEndpointIndex()

	// if cluster policy is random
	case "random":
		targetEndpoint = cluster.GetRandomEndpointIndex()

	// if cluster policy is least_connections
	case "least_connections":
		targetEndpoint = cluster.GetLeastConnectionsEndpointIndex()

	default:
		rejectionHandler(w, r)
		return
	}

	if targetEndpoint == -1 {
		// all endpoints are unhealthy
		rejectionHandler(w, r)
		return
	}

	// forward the request to current endpoint
	forwardRequest(w, r, cluster.Endpoints[targetEndpoint])
}

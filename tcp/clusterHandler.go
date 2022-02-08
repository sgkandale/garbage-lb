package tcp

import (
	"log"
	"net"

	"garbagelb/config"
)

func clusterHadler(src net.Conn, cluster *config.Cluster) {

	// no endpoints in cluster
	totalEndpoints := len(cluster.Endpoints)
	if totalEndpoints == 0 {
		rejectionHandler(src)
		return
	}

	// target endpoint
	targetEndpoint := 0

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
		log.Println("unknown cluster policy : ", cluster.Policy)
		return
	}

	if targetEndpoint == -1 {
		// all endpoints are unhealthy
		rejectionHandler(src)
		return
	}

	// forward the request to current endpoint
	forwardRequest(src, cluster.Endpoints[targetEndpoint])
}

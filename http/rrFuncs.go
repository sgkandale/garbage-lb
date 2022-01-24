package http

import (
	"garbagelb/config"
)

func getCurrentEndpointIndex(cluster *config.Cluster) int {
	// iterate over the endpoints to check if all of them are unhealthy
	allEndpointsUnhealthy := true
	for _, endpoint := range cluster.Endpoints {
		if endpoint.Healthy {
			allEndpointsUnhealthy = false
			break
		}
	}

	// if all endpoints are unhealthy, return -1
	if allEndpointsUnhealthy {
		return -1
	}

	cluster.IncrementRREndpointIndex()
	if cluster.Endpoints[cluster.RREndpointIndex].Healthy {
		return cluster.RREndpointIndex
	} else {
		return getCurrentEndpointIndex(cluster)
	}
}

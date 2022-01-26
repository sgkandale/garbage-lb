package http

import (
	"garbagelb/config"
	"math/rand"
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

func getRandomEndpointIndex(cluster *config.Cluster) int {
	// iterate over the endpoints to check if all of them are unhealthy
	allEndpointsUnhealthy := true
	healthyEndpoints := []int{}
	for endpointIndex, endpoint := range cluster.Endpoints {
		if endpoint.Healthy {
			allEndpointsUnhealthy = false
			healthyEndpoints = append(healthyEndpoints, endpointIndex)
			break
		}
	}

	// if all endpoints are unhealthy, return -1
	if allEndpointsUnhealthy {
		return -1
	}

	// get a random endpoint index
	randomEndpointIndex := healthyEndpoints[rand.Intn(len(healthyEndpoints))]
	return randomEndpointIndex

}

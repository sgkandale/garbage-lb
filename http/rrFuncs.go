package http

import (
	"math/rand"

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

func getLeastConnectionsEndpointIndex(cluster *config.Cluster) int {

	leastConnectionsEndpointIndex := -1
	leastConnections := -1

	// iterate over the endpoints
	for endpointIndex, endpoint := range cluster.Endpoints {
		if endpoint.Healthy {
			if leastConnections == -1 || endpoint.ActiveConnectionCount < leastConnections {
				leastConnections = endpoint.ActiveConnectionCount
				leastConnectionsEndpointIndex = endpointIndex
			}
		}
	}

	return leastConnectionsEndpointIndex
}

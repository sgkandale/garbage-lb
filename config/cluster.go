package config

import (
	"math/rand"
	"sync"
)

type ClusterHealth struct {
	Status         string     `json:"status"`
	HealthyCount   int        `json:"healthyCount"`
	UnhealthyCount int        `json:"unhealthyCount"`
	DegradedCount  int        `json:"degradedCount"`
	Mutex          sync.Mutex `json:"-"`
}

type Cluster struct {
	Name            string         `json:"name,omitempty"`
	Policy          string         `json:"policy,omitempty"`
	Endpoints       []*Endpoint    `json:"endpoints,omitempty"`
	Health          *ClusterHealth `json:"health"`
	RREndpointIndex int            `json:"-"`
	RequestCounter  int64          `json:"requestCounter,omitempty"`
	Mutex           sync.Mutex     `json:"-"`
}

func (cluster *Cluster) GetCurrentEndpointIndex() int {
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
		return cluster.GetCurrentEndpointIndex()
	}
}

func (cluster *Cluster) GetRandomEndpointIndex() int {
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

func (cluster *Cluster) GetLeastConnectionsEndpointIndex() int {

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

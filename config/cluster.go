package config

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
)

type ClusterHealth struct {
	Status         string     `json:"status"`
	HealthyCount   int        `json:"healthyCount"`
	UnhealthyCount int        `json:"unhealthyCount"`
	DegradedCount  int        `json:"degradedCount"`
	Mutex          sync.Mutex `json:"-"`
}

type BasicAuth struct {
	Enabled      bool   `json:"enabled,omitempty"`
	Username     string `json:"username,omitempty"`
	UsernameHash string `json:"-"`
	Password     string `json:"password,omitempty"`
	PasswordHash string `json:"-"`
}

type Cluster struct {
	Name            string         `json:"name,omitempty"`
	Policy          string         `json:"policy,omitempty"`
	BasicAuth       *BasicAuth     `json:"basicAuth,omitempty"`
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

func (cluster *Cluster) IncrementRREndpointIndex() {
	// lock the endpoints
	cluster.Mutex.Lock()
	// increment the endpoint index
	cluster.RREndpointIndex = (cluster.RREndpointIndex + 1) % len(cluster.Endpoints)
	// increment the request counter
	cluster.RequestCounter++
	// unlock the endpoints
	cluster.Mutex.Unlock()
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

func (cluster *Cluster) getBasicAuth() (*BasicAuth, error) {
	// username checks
	if len(cluster.BasicAuth.Username) == 0 {
		return nil, fmt.Errorf(
			"basic auth username is empty for cluster {%s}",
			cluster.Name,
		)
	}
	if len(cluster.BasicAuth.Username) < 4 {
		return nil, fmt.Errorf(
			"basic auth username should be atleast 4 characters long for cluster {%s}",
			cluster.Name,
		)
	}
	if len(cluster.BasicAuth.Username) > 256 {
		return nil, fmt.Errorf(
			"basic auth username cannot be longer than 256 characters long for cluster {%s}",
			cluster.Name,
		)
	}

	// password checks
	if len(cluster.BasicAuth.Password) == 0 {
		return nil, fmt.Errorf(
			"basic auth password is empty for cluster {%s}",
			cluster.Name,
		)
	}
	if len(cluster.BasicAuth.Password) < 4 {
		return nil, fmt.Errorf(
			"basic auth password should be atleast 4 characters long for cluster {%s}",
			cluster.Name,
		)
	}
	if len(cluster.BasicAuth.Password) > 256 {
		return nil, fmt.Errorf(
			"basic auth password cannot be longer than 256 characters long for cluster {%s}",
			cluster.Name,
		)
	}

	newBasicAuth := &BasicAuth{
		Enabled:      cluster.BasicAuth.Enabled,
		Username:     cluster.BasicAuth.Username,
		UsernameHash: fmt.Sprintf("%x", sha256.Sum256([]byte(cluster.BasicAuth.Username))),
		Password:     cluster.BasicAuth.Password,
		PasswordHash: fmt.Sprintf("%x", sha256.Sum256([]byte(cluster.BasicAuth.Password))),
	}

	return newBasicAuth, nil
}

func (cluster *Cluster) IsBasicAuthValid(r *http.Request) bool {
	if cluster.BasicAuth != nil {
		if cluster.BasicAuth.Enabled {
			username, password, ok := r.BasicAuth()
			if !ok {
				return false
			}
			usernameHash := fmt.Sprintf("%x", sha256.Sum256([]byte(username)))
			passwordHash := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))

			if usernameHash != cluster.BasicAuth.UsernameHash || passwordHash != cluster.BasicAuth.PasswordHash {
				return false
			}
		}
	}
	return true
}

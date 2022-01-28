package config

import "time"

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

func (endpoint *Endpoint) SetUnhealthy() {
	endpoint.Mutex.Lock()
	endpoint.Healthy = false
	endpoint.Mutex.Unlock()
}

func (endpoint *Endpoint) SetHealthy() {
	endpoint.Mutex.Lock()
	endpoint.Healthy = false
	endpoint.LastSeen = time.Now().Unix()
	endpoint.Mutex.Unlock()
}

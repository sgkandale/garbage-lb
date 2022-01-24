package config

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

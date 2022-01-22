package config

func (cluster *Cluster) AddEndpoint(endpoint Endpoint) {
	cluster.Endpoints = append(cluster.Endpoints, endpoint)
}

package config

type Listener struct {
	ID            string
	Name          string
	Port          string
	TLS           bool
	CertPath      string
	KeyPath       string
	DomainName    string
	Type          string
	Listening     bool
	TargetCluster Cluster
}

type Endpoint struct {
	ID      string
	Name    string
	Address string
	Port    string
	Healthy bool
}

type ClusterHealth struct {
	Status         string
	HealthyCount   int
	UnhealthyCount int
	DegradedCount  int
}

type Cluster struct {
	ID        string
	Name      string
	Type      string
	Endpoints []Endpoint
	Health    ClusterHealth
}

type Admin struct {
	Port    string
	Enabled bool
}

type ConfigStruct struct {
	Admin     Admin
	Listeners []Listener
	Clusters  []Cluster
}

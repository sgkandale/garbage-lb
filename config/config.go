package config

type Listener struct {
	Name                 string
	Port                 string
	TLS                  bool
	CertPath             string
	KeyPath              string
	DomainName           string
	Type                 string
	Listening            bool
	TargetCluster        string
	TargetClusterDetails *Cluster
}

type Endpoint struct {
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
	Name      string
	Type      string
	Policy    string
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

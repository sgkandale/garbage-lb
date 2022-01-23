package config

type Listener struct {
	Name       string
	Port       int
	TLS        bool
	CertPath   string
	KeyPath    string
	DomainName string
	Type       string
	Listening  bool
	Filter     *Filter
}

type Endpoint struct {
	Name    string
	Address string
	Port    int
	Healthy bool
}

type ClusterHealth struct {
	Status         string
	HealthyCount   int
	UnhealthyCount int
	DegradedCount  int
}

type Cluster struct {
	Name string
	// Type      string
	Policy    string
	Endpoints []*Endpoint
	Health    *ClusterHealth
}

type Rule struct {
	Name          string
	Type          string
	Value         string
	Subvalue      string
	Action        string
	Enabled       bool
	Cluster       string
	TargetCluster *Cluster
}

type Filter struct {
	Name  string
	Rules []*Rule
}

type Admin struct {
	Port    int
	Enabled bool
}

type ConfigStruct struct {
	Admin     *Admin
	Listeners []*Listener
	Clusters  []*Cluster
}

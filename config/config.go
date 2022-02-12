package config

type Rule struct {
	Name          string   `json:"name,omitempty"`
	Type          string   `json:"type,omitempty"`
	Value         string   `json:"value,omitempty"`
	Key           string   `json:"key,omitempty"`
	Comparison    string   `json:"comparison,omitempty"`
	Action        string   `json:"action,omitempty"`
	Enabled       bool     `json:"enabled,omitempty"`
	Cluster       string   `json:"cluster,omitempty"`
	TargetCluster *Cluster `json:"targetCluster,omitempty"`
}

type Filter struct {
	Name  string  `json:"name,omitempty"`
	Rules []*Rule `json:"rules,omitempty"`
}

type Admin struct {
	Port    int  `json:"port,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
}

type ConfigStruct struct {
	Admin     *Admin      `json:"admin,omitempty"`
	Listeners []*Listener `json:"listeners,omitempty"`
	Clusters  []*Cluster  `json:"clusters,omitempty"`
}

var Config = &ConfigStruct{}

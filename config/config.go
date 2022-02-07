package config

import "sync"

type Listener struct {
	Name                string     `json:"name,omitempty"`
	Port                int        `json:"port,omitempty"`
	Type                string     `json:"type,omitempty"`
	TLS                 bool       `json:"tls,omitempty"`
	CertPath            string     `json:"certPath,omitempty"`
	KeyPath             string     `json:"keyPath,omitempty"`
	Listening           bool       `json:"listening,omitempty"`
	Filter              *Filter    `json:"filter,omitempty"`
	HealthCheckInterval int        `json:"healthCheckInterval,omitempty"`
	ActiveConnections   int64      `json:"activeConnections,omitempty"`
	MaxConnections      int64      `json:"maxConnections,omitempty"`
	Mutex               sync.Mutex `json:"-"`
}

type Endpoint struct {
	Name                  string     `json:"name,omitempty"`
	Address               string     `json:"address,omitempty"`
	Port                  int        `json:"port,omitempty"`
	Protocol              string     `json:"protocol,omitempty"`
	Healthy               bool       `json:"healthy,omitempty"`
	ActiveConnectionCount int        `json:"activeConnectionCount,omitempty"`
	TotalRequestCount     int        `json:"totalRequestCount,omitempty"`
	LastSeen              int64      `json:"lastSeen"`
	Mutex                 sync.Mutex `json:"-"`
}

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

type Rule struct {
	Name          string   `json:"name,omitempty"`
	Type          string   `json:"type,omitempty"`
	Value         string   `json:"value,omitempty"`
	Key           string   `json:"key,omitempty"`
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

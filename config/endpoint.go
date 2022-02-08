package config

import (
	"sync"
	"time"
)

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

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

func (listener *Listener) IncrementActiveConnections() {
	listener.Mutex.Lock()
	listener.ActiveConnections++
	listener.Mutex.Unlock()
}

func (listener *Listener) DecrementActiveConnections() {
	listener.Mutex.Lock()
	listener.ActiveConnections--
	listener.Mutex.Unlock()
}

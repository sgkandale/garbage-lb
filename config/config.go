package config

type serverConfig struct {
	Port     string
	TLS      bool
	CertPath string
	KeyPath  string
}

type adminPortal struct {
	// Enabled bool
	Port string
}

type ConfigStruct struct {
	Server      serverConfig
	AdminPortal adminPortal
}

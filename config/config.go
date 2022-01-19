package config

type serverConfig struct {
	Port     string
	TLS      bool
	CertPath string
	KeyPath  string
}

type Config struct {
	Server serverConfig
}

package config

type Listener struct {
	Port       string
	TLS        bool
	CertPath   string
	KeyPath    string
	DomainName string
}

type Admin struct {
	Port string
}

type ConfigStruct struct {
	Admin     Admin
	Listeners []Listener
}

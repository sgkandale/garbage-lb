package defaults

type listenerType struct {
	Name        string
	TLSRequired bool
}

var ListenerTypes = []listenerType{
	{
		Name:        "http",
		TLSRequired: false,
	},
	{
		Name:        "tcp",
		TLSRequired: false,
	},
	{
		Name:        "tcp4",
		TLSRequired: false,
	},
	{
		Name:        "tcp6",
		TLSRequired: false,
	},
}

const ListenerHealthCheckInterval = 5

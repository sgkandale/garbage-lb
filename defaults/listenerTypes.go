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
		Name:        "https",
		TLSRequired: true,
	},
}

const ListenerHealthCheckInterval = 5

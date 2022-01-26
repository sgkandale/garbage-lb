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
}

const ListenerHealthCheckInterval = 5

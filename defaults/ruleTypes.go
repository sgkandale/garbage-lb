package defaults

type ruleRequirement struct {
	Name          string
	ValueRequired bool
	KeyRequired   bool
}

var RuleTypes = map[string][]ruleRequirement{
	"http": {
		{
			Name:          "path",
			ValueRequired: true,
			KeyRequired:   false,
		},
		{
			Name:          "header",
			ValueRequired: true,
			KeyRequired:   true,
		},
		{
			Name:          "cookie",
			ValueRequired: true,
			KeyRequired:   true,
		},
		{
			Name:          "source_ip",
			ValueRequired: true,
			KeyRequired:   false,
		},
		{
			Name:          "source_port",
			ValueRequired: true,
			KeyRequired:   false,
		},
		{
			Name:          "referrer",
			ValueRequired: true,
			KeyRequired:   false,
		},
		{
			Name:          "referer",
			ValueRequired: true,
			KeyRequired:   false,
		},
		{
			Name:          "method",
			ValueRequired: true,
			KeyRequired:   false,
		},
		{
			Name:          "host",
			ValueRequired: true,
			KeyRequired:   false,
		},
	},
	"tcp": {
		{
			Name:          "tcp_check",
			ValueRequired: false,
			KeyRequired:   false,
		},
	},
}

var RuleActionValues = []string{
	"forward",
	"reject",
}

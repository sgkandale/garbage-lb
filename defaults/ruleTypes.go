package defaults

type ruleRequirement struct {
	Name          string
	ValueRequired bool
	KeyRequired   bool
}

var RuleTypes = []ruleRequirement{
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
	},
	{
		Name:          "source_port",
		ValueRequired: true,
	},
	{
		Name:          "referrer",
		ValueRequired: true,
	},
	{
		Name:          "referer",
		ValueRequired: true,
	},
}

var RuleActionValues = []string{
	"forward",
	"reject",
}

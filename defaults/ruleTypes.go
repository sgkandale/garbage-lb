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
}

var RuleActionValues = []string{
	"allow",
	"reject",
}

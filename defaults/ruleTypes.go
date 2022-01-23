package defaults

type ruleRequirement struct {
	Name             string
	ValueRequired    bool
	SubvalueRequired bool
}

var RuleTypes = []ruleRequirement{
	{
		Name:             "path",
		ValueRequired:    true,
		SubvalueRequired: false,
	},
}

var RuleActionValues = []string{
	"allow",
	"reject",
}

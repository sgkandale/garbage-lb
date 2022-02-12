package defaults

import "strings"

var RuleComparisons = []string{
	"equals",
	"not_equals",
	"contains",
	"not_contains",
	"starts_with",
	"ends_with",
	"matches",
	"not_matches",
	"greater_than",
	"less_than",
	"greater_than_or_equal",
	"less_than_or_equal",
	"not_required",
}

func IsRuleComparisonValid(givenComparison, givenValue, incomingValue string) bool {
	switch givenComparison {
	case "equals":
		return givenValue == incomingValue
	case "not_equals":
		return givenValue != incomingValue
	case "contains":
		return strings.Contains(incomingValue, givenValue)
	case "not_contains":
		return !strings.Contains(incomingValue, givenValue)
	case "starts_with":
		return strings.HasPrefix(incomingValue, givenValue)
	case "ends_with":
		return strings.HasSuffix(incomingValue, givenValue)
	case "matches":
		return strings.EqualFold(incomingValue, givenValue)
	case "not_matches":
		return !strings.EqualFold(incomingValue, givenValue)
	case "greater_than":
		return incomingValue > givenValue
	case "less_than":
		return incomingValue < givenValue
	case "greater_than_or_equal":
		return incomingValue >= givenValue
	case "less_than_or_equal":
		return incomingValue <= givenValue
	case "not_required":
		return true
	}
	return false
}

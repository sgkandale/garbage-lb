package config

import "log"

func (filter *Filter) AddRule(rule *Rule, preparedConfig ConfigStruct) {

	if rule.Name == "" {
		log.Fatal("no rule name specified for ")
	}
}

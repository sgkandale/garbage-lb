package config

import (
	"fmt"
	"strings"

	"garbagelb/defaults"
)

func (configStruct *ConfigStruct) AddListener(givenListener *Listener) error {

	newListener := &Listener{}

	if givenListener.Name == "" {
		return fmt.Errorf("Listener Name cannot be empty")
	}

	for _, eachListener := range configStruct.Listeners {
		// check case insensitive name
		if strings.EqualFold(eachListener.Name, givenListener.Name) {
			return fmt.Errorf("listener name {%s} is already in use", givenListener.Name)
		}
	}
	newListener.Name = givenListener.Name

	if givenListener.Port < 1 || givenListener.Port > 65535 {
		return fmt.Errorf(
			"out of range port {%d} for listener {%s}",
			givenListener.Port,
			givenListener.Name,
		)
	}
	newListener.Port = givenListener.Port

	// TLS checks pending

	for _, listenerType := range defaults.ListenerTypes {
		if strings.EqualFold(listenerType, givenListener.Type) {
			newListener.Type = listenerType
		}
	}
	if newListener.Type == "" {
		return fmt.Errorf(
			"unsupported listener type {%s} for listener {%s}",
			givenListener.Type,
			givenListener.Name,
		)
	}

	newListener.Listening = givenListener.Listening

	// filter checks
	if givenListener.Filter == nil {
		return fmt.Errorf(
			"no filter provided for listener {%s}",
			givenListener.Name,
		)
	}
	if givenListener.Filter.Name == "" {
		return fmt.Errorf(
			"filter name is empty for listener {%s}",
			givenListener.Name,
		)
	}
	newFilter := &Filter{
		Name: givenListener.Filter.Name,
	}
	// rules checks
	for eachRuleIndex, eachRule := range givenListener.Filter.Rules {
		newRule := &Rule{}
		// check rule name validity
		if eachRule.Name == "" {
			return fmt.Errorf(
				"rule name is empty at index {%d} in filter {%s} in listener {%s}",
				eachRuleIndex,
				givenListener.Filter.Name,
				givenListener.Name,
			)
		}
		// check for existing rule name
		for existingRuleIndex, existingRule := range newFilter.Rules {
			if strings.EqualFold(existingRule.Name, eachRule.Name) {
				return fmt.Errorf(
					`duplicate rule name {%s} ::::
						 trace : 
						 listener {%s}
						 filter {%s}
						 first use at rule {%d}
						 second use at rule {%d}`,
					eachRule.Name,
					givenListener.Name,
					givenListener.Filter.Name,
					existingRuleIndex,
					eachRuleIndex,
				)
			}
		}
		newRule.Name = eachRule.Name
		// check rule types and values
		for _, eachRuleType := range defaults.RuleTypes {
			if strings.EqualFold(eachRuleType.Name, eachRule.Type) {
				newRule.Type = eachRuleType.Name
				if eachRuleType.ValueRequired && eachRule.Value == "" {
					return fmt.Errorf(
						`value required for rule type {%s} ::::
							 trace :
							 listener {%s}
							 filter {%s}
							 rule {%d}`,
						eachRule.Type,
						givenListener.Name,
						givenListener.Filter.Name,
						eachRuleIndex,
					)
				}
				newRule.Value = eachRule.Value
				if eachRuleType.KeyRequired && eachRule.Key == "" {
					return fmt.Errorf(
						`key required for rule type {%s} ::::
							 trace :
							 listener {%s}
							 filter {%s}
							 rule {%d}`,
						eachRule.Type,
						givenListener.Name,
						givenListener.Filter.Name,
						eachRuleIndex,
					)
				}
				newRule.Key = eachRule.Key
			}
		}
		if newRule.Type == "" {
			return fmt.Errorf(
				`invalid rule type {%s} ::::
					 trace :
					 listener {%s}
					 filter {%s}
					 rule {%d}`,
				eachRule.Type,
				givenListener.Name,
				givenListener.Filter.Name,
				eachRuleIndex,
			)
		}
		// rule action checks
		for _, eachAction := range defaults.RuleActionValues {
			if strings.EqualFold(eachAction, eachRule.Action) {
				newRule.Action = eachRule.Action
			}
		}
		if newRule.Action == "" {
			return fmt.Errorf(
				`invalid action {%s} ::::
					 trace :
					 listener {%s}
					 filter {%s}
					 rule {%d}`,
				eachRule.Action,
				givenListener.Name,
				givenListener.Filter.Name,
				eachRuleIndex,
			)
		}
		// cluster checks
		for _, existingCluster := range configStruct.Clusters {
			if existingCluster.Name == eachRule.Cluster {
				newRule.Cluster = eachRule.Cluster
				newRule.TargetCluster = existingCluster
			}
		}
		if newRule.Cluster == "" || newRule.TargetCluster == nil {
			return fmt.Errorf(
				`invalid cluster name {%s} ::::
					 trace :
					 listener {%s}
					 filter {%s}
					 rule {%d}`,
				eachRule.Cluster,
				givenListener.Name,
				givenListener.Filter.Name,
				eachRuleIndex,
			)
		}
		newRule.Enabled = eachRule.Enabled
		newFilter.Rules = append(newFilter.Rules, newRule)
	}

	newListener.Filter = newFilter

	// add newListener to config
	configStruct.Listeners = append(configStruct.Listeners, newListener)

	return nil

}

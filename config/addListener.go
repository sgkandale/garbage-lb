package config

import (
	"fmt"
	"strings"

	"garbagelb/internal/defaults"
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

	// checks for health check interval
	newListener.HealthCheckInterval = givenListener.HealthCheckInterval
	if newListener.HealthCheckInterval == 0 {
		newListener.HealthCheckInterval = defaults.ListenerHealthCheckInterval
	} else if newListener.HealthCheckInterval < 0 {
		return fmt.Errorf(
			"invalid health check interval {%d} for listener {%s}",
			newListener.HealthCheckInterval,
			newListener.Name,
		)
	} else if newListener.HealthCheckInterval > 120 {
		return fmt.Errorf(
			"health check interval {%d} is too long for listener {%s}",
			newListener.HealthCheckInterval,
			newListener.Name,
		)
	}

	// max connections checks
	newListener.MaxConnections = givenListener.MaxConnections
	if newListener.MaxConnections < 0 {
		return fmt.Errorf(
			"invalid max connections {%d} for listener {%s}",
			newListener.MaxConnections,
			newListener.Name,
		)
	}

	// type checks
	for _, listenerType := range defaults.ListenerTypes {
		if strings.EqualFold(listenerType.Name, givenListener.Type) {
			newListener.Type = listenerType.Name
			// TLS checks
			newListener.TLS = givenListener.TLS
			newListener.CertPath = givenListener.CertPath
			newListener.KeyPath = givenListener.KeyPath
			if newListener.TLS && newListener.CertPath == "" {
				return fmt.Errorf(
					"listener {%s} has TLS enabled but certPath is not provided",
					givenListener.Name,
				)
			}
			if newListener.TLS && newListener.KeyPath == "" {
				return fmt.Errorf(
					"listener {%s} has TLS enabled but keyPath is not provided",
					givenListener.Name,
				)
			}
			break
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
		listenerRuleTypes := defaults.RuleTypes[newListener.Type]
		if listenerRuleTypes == nil {
			return fmt.Errorf(
				"unsupported listener type {%s} for listener {%s}",
				newListener.Type,
				newListener.Name,
			)
		}
		for _, eachRuleType := range listenerRuleTypes {
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
		// rule comparison checks
		for _, eachRuleComparison := range defaults.RuleComparisons {
			if strings.EqualFold(eachRuleComparison, eachRule.Comparison) {
				newRule.Comparison = eachRuleComparison
				break
			}
		}
		if newRule.Comparison == "" {
			return fmt.Errorf(
				`unsupported comparison {%s}  ::::
					trace : 
					listener {%s}
					filter {%s}
					rule {%s}`,
				eachRule.Comparison,
				givenListener.Name,
				givenListener.Filter.Name,
				eachRule.Name,
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

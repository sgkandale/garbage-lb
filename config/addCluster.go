package config

import (
	"fmt"
	"strings"

	"garbagelb/internal/defaults"
)

func (configStruct *ConfigStruct) AddCluster(givenCluster *Cluster) error {

	newCluster := &Cluster{}

	if givenCluster.Name == "" {
		return fmt.Errorf("cluster name cannot be empty")
	}

	// check for existing cluster name
	for eachClusterIndex, eachCluster := range configStruct.Clusters {
		// check case insensitive name
		if strings.EqualFold(eachCluster.Name, givenCluster.Name) {
			return fmt.Errorf(
				"cluster name {%s} is already defined at index {%d} as {%s}",
				givenCluster.Name,
				eachClusterIndex,
				eachCluster.Name,
			)
		}
	}
	newCluster.Name = givenCluster.Name

	// check for cluster policy
	for _, eachPolicy := range defaults.ClusterPolicies {
		if eachPolicy == givenCluster.Policy {
			newCluster.Policy = givenCluster.Policy
		}
	}
	if newCluster.Policy == "" {
		return fmt.Errorf(
			"unsupported cluster policy {%s} for cluster {%s}",
			givenCluster.Policy,
			givenCluster.Name,
		)
	}

	// endpoint checks
	for givenEndpointIndex, givenEndpoint := range givenCluster.Endpoints {
		newEndpoint := &Endpoint{}
		// endpoint name validity
		if givenEndpoint.Name == "" {
			return fmt.Errorf(
				"endpoint name is empty at index {%d} in cluster {%s}",
				givenEndpointIndex,
				givenCluster.Name,
			)
		}
		// check for duplicate endpoint name
		for existingEndpointIndex, existingEndpoint := range newCluster.Endpoints {
			if strings.EqualFold(existingEndpoint.Name, givenEndpoint.Name) {
				return fmt.Errorf(
					"endpoint name {%s} is already defined at index {%d} as {%s} in cluster {%s}",
					givenEndpoint.Name,
					existingEndpointIndex,
					existingEndpoint.Name,
					givenCluster.Name,
				)
			}
		}
		newEndpoint.Name = givenEndpoint.Name
		newEndpoint.Address = givenEndpoint.Address
		// check endpoint port range
		if givenEndpoint.Port < 1 || givenEndpoint.Port > 65535 {
			return fmt.Errorf(
				"port out of range for endpoint at index {%d} in cluster {%s}",
				givenEndpointIndex,
				givenCluster.Name,
			)
		}
		// endpoint protocol checks
		newEndpoint.Protocol = defaults.GetEndpointProtocol(givenEndpoint.Protocol)
		if newEndpoint.Protocol == "" {
			return fmt.Errorf(
				"unsupported protocol {%s} for endpoint at index {%d} in cluster {%s}",
				givenEndpoint.Protocol,
				givenEndpointIndex,
				givenCluster.Name,
			)
		}
		newEndpoint.Port = givenEndpoint.Port
		newEndpoint.Healthy = false
		newCluster.Endpoints = append(newCluster.Endpoints, newEndpoint)
	}

	newCluster.Health = &ClusterHealth{
		Status:         "Unknown",
		HealthyCount:   0,
		UnhealthyCount: len(newCluster.Endpoints),
		DegradedCount:  0,
	}

	configStruct.Clusters = append(configStruct.Clusters, newCluster)

	return nil
}

package config

import (
	"log"
	"strings"

	"garbagelb/defaults"
)

func (configStruct *ConfigStruct) AddCluster(cluster *Cluster) {

	newCluster := &Cluster{}

	if cluster.Name == "" {
		log.Fatal("Cluster Name cannot be empty")
	}

	// check for existing cluster name
	for _, eachCluster := range configStruct.Clusters {
		// check case insensitive name
		if strings.EqualFold(eachCluster.Name, cluster.Name) {
			log.Fatal("Cluster name already used {%s}" + cluster.Name)
		}
	}

	// check for cluster policy
	for _, eachPolicy := range defaults.ClusterPolicies {
		if eachPolicy == cluster.Policy {
			newCluster.Policy = cluster.Policy
		}
	}
	if newCluster.Policy == "" {
		log.Fatalf("invalid cluster policy {%s} for cluster {%s}", cluster.Policy, cluster.Name)
	}

	for _, eachEndpoint := range cluster.Endpoints {
		newCluster.AddEndpoint(eachEndpoint)
	}

	configStruct.Clusters = append(configStruct.Clusters, cluster)

}

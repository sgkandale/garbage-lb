package config

import "log"

func (configStruct *ConfigStruct) AddCluster(cluster *Cluster) {

	newCluster := &Cluster{}

	if cluster.Name == "" {
		log.Fatal("Cluster Name cannot be empty")
	}

	for _, eachCluster := range configStruct.Clusters {
		if eachCluster.Name == cluster.Name {
			log.Fatal("Cluster name already used : " + cluster.Name)
		}
	}

	for _, eachEndpoint := range cluster.Endpoints {
		newCluster.AddEndpoint(eachEndpoint)
	}

	configStruct.Clusters = append(configStruct.Clusters, cluster)

}

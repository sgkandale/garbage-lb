package tcp

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"garbagelb/config"
	"garbagelb/defaults"
)

func pingEndpoint(endpoint *config.Endpoint) error {

	testConn, err := net.DialTimeout(
		endpoint.Protocol,
		fmt.Sprintf("%s:%d", endpoint.Address, endpoint.Port),
		10*time.Second,
	)
	if testConn != nil {
		testConn.Close()
	}
	return err
}

func PerformHealthChecks(listener *config.Listener) {

	if listener.Filter != nil {
		for _, rule := range listener.Filter.Rules {
			if rule.Enabled {
				if rule.TargetCluster != nil {
					endpointsWaitGroup := sync.WaitGroup{}
					endpointsWaitGroup.Add(len(rule.TargetCluster.Endpoints))
					for _, endpoint := range rule.TargetCluster.Endpoints {
						go func(endpoint *config.Endpoint) {
							defer endpointsWaitGroup.Done()
							err := pingEndpoint(endpoint)
							if err != nil {
								log.Println(err)
								endpoint.SetUnhealthy()
							} else {
								endpoint.SetHealthy()
							}
						}(endpoint)
					}
					endpointsWaitGroup.Wait()
					totalEndpoints := 0
					if rule.TargetCluster.Endpoints != nil {
						totalEndpoints = len(rule.TargetCluster.Endpoints)
					}
					healthyEndpoints := 0
					unhealthyEndpoints := 0
					degradedCount := 0
					for _, endpoint := range rule.TargetCluster.Endpoints {
						if endpoint.Healthy {
							healthyEndpoints++
						} else {
							if endpoint.LastSeen < time.Now().Unix()-defaults.EndpointDegradationInterval {
								degradedCount++
							} else {
								unhealthyEndpoints++
							}
						}
					}
					// get the percentage of healthy endpoints
					var healthyPercentage float64 = 0
					if totalEndpoints > 0 {
						healthyPercentage = float64(healthyEndpoints) / float64(totalEndpoints)
					}
					// lock the cluster health
					rule.TargetCluster.Health.Mutex.Lock()
					if healthyPercentage < defaults.HealthyClusterThreshold {
						rule.TargetCluster.Health.Status = "Unhealthy"
					} else {
						rule.TargetCluster.Health.Status = "Healthy"
					}
					rule.TargetCluster.Health.HealthyCount = healthyEndpoints
					rule.TargetCluster.Health.UnhealthyCount = unhealthyEndpoints
					rule.TargetCluster.Health.DegradedCount = degradedCount
					// unlock the cluster health
					rule.TargetCluster.Health.Mutex.Unlock()
				}
			}
		}
	}

}

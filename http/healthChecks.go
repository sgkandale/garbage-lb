package http

import (
	"fmt"
	"log"
	goHttp "net/http"
	"sync"
	"time"

	"garbagelb/config"
)

var client = goHttp.Client{
	Timeout: 10 * time.Second,
}

func pingEndpoint(endpoint *config.Endpoint) (int, error) {

	req, err := goHttp.NewRequest(
		"HEAD",
		fmt.Sprintf(
			"%s://%s:%d",
			endpoint.Protocol,
			endpoint.Address,
			endpoint.Port,
		),
		nil,
	)
	if err != nil {
		return 0, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}

	resp.Body.Close()

	return resp.StatusCode, nil
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
							statusCode, err := pingEndpoint(endpoint)
							if err != nil {
								log.Println(err)
								endpoint.Healthy = false
							} else {
								if statusCode == goHttp.StatusServiceUnavailable {
									log.Printf(
										"endpoint responded with status code %d\n",
										statusCode,
									)
									endpoint.Healthy = false
								} else {
									endpoint.Healthy = true
								}
							}
						}(endpoint)
					}
				}
			}
		}
	}

}

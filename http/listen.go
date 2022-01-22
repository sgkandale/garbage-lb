package http

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"garbagelb/config"
)

func (server *HTTPServer) Listen(wg *sync.WaitGroup, listener *config.Listener) {

	server.Addr = fmt.Sprintf(":%d", listener.Port)

	targetCluster := &config.Cluster{}

	targetClusterName := listener.TargetCluster
	for _, cluster := range config.Config.Clusters {
		if cluster.Name == targetClusterName {
			targetCluster = cluster
		}
	}

	switch strings.ToLower(targetCluster.Policy) {
	case "round_robin":
		server.Handler = http.HandlerFunc(server.LBHandler)
	default:
		log.Printf(
			"unsupported cluster policy : %s",
			targetCluster.Policy,
		)
		wg.Done()
		return
	}

	log.Println(
		fmt.Sprintf(
			"%s starting at port %d...",
			listener.Name,
			listener.Port,
		),
	)

	go func() {
		defer wg.Done()
		log.Fatal(server.ListenAndServe())
	}()
}

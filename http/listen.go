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

	server.Addr = fmt.Sprintf(":%s", listener.Port)

	targetCluster := config.Cluster{}

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
		panic(
			fmt.Sprintf(
				"unsupported cluster policy : %s",
				targetCluster.Policy,
			),
		)
	}

	log.Println(
		fmt.Sprintf(
			"%s starting at port %s...",
			listener.Name,
			listener.Port,
		),
	)

	go func() {
		defer wg.Done()
		log.Fatal(server.ListenAndServe())
	}()
}

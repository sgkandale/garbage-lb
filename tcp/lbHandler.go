package tcp

import (
	"log"
	"net"
)

func (server *TCPServer) LBHandler(src net.Conn) {

	// get max connections count
	maxAllowedConnections := server.Listener.MaxConnections
	if maxAllowedConnections > 0 {
		// compare active connections count
		if server.Listener.ActiveConnections >= maxAllowedConnections {
			log.Println("max connections reached")
			return
		} else {
			// increment active connections count
			server.Listener.Mutex.Lock()
			server.Listener.ActiveConnections++
			server.Listener.Mutex.Unlock()
		}

		// decrement active connections on exit
		defer func() {
			server.Listener.Mutex.Lock()
			server.Listener.ActiveConnections--
			server.Listener.Mutex.Unlock()
		}()
	}

	if server.Listener.Filter != nil {
	rulesIterator:
		// iterate over the rules
		for _, eachRule := range server.Listener.Filter.Rules {
			// if the rule is enabled
			if eachRule.Enabled {
				// switch over the rule type
				switch eachRule.Type {
				case "tcp_check":
					clusterHadler(src, eachRule.TargetCluster)
				default:
					log.Printf("unknown rule type {%s}", eachRule.Type)
					return
				}
			}
			// continue to next rule
			continue rulesIterator
		}
	}
}

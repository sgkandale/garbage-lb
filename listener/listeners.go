package listener

import (
	"sync"

	"garbagelb/config"
)

type ServerHandler interface {
	Listen(wg *sync.WaitGroup, listener *config.Listener)
	Terminate(wg *sync.WaitGroup, listener *config.Listener)
}

type Listener struct {
	ServerHandler   ServerHandler
	ListenerDetails *config.Listener
}

var Listeners []*Listener

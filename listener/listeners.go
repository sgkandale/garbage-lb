package listener

import "sync"

type ServerHandler interface {
	Listen(wg *sync.WaitGroup)
	Terminate(wg *sync.WaitGroup)
}

type Listener struct {
	Name          string
	Port          string
	ServerHandler ServerHandler
}

var Listeners []*Listener

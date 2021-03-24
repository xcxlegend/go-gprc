package ws

import "sync"

type ClientManager struct {
	clients sync.Map
}

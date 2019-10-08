package swatcher_server

import (
	"swatcher-server/pkg/store"
)

type Client struct {
	serverEvents chan store.Event
	clientEvents chan store.Event
}

type SWServerImpl struct {
	pending       map[int64]chan interface{}
	activeClients map[int64]chan *Client
	store         store.Storage
}

func NewSWServerImpl(storage store.Storage) *SWServerImpl {
	return &SWServerImpl{
		activeClients: make(map[int64]chan *Client),
		store:         storage,
	}
}

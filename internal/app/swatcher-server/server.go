package swatcher_server

import (
	"github.com/segmentio/kafka-go"
	"swatcher-server/pkg/store"
)

type Client struct {
	serverEvents chan store.Event
	clientEvents chan store.Event
}

type SWServerImpl struct {
	pending       map[int64]chan interface{}
	activeClients map[int64]chan *Client
	rClientEvents *kafka.Reader
	wClientEvents *kafka.Writer
	store         store.Storage
}

func NewSWServerImpl(storage store.Storage, rClientEvents *kafka.Reader, wClientEvents *kafka.Writer) *SWServerImpl {
	return &SWServerImpl{
		activeClients: make(map[int64]chan *Client),
		rClientEvents: rClientEvents,
		wClientEvents: wClientEvents,
		store:         storage,
	}
}

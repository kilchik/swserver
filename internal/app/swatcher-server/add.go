package swatcher_server

import (
	"context"
	"net/http"
	"swatcher-server/pkg/store"
)

type addMessage struct {
	Title  string `json:"title"`
	Magnet string `json:"magnet"`
}

func (s *SWServerImpl) Add(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var uid int64 = 42

	// TODO: check if already in db
	// TODO: insert into movies table and get uuid

	s.store.AddSrvEvent(ctx, uid, &store.Event{Type: store.SrvEventTypeAdd})
}

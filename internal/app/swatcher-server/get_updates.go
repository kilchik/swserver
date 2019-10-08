package swatcher_server

import (
	"context"
	"database/sql"
	pb "github.com/kilchik/swserver/pkg/swatcher-server"
	"github.com/pkg/errors"
)

func (s *SWServerImpl) GetUpdates(ctx context.Context, req *pb.GetUpdatesRequest) (resp *pb.GetUpdatesResponse, err error) {
	// TODO: auth
	for {
		events, err := s.store.GetNewSrvEvents(ctx, req.GetClientId(), req.GetRev())
		if err != nil {
			if err == sql.ErrNoRows {
				<-s.pending[req.GetClientId()]
				continue
			}
			return nil, errors.Wrap(err, "get updates: get new server event from db")
		}

		var updates []*pb.ServerUpdate
		for _, e := range events {
			updates = append(updates, &pb.ServerUpdate{
				Type: pb.ServerUpdate_UpdateType(e.Type),
				Data: string(e.Data),
			})
		}
		return &pb.GetUpdatesResponse{Updates: updates}, nil
	}
}

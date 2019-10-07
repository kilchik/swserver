package swatcher_server

import (
	"context"
	"encoding/json"
	"github.com/kilchik/logo/pkg/logo"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
	"io"
	"swatcher-server/pkg/store"
	pb "swatcher-server/pkg/swatcher-server"
)

// TODO: debug log interceptor
// TODO: why no context in stream?
func (s *SWServerImpl) SendUpdates(req pb.SwatcherServer_SendUpdatesServer) error {
	ctx := context.Background()
	for {
		update, err := req.Recv()
		if err == io.EOF {
			logo.Debug(ctx, "send updates: accepted stream message from client")
			return nil
		}
		if err != nil {
			logo.Error(ctx, "send updates: receive request stream: %v", err)
			return errors.Wrap(err, "failed to receive request stream")
		}

		event, err := json.Marshal(store.Event{
			Type: int32(update.Type),
			Data: []byte(update.Data),
		})
		if err != nil {
			logo.Error(ctx, "send updates: encode event from request: %v", err)
			return errors.Wrap(err, "failed to encode event from request")
		}

		if err := s.wClientEvents.WriteMessages(context.Background(), kafka.Message{Key: nil, Value: event}); err != nil {
			logo.Error(ctx, "send updates: write message to kafka: %v", err)
			return errors.Wrap(err, "failed to write message to kafka")
		}
	}
}

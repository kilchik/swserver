package main

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/kilchik/logo/pkg/logo"
	pb "github.com/kilchik/swserver/pkg/swatcher-server"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	swatcher_server "swatcher-server/internal/app/swatcher-server"
	"swatcher-server/pkg/config"
	"swatcher-server/pkg/store"
)

func main() {
	ctx := context.Background()

	// Initialize config
	config, err := config.Init("swatcher-server.toml")
	if err != nil {
		fmt.Printf("init config: %v\n", err)
		os.Exit(1)
	}

	// Initialize logger
	logo.Init(config.GetEnableDebugLogs())

	fatal := func(ctx context.Context, format string, args ...interface{}) {
		logo.Error(ctx, format, args...)
		os.Exit(1)
	}

	// Initiate storage
	db, err := sqlx.Connect("pgx", config.GetDatabaseDsn())
	if err != nil {
		fatal(ctx, "connect to db: %v", err)
	}
	storage := store.NewStorage(db)

	srv := swatcher_server.NewSWServerImpl(storage)

	// Create and start HTTP server
	http.HandleFunc("/add", srv.Add)
	go func() {
		if err := http.ListenAndServe(config.GetListenHttpAddr(), nil); err != nil {
			fatal(ctx, "listen http: %v", err)
		}
	}()

	// Create and start gRPC server
	listener, err := net.Listen("tcp", config.GetListenGrpcAddr())
	if err != nil {
		fatal(ctx, "listen grpc: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterSwatcherServerServer(grpcServer, srv)
	if err := grpcServer.Serve(listener); err != nil {
		fatal(ctx, "grpc server run: %v", err)
	}
}

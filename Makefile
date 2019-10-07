migrate:
	goose -dir migrations/ postgres "host=localhost port=5432 dbname=swdb user=testusr password=123 sslmode=disable" up sql

migrate-new:
	goose -dir migrations/ postgres "host=localhost port=5432 dbname=swdb user=testusr password=123 sslmode=disable" create $(name) sql

build:
	migrate
	protoc --proto_path=. --go_out=plugins=grpc:pkg/swatcher-server swatcher-server.proto
	go generate ./...
	go build

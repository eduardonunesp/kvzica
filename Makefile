.phony: proto

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	pkg/proto/*.proto

build:
	go build -o bin/kvzica cmd/kvzica/main.go

run:
	go run cmd/kvzica/main.go

client:
	go run cmd/kvclient/main.go
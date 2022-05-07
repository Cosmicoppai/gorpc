gen:
	protoc --proto_path proto --go_out=:pb --go-grpc_out=:pb --go_opt=paths=import proto/*.proto
clean:
	rm --r src/proto

server:
	go run cmd/server/main.go --port 8096

client:
	go run cmd/client/main.go --address 0.0.0.0:8096

test:
	go test --cover -race ./...
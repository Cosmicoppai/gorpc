gen:
	protoc --proto_path pb --go_out=src --go_opt=paths=import pb/*.proto
clean:
	rm --r src/proto

run:
	go run src/main.go
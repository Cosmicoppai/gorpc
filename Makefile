gen:
	protoc --proto_path proto --go_out=src --go_opt=paths=import proto/*.proto
clean:
	rm --r src/proto

run:
	go run src/main.go
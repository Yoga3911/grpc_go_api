proto:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto

server:
	go run main.go

evans:
	C:\\Evans\evans.exe --host localhost --port 9090 -r repl

.PHONY: proto server evans
gen-protoc:
	protoc --proto_path=. --go_out=. --go-grpc_out=paths=source_relative:. --go_opt=paths=source_relative todo/todo.proto
run-server:
	go run server/main.go
run-client:
	go run client/main.go
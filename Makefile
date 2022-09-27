
run:
	go run ./server port 50052

gen:
	protoc -I proto/ proto/service.proto --go_out=. --go-grpc_out=.
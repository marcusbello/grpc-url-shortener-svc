

gen-rpc:
	protoc -I proto/ proto/service.proto --go_out=. --go-grpc_out=.
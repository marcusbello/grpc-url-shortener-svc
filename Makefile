
runapp:
	docker-compose up

run:
	go run ./server port 50052

gen:
	protoc -I proto/ proto/service.proto --go_out=. --go-grpc_out=.


server:
	docker build -f Dockerfile-server -t shorturl-server .


client:
	docker build -f Dockerfile-client -t shorturl-server .

down:
	docker-compose down

test:
	go test .


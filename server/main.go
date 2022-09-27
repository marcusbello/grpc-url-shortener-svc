package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc-url-shortener-svc/proto"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50052, "The server port")
)

type server struct {
	pb.UnimplementedShortenerSvcServer
}

func (s *server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	url := in.Url
	log.Printf("Shortener.Add >> RECEIVED: %s", url)
	response := &pb.AddResponse{ShortUrl: url}
	return response, nil
}

func (s *server) Show(ctx context.Context, in *pb.ShowRequest) (*pb.ShowResponse, error) {

	return &pb.ShowResponse{}, nil
}

func (s *server) GetUrl(ctx context.Context, in *pb.GetUrlRequest) (*pb.GetUrlResponse, error) {

	return &pb.GetUrlResponse{}, nil
}

func (s *server) List(ctx context.Context, in *pb.ListRequest) (*pb.ListResponse, error) {

	return &pb.ListResponse{}, nil
}

func (s *server) Update(ctx context.Context, in *pb.UpdateRequest) (*pb.UpdateResponse, error) {

	return &pb.UpdateResponse{}, nil
}

func (s *server) Remove(ctx context.Context, in *pb.RemoveRequest) (*pb.RemoveResponse, error) {

	return &pb.RemoveResponse{}, nil
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	s := grpc.NewServer(opts...)
	pb.RegisterShortenerSvcServer(s, &server{})
	log.Printf("server listening on %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}

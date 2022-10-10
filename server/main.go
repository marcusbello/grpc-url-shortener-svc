package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"grpc-url-shortener-svc/db"
	"grpc-url-shortener-svc/model"
	pb "grpc-url-shortener-svc/proto"
	"log"
	"net"
	"time"
)

var (
	port        = flag.Int("port", 50052, "The server port")
	dbURL       = "postgres://postgres:postgres@db:5432/postgres?sslmode=disable"
	goerBaseUrl = "https://goer.ng"
)

type server struct {
	pb.UnimplementedShortenerSvcServer
	db *db.PGData
}

func (s *server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {

	// Random user for now
	createdBy := "Marcus"

	//Get ShortCode from shortid package
	shortCode, err := s.db.GenShortCode()
	if err != nil {
		return nil, err
	}

	//Create store for AddRequest
	data := &model.AddRequest{
		Url:       in.Url,
		ShortCode: shortCode,
		GoerUrl:   fmt.Sprintf("%s%s", goerBaseUrl, shortCode),
		CreatedBy: createdBy,
		CreatedAt: time.Now(),
	}
	log.Printf("Shortener.Add >> RECEIVED: %s", data)
	err = s.db.Add(ctx, data)
	if err != nil {
		return nil, err
	}
	response := &pb.AddResponse{ShortUrl: data.ShortCode}
	return response, nil
}

func (s *server) Show(ctx context.Context, in *pb.ShowRequest) (*pb.ShowResponse, error) {
	res := &pb.ShowResponse{}
	return res, nil
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

	pgData, err := db.NewPGRepo(dbURL)
	if err != nil {
		log.Fatalf("Failed to connect DB : %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	s := grpc.NewServer(opts...)
	pb.RegisterShortenerSvcServer(s, &server{db: pgData})
	log.Printf("server listening on %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}

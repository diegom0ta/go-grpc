package main

import (
	"bookshop/server/bookshop/pb"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedInventoryServer
}

func getSampleBooks() []*pb.Book {
	b := &pb.GetBookListResponse{}
	return b.GetBooks()
}

func (s *server) GetBookList(ctx context.Context, in *pb.GetBookListRequest) (*pb.GetBookListResponse, error) {
	return &pb.GetBookListResponse{
		Books: getSampleBooks(),
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterInventoryServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

package main

import (
	"bookshop/server/api"
	"bookshop/server/bookshop/pb"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedInventoryServer
}

func getSampleBooks() []*pb.Book {
	books := api.BooksHandler()
	gb := api.GetBytes(books)
	b := &pb.GetBookListResponse{}
	for _, book := range gb {
		b, ok := book.([]*pb.Book)
		if !ok {
			log.Fatal(b)
		}
	}
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
	reflection.Register(s)
	pb.RegisterInventoryServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

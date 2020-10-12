package main

import (
	"context"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	f "github.com/amine7777/go-grpc-api/functions"
	"github.com/amine7777/go-grpc-api/proto"
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp", ":4242")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}
func (s *server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()
	result := f.AddFunc(a, b)
	return &proto.Response{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()
	result := f.MultFunc(a, b)
	return &proto.Response{Result: result}, nil
}

func (s *server) PrimeNumbers(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	number := request.GetNumber()
	prime := f.PrimeFunc(number)
	return &proto.Response{Prime: prime}, nil
}

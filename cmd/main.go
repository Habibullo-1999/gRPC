package main

import (
	"log"
	"net"

	"github.com/Habibullo-1999/gRPC/pkg/api"
	"github.com/Habibullo-1999/gRPC/pkg/api/adder"
	"google.golang.org/grpc"

)

func main() {
	s := grpc.NewServer()

	srv := &adder.GRPCServer{}

	api.RegisterAdderServer(s, srv)

	l, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}

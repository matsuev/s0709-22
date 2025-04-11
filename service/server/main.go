package main

import (
	"log"
	"net"
	"s0709-22/internal/demoapi"
	"s0709-22/internal/testapi"

	"github.com/matsuev/demoapi/pkg/example"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp4", "127.0.0.1:10000")
	if err != nil {
		log.Fatalln(err)
	}

	srv := grpc.NewServer()

	svc := NewDemoService()

	example.Hello()

	demoapi.RegisterDemoServiceServer(srv, svc)
	testapi.RegisterTestServiceServer(srv, testapi.UnimplementedTestServiceServer{})

	if err := srv.Serve(listener); err != nil {
		log.Fatalln()
	}
}

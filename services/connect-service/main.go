package main

import (
	"log"
	"net"
	"s0709-22/internal/proxyproto"
	"s0709-22/services/connect-service/internal/service"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp4", "127.0.0.1:10000")
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := listener.Close(); err != nil {
			log.Println(err)
		}
	}()

	srv := grpc.NewServer()

	svc := service.New()

	proxyproto.RegisterCentrifugoProxyServer(srv, svc)

	if err := srv.Serve(listener); err != nil {
		log.Println()
	}
}

package main

import (
	"context"
	"log"
	"s0709-22/internal/demoapi"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient("127.0.0.1:10000", opts...)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println(err)
		}
	}()

	client := demoapi.NewDemoServiceClient(conn)

	response, err := client.Echo(context.Background(), &demoapi.EchoRequest{
		Message: "gashjgfjagsdfjhg",
	})
	if err != nil {
		log.Println(err)
	} else {
		log.Println(response)
	}
}

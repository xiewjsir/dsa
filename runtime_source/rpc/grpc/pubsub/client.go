package main

import (
	"context"
	"log"
	
	"google.golang.org/grpc"
	"dsa/runtime_source/rpc/grpc/proto/hello"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	
	client :=hello.NewPubsubServiceClient(conn)
	
	_, err = client.Publish(
		context.Background(), &hello.String{Value: "golang: hello Go"},
	)
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Publish(
		context.Background(), &hello.String{Value: "docker: hello Docker"},
	)
	if err != nil {
		log.Fatal(err)
	}
}

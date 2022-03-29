package main

import (
	"context"
	"fmt"
	"io"
	"log"
	
	"dsa/runtime_source/rpc/grpc/proto/hello"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	
	client := hello.NewPubsubServiceClient(conn)
	stream, err := client.Subscribe(
		context.Background(), &hello.String{Value: "golang:"},
	)
	if err != nil {
		log.Fatal(err)
	}
	
	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		
		fmt.Println(reply.GetValue())
	}
}

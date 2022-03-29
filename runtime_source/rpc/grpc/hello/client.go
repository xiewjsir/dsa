package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"
	
	"dsa/runtime_source/rpc/grpc/proto/hello"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	
	client := hello.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &hello.String{Value: "hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())
	
	
	stream, err := client.Channel(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	
	go func() {
		for {
			if err := stream.Send(&hello.String{Value: "hi"}); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}
	}()
	
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

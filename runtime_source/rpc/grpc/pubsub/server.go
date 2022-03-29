package main

import (
	"context"
	"net"
	"strings"
	"time"
	"dsa/runtime_source/rpc/grpc/proto/hello"
	"github.com/moby/moby/pkg/pubsub"
	"google.golang.org/grpc"
	"log"
)

type PubsubService struct {
	pub *pubsub.Publisher
}

func NewPubsubService() *PubsubService {
	return &PubsubService{
		pub: pubsub.NewPublisher(100*time.Millisecond, 10),
	}
}

func (p *PubsubService) Publish(ctx context.Context, arg *hello.String) (*hello.String, error) {
	p.pub.Publish(arg.GetValue())
	return &hello.String{}, nil
}

func (p *PubsubService) Subscribe(arg *hello.String, stream hello.PubsubService_SubscribeServer) error {
	ch := p.pub.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key,arg.GetValue()) {
				return true
			}
		}
		return false
	})
	
	for v := range ch {
		if err := stream.Send(&hello.String{Value: v.(string)}); err != nil {
			return err
		}
	}
	
	return nil
}


func main() {
	grpcServer := grpc.NewServer()
	hello.RegisterPubsubServiceServer(grpcServer, new(PubsubService))
	
	conn, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(conn)
}
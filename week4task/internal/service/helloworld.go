package service

import (
	"golang.org/x/net/context"
	pb "helloworld/api/helloworld/v1"
)

// server is used to implement HelloWord.GreeterServer.
type HelloWord struct {
	*pb.UnimplementedGreeterServer
}

// NewHelloWord new a greeter service.
func NewHelloWord() *HelloWord {
	return &HelloWord{}
}

func (s *HelloWord) mustEmbedUnimplementedGreeterServer() {
	panic("implement me")
}

// SayHello implements helloworld.GreeterServer
func (s *HelloWord) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

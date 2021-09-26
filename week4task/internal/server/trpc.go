package server

import (
	"google.golang.org/grpc"
	pb "helloworld/api/helloworld/v1"
	"helloworld/internal/service"
)

// server is used to implement helloworld.GreeterServer.
func NewGRPCServer(greeter *service.HelloWord) *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, greeter)
	return s
}

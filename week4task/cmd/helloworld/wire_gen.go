// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"google.golang.org/grpc"
	"helloworld/internal/server"
	"helloworld/internal/service"
)

// Injectors from wire.go:

func initApp() *grpc.Server {
	helloWord := service.NewHelloWord()
	grpcServer := server.NewGRPCServer(helloWord)
	return grpcServer
}

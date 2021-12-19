// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"blog/internal/blog-interface/biz"
	"blog/internal/blog-interface/conf"
	"blog/internal/blog-interface/data"
	"blog/internal/blog-interface/server"
	"blog/internal/blog-interface/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/sdk/trace"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, tracerProvider *trace.TracerProvider, logger log.Logger) (*kratos.App, func(), error) {
	conn := data.NewRedisCli(confData)
	client := data.NewMysqlClient(confData)
	reader := data.NewKafkaConsumer(confData)
	elasticClient := data.NewEsInstance(confData)
	dataData, cleanup, err := data.NewData(conn, client, reader, elasticClient, logger)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	userUseCase := biz.NewUserUseCase(userRepo, logger)
	userService := service.NewUserService(userUseCase, logger)
	httpServer := server.NewHTTPServer(confServer, userService, tracerProvider, logger)
	grpcServer := server.NewGRPCServer(confServer, userService, tracerProvider, logger)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}

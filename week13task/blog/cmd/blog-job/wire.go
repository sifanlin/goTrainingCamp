//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"blog/internal/blog-job/biz"
	"blog/internal/blog-job/conf"
	"blog/internal/blog-job/data"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(conf *conf.Data, log log.Logger) error {
	panic(wire.Build(biz.ProviderSet, data.ProviderSet))
}

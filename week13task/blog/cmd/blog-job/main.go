package main

import (
	"flag"
	"github.com/go-kratos/kratos/v2/log"
	"os"

	"blog/internal/blog-job/biz"
	"blog/internal/blog-job/conf"
	"blog/internal/blog-job/data"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}
func main() {
	cfg := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	logger := log.NewStdLogger(os.Stdout)
	if err := cfg.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := cfg.Scan(&bc); err != nil {
		panic(err)
	}
	conn := data.NewRedisCli(bc.Data)
	reader := data.NewKafkaConsumer(bc.Data)
	elasticClient := data.NewEsInstance(bc.Data)
	dataData, cleanup, err := data.NewData(conn, reader, elasticClient, logger)
	defer func() {
		cleanup()
	}()
	err = biz.ArticleConsumer(dataData, logger)
	if err != nil {
		panic(err)
	}
}

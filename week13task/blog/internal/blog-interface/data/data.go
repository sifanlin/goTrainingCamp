package data

import (
	"context"
	"fmt"

	"blog/internal/blog-interface/conf"
	"blog/internal/blog-interface/data/ent"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/gomodule/redigo/redis"
	"github.com/google/wire"
	"github.com/olivere/elastic"
	"github.com/segmentio/kafka-go"

	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewEsInstance, NewUserRepo, NewRedisCli, NewMysqlClient, NewKafkaConsumer)

// Data .
type Data struct {
	Consumer      *kafka.Reader
	ElasticClient *elastic.Client
	RedisCli      redis.Conn
	MysqlClient   *ent.Client
}

func NewData(redisClient redis.Conn, entCli *ent.Client, consumer *kafka.Reader, client *elastic.Client, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		defer client.Stop()
		defer consumer.Close()
		defer entCli.Close()
		defer redisClient.Close()
		log.NewHelper(logger).Info("closing the data resources")
	}
	d := &Data{
		Consumer:      consumer,
		ElasticClient: client,
		RedisCli:      redisClient,
		MysqlClient:   entCli,
	}
	return d, cleanup, nil
}

func NewMysqlClient(c *conf.Data) *ent.Client {
	client, err := ent.Open(c.Database.GetDriver(), c.Database.GetSource())
	if err != nil {
		panic(err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		panic(err)
	}
	return client
}

func NewKafkaConsumer(conf *conf.Data) *kafka.Reader {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  conf.Kafka.Address,
		GroupID:  "group-a",
		Topic:    "topic",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	return reader
}

// NewEsInstance 创建es实例
func NewEsInstance(conf *conf.Data) *elastic.Client {
	esConfig := conf.Elastic
	url := fmt.Sprintf("%s://%s:%s", esConfig.Scheme, esConfig.IP, esConfig.Port)
	client, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetHealthcheck(esConfig.HealthCheckEnabled),
		elastic.SetBasicAuth(esConfig.BasicAuthUsername, esConfig.BasicAuthPassword),
		//see https://github.com/olivere/elastic/wiki/Connection-Problems
		elastic.SetSniff(esConfig.SnifferEnabled))

	if err != nil {
		panic(err)
	}
	return client
}
func NewRedisCli(conf *conf.Data) redis.Conn {
	client, err := redis.Dial(conf.Redis.Network, conf.Redis.Addr)
	if err != nil {
		panic(err)
	}
	return client
}

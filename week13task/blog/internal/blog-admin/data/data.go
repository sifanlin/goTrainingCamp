package data

import (
	"blog/internal/blog-admin/conf"
	"blog/internal/blog-admin/data/ent"
	"context"
	"github.com/google/wire"

	"github.com/Shopify/sarama"
	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewMysqlClient,
	NewUserRepo,
	NewKafkaProducer,
)

// Data .
type Data struct {
	db *ent.Client
	kp sarama.AsyncProducer
}

// NewData .
func NewData(producer sarama.AsyncProducer, client *ent.Client, logger log.Logger) (*Data, func(), error) {

	d := &Data{
		db: client,
		kp: producer,
	}
	cleanup := func() {
		d.db.Close()
		d.kp.Close()
		log.NewHelper(logger).Info("closing the data resources")
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

func NewKafkaProducer(conf *conf.Data) sarama.AsyncProducer {
	c := sarama.NewConfig()
	p, err := sarama.NewAsyncProducer(conf.Kafka.Address, c)
	if err != nil {
		panic(err)
	}
	return p
}

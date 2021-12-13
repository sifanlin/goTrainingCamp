package data

import (
	"context"
	"github.com/pkg/errors"

	"blog/internal/blog-interface/conf"
	"blog/internal/blog-interface/data/ent"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo)

// Data .
type Data struct {
	db *ent.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	client, err := ent.Open(c.Database.GetDriver(), c.Database.GetSource())
	if err != nil {
		return nil, nil, errors.Wrapf(err, "ent open err :err")
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, nil, errors.Wrapf(err, "client.Schema.Create err")
	}
	cleanup := func() {
		defer client.Close()
		log.NewHelper(logger).Info("closing the data resources")
	}
	d := &Data{
		db: client,
	}
	return d, cleanup, nil
}

func newKafkaSender() {

}

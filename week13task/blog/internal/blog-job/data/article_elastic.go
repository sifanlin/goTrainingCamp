package data

import (
	"context"
	"github.com/olivere/elastic"
)

func NewElastic(Client elastic.Client) *elasticClient {
	return &elasticClient{
		client: Client,
	}
}

type elasticClient struct {
	client elastic.Client
}

// CreateIndex 创建索引
func (a *elasticClient) CreateIndex(ctx context.Context, index, mapping string) error {
	exists, err := a.client.IndexExists(index).Do(ctx)
	if err != nil {
		return err
	}
	if exists {
		return nil
	} else {
		_, err = a.client.CreateIndex(index).Body(mapping).Do(ctx)
		return err
	}
}

// Put 单条更新
func (a *elasticClient) Put(ctx context.Context, index, mapping, iType, id string, body interface{}) error {
	err := a.CreateIndex(ctx, index, mapping)
	if err != nil {
		return err
	}
	_, err = a.client.Index().
		Index(index).
		Type(iType).
		Id(id).
		BodyJson(body).
		Do(context.Background())
	return err
}

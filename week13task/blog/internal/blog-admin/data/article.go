package data

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"

	"blog/internal/blog-admin/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type ArticleRepo struct {
	data *Data
	log  *log.Helper
}

func (a *ArticleRepo) CreateArticle(ctx context.Context, article *biz.Article) error {
	ret, err := a.data.db.Article.Create().SetCid(article.CateID).SetContent(article.Content).
		SetTitle(article.Title).Save(ctx)
	if err != nil {
		return err
	}
	ar, err := json.Marshal(article)
	if err != nil {
		return err
	}
	a.data.kp.Input() <- &sarama.ProducerMessage{
		Topic: "article",
		Key:   sarama.ByteEncoder(fmt.Sprint(ret.ID)),
		Value: sarama.ByteEncoder(ar),
	}
	return nil
}

package biz

import (
	pb "blog/api/interface/v1"
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"strconv"

	jobData "blog/internal/blog-job/data"

	"github.com/fatih/structs"
	"github.com/go-kratos/kratos/v2/log"
)

var articleMap = `{
	"mappings":{
		"properties":{
			"id": 			{ "type": "long" },
			"title": 		{ "type": "text" },
			"desc":			{ "type": "text" },
			"content":		{ "type": "text" }
			}
		}
	}`

const (
	articleIndex     = "articleIndex"
	articleIndexType = "doc"
)

const (
	articleRedisPrefix = "articleIndex:%d"
)

type ArticleMessage struct {
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Cid holds the value of the "cid" field.
	Cid int32 `json:"cid,omitempty"`
	// Desc holds the value of the "desc" field.
	Desc string `json:"desc,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// Img holds the value of the "img" field.
	Img string `json:"img,omitempty"`
}

type ArticleUseCase struct {
	repo Article
	log  *log.Helper
}

type Article interface {
	CreateIndex(ctx context.Context, index, mapping string) error
	Put(ctx context.Context, index, mapping, iType, id string, body interface{}) error
}

func ArticleConsumer(data *jobData.Data, logger log.Logger) error {
	go func() {
		for {
			m, err := data.Consumer.FetchMessage(context.Background())
			if err != nil {
				break
			}
			var article ArticleMessage
			err = json.Unmarshal(m.Value, &article)
			if err != nil {
				continue
			}
			body := structs.Map(article)
			articleCli := jobData.NewElastic(data.ElasticClient)
			// 写es
			ID := strconv.Itoa(article.ID)
			err = articleCli.Put(context.Background(), articleIndex, articleMap, articleIndexType, ID, body)
			if err != nil {
				log.NewHelper(logger).Errorf(" articleCli.Put err,err info %v", err)
			}
			// 写redis
			articleKey := fmt.Sprintf(articleRedisPrefix, article.ID)
			mp := jsonpb.Marshaler{}
			pbArticle := pb.ArticleDetailReply{
				Id:      int64(article.ID),
				Title:   article.Title,
				Content: article.Content,
				CateID:  article.Cid,
			}
			articleStr, err := mp.MarshalToString(&pbArticle)
			if err != nil {
				log.NewHelper(logger).Errorf(" RedisCli.Send err,err info %v", err)
			}
			data.RedisCli.Send("set", articleKey, articleStr)
			data.RedisCli.Flush()
			_, err = data.RedisCli.Receive()
			if err != nil {
				log.NewHelper(logger).Errorf(" articleCli.Put err,err info %v", err)
			}
		}
	}()
	return nil
}

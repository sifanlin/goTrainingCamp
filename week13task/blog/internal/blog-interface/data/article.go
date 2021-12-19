package data

import (
	"bytes"
	"context"
	"fmt"

	pb "blog/api/interface/v1"
	"blog/internal/blog-interface/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang/protobuf/jsonpb"
	"github.com/pkg/errors"
)

const (
	articleGetCommand      = "get"
	articleRedisPrefix     = "article:%d"
	articleAttrRedisPrefix = "articleAttr:%d"
)

type ArticleRepo struct {
	data *Data
	log  *log.Helper
}

func (a *ArticleRepo) ArticleList(ctx context.Context, pageSize, pageNumber int32) (list *biz.ArticleList, err error) {

	return &biz.ArticleList{}, nil
}

func (a *ArticleRepo) ArticleDetail(ctx context.Context, ID int64) (*biz.Article, error) {
	articleKey := fmt.Sprintf(articleRedisPrefix, ID)
	articleAttrKey := fmt.Sprintf(articleAttrRedisPrefix, ID)
	// 从redis中取出文章和点赞数 pipeline
	a.data.RedisCli.Send(articleGetCommand, articleKey)
	a.data.RedisCli.Send(articleGetCommand, articleAttrKey)
	a.data.RedisCli.Flush()
	article, err := a.data.RedisCli.Receive()
	if err != nil {
		return nil, err
	}
	articleAttr, err := a.data.RedisCli.Receive()
	if err != nil {
		return nil, err
	}
	articleString, ok := article.(string)
	if ok {
		return nil, errors.Errorf("type err %v", article)
	}
	reader := bytes.NewReader([]byte(articleString))
	pbArticle := pb.ArticleDetailReply{}
	err = jsonpb.Unmarshal(reader, &pbArticle)
	if err != nil {
		return nil, err
	}

	articleAttrString, ok := articleAttr.(string)
	if ok {
		return nil, errors.Errorf("type err %v", article)
	}
	articleAttReader := bytes.NewReader([]byte(articleAttrString))
	articleAttPBArticle := pb.ArticleDetailReply{}
	err = jsonpb.Unmarshal(articleAttReader, &articleAttPBArticle)
	if err != nil {
		return nil, err
	}
	return &biz.Article{
		ID:       pbArticle.Id,
		Title:    pbArticle.Title,
		Content:  pbArticle.Content,
		CateID:   pbArticle.CateID,
		ClickNum: articleAttPBArticle.ClickNum,
	}, nil
}

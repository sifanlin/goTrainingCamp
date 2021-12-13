package data

import (
	"context"

	"blog/internal/blog-interface/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type ArticleRepo struct {
	data *Data
	log  *log.Helper
}

func (a *ArticleRepo) ArticleList(ctx context.Context, pageSize, pageNumber int32) (list *biz.ArticleList, err error) {
	return &biz.ArticleList{}, nil
}

func (a *ArticleRepo) ArticleDetail(context.Context, int64) (*biz.Article, error) {
	return &biz.Article{}, nil
}

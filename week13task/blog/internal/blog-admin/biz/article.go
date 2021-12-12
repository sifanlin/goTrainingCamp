package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Article struct {
	ID      int64
	Title   string
	Content string
	CateID  int32
}

type ArticleList struct {
	list  []*Article
	total int32
}

func NewArticleUseCase(repo ArticleRepo, logger log.Logger) *ArticleUseCase {
	return &ArticleUseCase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

type ArticleUseCase struct {
	repo ArticleRepo
	log  *log.Helper
}

type ArticleRepo interface {
	// CreateArticle 在es中获取
	CreateArticle(ctx context.Context, article *Article) error
}

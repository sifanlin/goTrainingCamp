package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Article struct {
	ID       int64
	Title    string
	Content  string
	CateID   int32
	ClickNum int64
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
	// ArticleList 在es中获取
	ArticleList(ctx context.Context, pageSize, pageNumber int32) (list *ArticleList, err error)
	// ArticleDetail 在 redis中拿到
	ArticleDetail(context.Context, int64) (*Article, error)
}

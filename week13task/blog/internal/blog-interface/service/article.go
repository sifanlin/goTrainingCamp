package service

import (
	"context"

	pb "blog/api/v1"
)

type ArticleService struct {
	pb.UnimplementedArticleServer
}

func NewArticleService() *ArticleService {
	return &ArticleService{}
}

func (s *ArticleService) ArticleList(ctx context.Context, req *pb.ArticleListRequest) (*pb.ArticleListReply, error) {
	return &pb.ArticleListReply{}, nil
}
func (s *ArticleService) ArticleDetail(ctx context.Context, req *pb.ArticleDetailRequest) (*pb.ArticleDetailReply, error) {
	return &pb.ArticleDetailReply{}, nil
}

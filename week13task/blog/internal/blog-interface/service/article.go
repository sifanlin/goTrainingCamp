package service

import (
	"context"

	pb "blog/api/interface/v1"

	"go.opentelemetry.io/otel"
)

type ArticleService struct {
	pb.UnimplementedArticleServer
}

func NewArticleService() *ArticleService {
	return &ArticleService{}
}

func (s *ArticleService) ArticleList(ctx context.Context, req *pb.ArticleListRequest) (*pb.ArticleListReply, error) {
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "ArticleList")
	defer span.End()
	return &pb.ArticleListReply{}, nil
}
func (s *ArticleService) ArticleDetail(ctx context.Context, req *pb.ArticleDetailRequest) (*pb.ArticleDetailReply, error) {
	return &pb.ArticleDetailReply{}, nil
}

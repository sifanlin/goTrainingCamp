package service

import (
	"context"

	pb "blog/api/v1"
)

type CategoryService struct {
	pb.UnimplementedCategoryServer
}

func NewCategoryService() *CategoryService {
	return &CategoryService{}
}

func (s *CategoryService) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.CreateCategoryReply, error) {
	return &pb.CreateCategoryReply{}, nil
}

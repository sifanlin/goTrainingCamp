package service

import (
	"context"

	"blog/internal/blog-admin/biz"

	pb "blog/api/admin/v1"
	"github.com/go-kratos/kratos/v2/log"
)

type UserService struct {
	pb.UnimplementedUserServer
	user *biz.UserUseCase
	log  *log.Helper
}

func NewUserService(user *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{
		user: user,
		log:  log.NewHelper(logger),
	}
}
func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	return &pb.CreateUserReply{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.CreateUserReply, error) {
	return &pb.CreateUserReply{}, nil
}

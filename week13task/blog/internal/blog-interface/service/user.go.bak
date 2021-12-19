package service

import (
	"blog/internal/blog-interface/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"

	pb "blog/api/v1"
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
	s.log.Infof("req 333 %+v", req)
	err := s.user.Create(ctx, &biz.User{
		UserName: req.Username,
		Password: req.Password, //todo 加密
		Role:     1,
	})
	if err != nil {
		s.log.Errorf("create user err %+v", err)
		return &pb.CreateUserReply{Message: err.Error()}, nil
	}
	return &pb.CreateUserReply{}, nil
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.CreateUserReply, error) {

	return &pb.CreateUserReply{}, nil
}

func (s *UserService) TestUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	return &pb.CreateUserReply{}, nil
}

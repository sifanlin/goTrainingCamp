package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	ID       int64
	UserName string
	Password string
	Role     int
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo: repo,
	}
}

type UserRepo interface {
	CreateUser(context.Context, *User) error
	GetUser(context.Context, *User) (*User, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func (uc *UserUseCase) Create(ctx context.Context, user *User) error {
	return uc.repo.CreateUser(ctx, user)
}

func (uc *UserUseCase) GetUserDetail(ctx context.Context, user *User) (*User, error) {
	return uc.repo.GetUser(ctx, user)
}

package data

import (
	"context"
	"time"

	"blog/internal/blog-interface/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u *userRepo) CreateUser(ctx context.Context, user *biz.User) error {
	_, err := u.data.db.User.Create().
		SetCreateTime(time.Now()).SetUserName(user.UserName).
		SetPassword(user.Password).SetRole(user.Role).Save(ctx)
	return err
}

func (u *userRepo) GetUser(ctx context.Context, bizUser *biz.User) (*biz.User, error) {
	user, err := u.data.db.User.Get(ctx, int(bizUser.ID))
	if err != nil {
		return nil, errors.Wrapf(err, "get user err req userID:%v", bizUser.ID)
	}
	return &biz.User{
		ID:       int64(user.ID),
		UserName: user.UserName,
	}, err
}

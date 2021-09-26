package data

import (
	"context"
	"helloworld/internal/biz"
)

type greeterRepo struct {
}

// NewGreeterRepo .
func NewGreeterRepo() biz.GreeterRepo {
	return &greeterRepo{}
}

func (r *greeterRepo) CreateGreeter(ctx context.Context, g *biz.Greeter) error {
	return nil
}

func (r *greeterRepo) UpdateGreeter(ctx context.Context, g *biz.Greeter) error {
	return nil
}

package data

import (
	"context"
	"strike_layout_template/internal/biz"
)

type greeterRepo struct {
	data *Data
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
	}
}

// FindByID implements biz.GreeterRepo.
func (g *greeterRepo) FindByID(context.Context, int64) (*biz.Greeter, error) {
	panic("unimplemented")
}

// ListAll implements biz.GreeterRepo.
func (g *greeterRepo) ListAll(context.Context) ([]*biz.Greeter, error) {
	panic("unimplemented")
}

// ListByHello implements biz.GreeterRepo.
func (g *greeterRepo) ListByHello(context.Context, string) ([]*biz.Greeter, error) {
	panic("unimplemented")
}

// Save implements biz.GreeterRepo.
func (g *greeterRepo) Save(context.Context, *biz.Greeter) (*biz.Greeter, error) {
	panic("unimplemented")
}

// Update implements biz.GreeterRepo.
func (g *greeterRepo) Update(context.Context, *biz.Greeter) (*biz.Greeter, error) {
	panic("unimplemented")
}

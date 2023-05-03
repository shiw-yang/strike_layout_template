//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"github.com/google/wire"

	"strike_layout_template/internal/biz"
	"strike_layout_template/internal/conf"
	"strike_layout_template/internal/data"
	"strike_layout_template/internal/server"
	"strike_layout_template/internal/service"
)

func wireApp(*conf.Server, *conf.Data) (*server.Server, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}

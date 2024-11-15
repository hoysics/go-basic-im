//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/hoysics/basic-im/app/state/internal/biz"
	"github.com/hoysics/basic-im/app/state/internal/conf"
	"github.com/hoysics/basic-im/app/state/internal/data"
	"github.com/hoysics/basic-im/app/state/internal/server"
	"github.com/hoysics/basic-im/app/state/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Registry, *conf.State, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}

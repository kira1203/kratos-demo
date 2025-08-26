//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos-layout/internal/bondProductApi/biz"
	"github.com/go-kratos/kratos-layout/internal/bondProductApi/data"
	"github.com/go-kratos/kratos-layout/internal/bondProductApi/server"
	"github.com/go-kratos/kratos-layout/internal/bondProductApi/service"
	"github.com/go-kratos/kratos-layout/internal/config"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*config.Server, *config.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		data.GormDataProviderSet,
		data.ProviderSet,
		biz.BizProviderSet,
		service.ServiceProviderSet,
		server.ServerProviderSet,
		newApp,
	),
	)
}

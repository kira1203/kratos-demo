//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos-layout/internal/config"
	"github.com/go-kratos/kratos-layout/internal/gateway"
	"github.com/go-kratos/kratos-layout/internal/gateway/registry"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*config.Server, *config.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		gateway.ServerProviderSet,
		registry.NacosRegistryProviderSet,
		newApp,
	),
	)
}

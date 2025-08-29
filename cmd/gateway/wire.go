//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"kratos-demo/internal/config"
	"kratos-demo/internal/gateway"
	"kratos-demo/internal/gateway/registry"
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

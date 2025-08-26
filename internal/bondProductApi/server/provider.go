package server

import (
	"github.com/go-kratos/kratos-layout/internal/bondProductApi/service"
	"github.com/google/wire"
)

// ServerProviderSet is server providers.
var ServerProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer, NewRegisterSet)

func NewRegisterSet(
	product *service.ProductService,
) []Register {
	return []Register{product}
}

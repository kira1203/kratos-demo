package server

import (
	"github.com/google/wire"
	"kratos-demo/internal/bondProductApi/service"
)

// ServerProviderSet is server providers.
var ServerProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer, NewRegisterSet)

func NewRegisterSet(
	product *service.ProductService,
) []Register {
	return []Register{product}
}

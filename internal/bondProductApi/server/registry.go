package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type Register interface {
	HttpRegister(*http.Server)
	GrpcRegister(*grpc.Server)
}

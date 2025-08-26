package gateway

import (
	"github.com/google/wire"
)

var ServerProviderSet = wire.NewSet(NewHTTPServer, NewHandler)

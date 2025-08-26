package registry

import "github.com/google/wire"

var NacosRegistryProviderSet = wire.NewSet(NewNacosClientRegistry)

package registry

import (
	"context"
	"github.com/go-kratos/kratos/v2/registry"
)

const (
	namespace = "saas"
)

// RegistryClient 定义注册中心客户端接口
type RegistryClient interface {
	GetService(ctx context.Context, serviceName string) ([]*registry.ServiceInstance, error)
	// 可根据需要添加其他方法，如 Register、Deregister 等
}

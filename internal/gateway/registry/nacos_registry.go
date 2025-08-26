package registry

import (
	"context"
	"github.com/go-kratos/kratos-layout/internal/config"
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/registry"
)

// NacosClient 实现 RegistryClient 接口
type NacosClient struct {
	discovery registry.Discovery
}

// GetService 从 Nacos 获取服务实例
func (c *NacosClient) GetService(ctx context.Context, serviceName string) ([]*registry.ServiceInstance, error) {
	return c.discovery.GetService(ctx, serviceName)
}

func NewNacosClientRegistry() (RegistryClient, error) {
	client, err := config.GetNacosClient().GetNamingClient(namespace)
	if err != nil {
		return nil, err
	}

	discovery := nacos.New(client)

	return &NacosClient{discovery: discovery}, nil
}

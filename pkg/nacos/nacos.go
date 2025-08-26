package nacos

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"sync"
)

type NacosCli struct {
	configCliMap map[string]config_client.IConfigClient // 配置客户端（用于配置中心）
	namingCliMap map[string]naming_client.INamingClient // 命名客户端（用于服务发现）
	cacheLogDir  string
	logDir       string
	LogLvl       string
	lock         sync.Mutex
	addr         string
	port         int32
	accessKey    string
	secretKey    string
	region       string
}

type NacosCliOption func(*NacosCli)

func NewNacosOpts(opts ...NacosCliOption) *NacosCli {
	c := &NacosCli{
		configCliMap: make(map[string]config_client.IConfigClient),
		namingCliMap: make(map[string]naming_client.INamingClient),
		LogLvl:       "warn",
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func WithCacheDir(cacheLogDir string) NacosCliOption {
	return func(config *NacosCli) {
		config.cacheLogDir = cacheLogDir
	}
}

func WithLogDir(dir string) NacosCliOption {
	return func(config *NacosCli) {
		config.logDir = dir
	}
}
func WithLogLevel(lvl string) NacosCliOption {
	return func(config *NacosCli) {
		config.LogLvl = lvl
	}
}

func WithAddr(addr string) NacosCliOption {
	return func(config *NacosCli) {
		config.addr = addr
	}
}

func WithPort(port int32) NacosCliOption {
	return func(config *NacosCli) {
		config.port = port
	}
}

func WithRegionId(region string) NacosCliOption {
	return func(config *NacosCli) {
		config.region = region
	}
}

func WithKmsAK(ak string) NacosCliOption {
	return func(config *NacosCli) {
		config.accessKey = ak
	}
}

func WithKmsSK(sk string) NacosCliOption {
	return func(config *NacosCli) {
		config.secretKey = sk
	}
}

func (c *NacosCli) getConfigCli(ns string) (config_client.IConfigClient, error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if cli, ok := c.configCliMap[ns]; ok {
		return cli, nil
	}
	cliCfg := *constant.NewClientConfig(
		constant.WithTimeoutMs(5000),
		constant.WithNamespaceId(ns),
		constant.WithOpenKMS(true),
		constant.WithRegionId(c.region),
		constant.WithSecretKey(c.secretKey),
		constant.WithAccessKey(c.accessKey),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir(c.logDir),
		constant.WithCacheDir(c.cacheLogDir),
		constant.WithLogLevel(c.LogLvl),
	)
	sc := []constant.ServerConfig{
		{
			IpAddr: c.addr,
			Port:   uint64(c.port),
		},
	}
	nacosCli, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cliCfg,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		fmt.Println("init nacos client error:", err)
		return nil, err
	}
	c.configCliMap[ns] = nacosCli
	return nacosCli, nil
}

func (c *NacosCli) GetCfgFromNacos(id, group, ns string) (string, error) {
	cli, err := c.getConfigCli(ns)
	if err != nil {
		return "", err
	}
	return cli.GetConfig(vo.ConfigParam{
		Group:  group,
		DataId: id,
	})
}

func (c *NacosCli) GetNamingClient(ns string) (naming_client.INamingClient, error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if cli, ok := c.namingCliMap[ns]; ok {
		return cli, nil
	}

	cliCfg := *c.getClientConfig() // 复用配置构造
	sc := c.getServerConfigs()

	namingCli, err := clients.NewNamingClient(vo.NacosClientParam{
		ClientConfig:  &cliCfg,
		ServerConfigs: sc,
	})
	if err != nil {
		return nil, err
	}

	c.namingCliMap[ns] = namingCli
	return namingCli, nil
}

func (c *NacosCli) getClientConfig() *constant.ClientConfig {
	return constant.NewClientConfig(
		constant.WithTimeoutMs(5000),
		constant.WithNamespaceId(""), // 注意：这里先不设 Namespace，后面再填
		constant.WithOpenKMS(true),
		constant.WithRegionId(c.region),
		constant.WithSecretKey(c.secretKey),
		constant.WithAccessKey(c.accessKey),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir(c.logDir),
		constant.WithCacheDir(c.cacheLogDir),
		constant.WithLogLevel(c.LogLvl),
	)
}

func (c *NacosCli) getServerConfigs() []constant.ServerConfig {
	return []constant.ServerConfig{
		{
			IpAddr: c.addr,
			Port:   uint64(c.port),
		},
	}
}

func (c *NacosCli) CloseNacosConns() {
	c.lock.Lock()
	defer c.lock.Unlock()

	//// 关闭所有配置客户端
	//for ns, configClient := range c.configCliMap {
	//	configClient.CloseClient()
	//	delete(c.configCliMap, ns)
	//}
	//
	//// 关闭所有命名客户端
	//for ns, namingClient := range c.namingCliMap {
	//	namingClient.CloseClient()
	//	delete(c.namingCliMap, ns)
	//}
}

// 服务注册
func (c *NacosCli) RegisterServiceInstance(serviceName, groupName, ip string, port uint64, ns string, healthy bool) error {
	cli, err := c.GetNamingClient(ns)
	if err != nil {
		return err
	}

	success, err := cli.RegisterInstance(vo.RegisterInstanceParam{
		ServiceName: serviceName,
		GroupName:   groupName,
		Ip:          ip,
		Port:        port,
		Weight:      1.0,
		Enable:      true,
		Healthy:     healthy,
		Ephemeral:   true,
		Metadata:    map[string]string{},
	})
	if !success || err != nil {
		return fmt.Errorf("failed to register instance %s:%d to service '%s': %w", ip, port, serviceName, err)
	}

	return nil
}

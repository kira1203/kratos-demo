package gateway

import (
	"context"
	"kratos-demo/internal/gateway/registry"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

// Handler 定义网关的处理逻辑
type Handler struct {
	proxy *httputil.ReverseProxy
}

func NewHandler(client registry.RegistryClient) (*Handler, error) {
	// 创建 ReverseProxy（此处暂时占位，动态选择实例）
	proxy := httputil.NewSingleHostReverseProxy(&url.URL{})

	// 动态选择服务实例
	proxy.Director = func(req *http.Request) {
		// 根据请求路径动态选择服务名
		serviceName := determineServiceName(req.URL.Path)
		if serviceName == "" {
			log.Printf("No service found for path: %s", req.URL.Path)
			return
		}

		// 从 Nacos 获取服务实例
		instances, err := client.GetService(context.Background(), serviceName)
		if err != nil || len(instances) == 0 {
			log.Printf("Failed to discover service: %v", err)
			return
		}
		// 选择第一个可用实例（可扩展为负载均衡逻辑）
		instance := instances[0]
		if len(instance.Endpoints) == 0 {
			log.Printf("No endpoints available for instance: %v", instance.ID)
			return
		}
		endpoint := instance.Endpoints[0]
		targetURL, err := url.Parse(endpoint)
		if err != nil {
			log.Printf("Failed to parse endpoint %s: %v", endpoint, err)
			return
		}
		req.URL.Scheme = targetURL.Scheme
		req.URL.Host = targetURL.Host
		req.URL.Path = singleJoiningSlash("", req.URL.Path)
		req.Host = targetURL.Host
	}

	return &Handler{
		proxy: proxy,
	}, nil
}

// singleJoiningSlash 辅助函数，确保路径正确拼接
func singleJoiningSlash(a, b string) string {
	if a == "" {
		return b
	}
	return a + "/" + b
}

// determineServiceName 根据请求路径确定服务名
func determineServiceName(path string) string {
	// 自定义路由规则，根据路径前缀映射服务名
	switch {
	case strings.HasPrefix(path, "/api/v1/product"):
		return "flp-note-product-api"
	case strings.HasPrefix(path, "/api/v1/user"):
		return "flp-note-user-api"
	// 添加更多路由规则...
	default:
		return "" // 未匹配到服务
	}
}

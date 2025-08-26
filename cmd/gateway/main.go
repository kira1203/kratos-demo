package main

import (
	"fmt"
	"github.com/go-kratos/kratos-layout/internal/config"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	"os"
)

const (
	AppName           = "gateway"
	AppNacosConfigKey = "gateway@@saas@@saas"
	AppNamespace      = "saas"
	NacosRedisRWKey   = ""
)

var configKeys = map[string][]string{
	AppNacosConfigKey: {"", "YAML"},
}

var id, _ = os.Hostname()

func init() {
	config.LoadEnvFile()
	defer config.CloseNacosConns()
	err := config.NewNacosClientInsFromEnv(AppName, "/data/logs")
	if err != nil {
		panic(err)
	}

	for k, v := range configKeys {
		if err := config.GetViperCfgFromNacos(k, v[0], v[1]); err != nil {
			fmt.Printf("GetViperCfgFromNacos key:%s error:%s\n", v, err.Error())
			panic(err)
		}
	}
}

func newApp(logger log.Logger, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(AppName),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
		),
	)
}

func main() {
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", AppName,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)

	var bc config.Bootstrap
	if err := config.GetCfg("", &bc); err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// 启动网关
	if err := app.Start(nil); err != nil {
		panic(err)
	}
}

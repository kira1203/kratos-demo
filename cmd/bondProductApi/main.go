package main

import (
	"fmt"
	"github.com/go-kratos/kratos-layout/internal/config"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"

const (
	AppName           = "saas-bond-product-api"
	AppNacosConfigKey = "saas-bond-product-api@@saas@@saas"
	NacosRedisRWKey   = ""
	NacosDBRWKey      = "cipher-mysql-lupu_bond-rw@@common@@lupu"
)

var configKeys = map[string][]string{
	AppNacosConfigKey: {"", "YAML"},
	NacosDBRWKey:      {"mysql-bond-db-rw", "json"},
}

var id, _ = os.Hostname()

func Init() {
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

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(AppName),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
	)
}

func main() {

	a := make([]int, 3, 3) // len=3, cap=5 → append 不会扩容
	a[0], a[1], a[2] = 1, 2, 3

	fmt.Printf("a before: %v (len=%d, cap=%d)\n", a, len(a), cap(a))

	func(s []int) {
		fmt.Printf("s before append: %v (len=%d, cap=%d)\n", s, len(s), cap(s))
		s = append(s, 4) // cap=5 >= 4，不会扩容
		fmt.Printf("s after append: %v (len=%d, cap=%d)\n", s, len(s), cap(s))
		s[0] = 999 // 修改底层数组第0个元素
	}(a)

	fmt.Printf("a after: %v\n", a)

	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", AppName,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	//c := config.New(
	//	config.WithSource(
	//		file.NewSource(flagconf),
	//	),
	//)
	//defer c.Close()
	//
	//if err := c.Load(); err != nil {
	//	panic(err)
	//}

	var bc config.Bootstrap
	//if err := c.Scan(&bc); err != nil {
	//	panic(err)
	//}
	if err := config.GetCfg("", &bc); err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}

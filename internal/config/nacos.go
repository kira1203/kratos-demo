package config

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	pkgnacos "github.com/go-kratos/kratos-layout/pkg/nacos"
	"github.com/spf13/viper"
)

var (
	ins *pkgnacos.NacosCli
)

const (
	KMS_ACCESS_KEY  = "KMS_ACCESS_KEY"
	kMS_SECRET_KEY  = "KMS_SECRET_KEY"
	NACOS_URL       = "NACOS_URL"
	NACOS_PORT      = "NACOS_PORT"
	NACOS_REGION_ID = "NACOS_REGION_ID"
	NACOS_CACHE_DIR = "NACOS_CACHE_DIR"
)

func GetNacosClient() *pkgnacos.NacosCli {
	return ins
}

func GetViperCfgFromNacos(nacosKey, localKey, cfgType string) error {
	b, err := GetConfigFromNacos(nacosKey)
	if err != nil {
		return err
	}
	v := viper.New()
	v.SetConfigType(cfgType)
	err = v.ReadConfig(bytes.NewBufferString(b))
	if err != nil {
		return err
	}
	cfg := v.AllSettings()
	if len(cfg) == 0 {
		return fmt.Errorf("no config for key %s", nacosKey)
	}
	if localKey == "" {
		return viper.MergeConfigMap(cfg)
	}
	viper.Set(localKey, cfg)
	return nil
}

func NewNacosClientInsFromEnv(app, dir string) error {
	url := os.Getenv(NACOS_URL)
	if url == "" {
		return errors.New("cant get env var NACOS_URL")
	}
	port, er := strconv.Atoi(os.Getenv(NACOS_PORT))
	if er != nil {
		port = 8848
	}

	if dir == "" {
		dir = "/data/logs"
	}

	ins = pkgnacos.NewNacosOpts(
		pkgnacos.WithAddr(os.Getenv(NACOS_URL)),
		pkgnacos.WithPort(int32(port)),
		pkgnacos.WithCacheDir(dir+"/"+app),
		pkgnacos.WithLogLevel("warn"),
		pkgnacos.WithLogDir(dir+"/"+app),
		pkgnacos.WithKmsAK(os.Getenv(KMS_ACCESS_KEY)),
		pkgnacos.WithKmsSK(os.Getenv(kMS_SECRET_KEY)),
		pkgnacos.WithRegionId(os.Getenv(NACOS_REGION_ID)),
	)
	return nil
}

func GetCfgByNacosKey(nacosKey, key, cfgType string, cfg interface{}) error {
	b, err := GetConfigFromNacos(nacosKey)
	if err != nil {
		return err
	}
	v := viper.New()
	v.SetConfigType(cfgType)
	err = v.ReadConfig(bytes.NewBufferString(b))
	if err != nil {
		return err
	}
	if key == "" {
		return viper.Unmarshal(cfg)
	}
	return viper.Sub(key).Unmarshal(cfg)
}

func CloseNacosConns() {
	if ins == nil {
		return
	}
	ins.CloseNacosConns()
}

// get config string  so you can Unmarshal yourself !
func GetConfigFromNacos(key string) (string, error) {
	return GetConfigFromNacosSep(key, "@@")
}

func GetConfigFromNacosSep(key, sep string) (string, error) {
	arr := strings.Split(key, sep)
	return ins.GetCfgFromNacos(arr[0], arr[1], arr[2])
}

func Get(key, sep string) (string, error) {
	arr := strings.Split(key, sep)
	return ins.GetCfgFromNacos(arr[0], arr[1], arr[2])
}

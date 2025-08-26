package config

import (
	"fmt"
	"github.com/joho/godotenv"

	"github.com/spf13/viper"
)

func LoadEnvFile(dir ...string) {
	if len(dir) == 0 {
		err := godotenv.Load()
		if err != nil {
			fmt.Printf("cant read env file:%v\n", err)
		}
		return
	}
	er := godotenv.Load(dir...)
	if er != nil {
		fmt.Printf("cant read env file:%v", er)
	}
}

func GetCfg(key string, cfg interface{}) error {
	if key == "" {
		return viper.Unmarshal(cfg)
	}
	return viper.UnmarshalKey(key, cfg)
}

func GetCfgString(key string) string {
	return viper.GetString(key)
}

func GetCfgStringSlice(key string) []string {
	return viper.GetStringSlice(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}

func ClearAllConfig() {
	viper.Reset()
}

func PrintConfig(vps ...viper.Viper) {
	temp := viper.New()
	if len(vps) == 0 {
		vps = append(vps, *viper.GetViper())
	}
	for _, vp := range vps {
		for _, key := range vp.AllKeys() {
			temp.Set(key, viper.Get(key))
		}
	}
}

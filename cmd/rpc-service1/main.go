package config

import (
	"flag"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Address string
}

func LoadConfig() *Config {
	// 1. 从标志读取
	address := flag.String("address", "", "RPC server address")
	flag.Parse()

	// 2. 从环境变量读取
	if *address == "" {
		*address = os.Getenv("RPC_ADDRESS")
	}

	// 3. 从配置文件读取
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if *address == "" {
		*address = viper.GetString("address")
	}

	return &Config{
		Address: *address,
	}
}

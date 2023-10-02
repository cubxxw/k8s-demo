package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

const (
	DefaultConfigPath = "configs"
	DefaultConfigFile = "config"
)

type Config struct {
	Address string
}

func getAddressFromFlag() string {
	address := flag.String("address", "", "RPC server address")
	flag.Parse()
	return *address
}

func getAddressFromEnv() string {
	return os.Getenv("RPC_ADDRESS")
}

func getAddressFromConfig(configPath string) string {
	log.Println("DefaultConfigFile is", DefaultConfigFile)
	log.Println("configPath is", configPath)

	viper.SetConfigName(DefaultConfigFile)
	viper.AddConfigPath(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
	return viper.GetString("address")
}

func LoadConfig(configPath string) *Config {
	address := getAddressFromFlag()
	if address == "" {
		address = getAddressFromEnv()
	}
	if address == "" {
		address = getAddressFromConfig(configPath)
	}

	log.Println("Address is", address) // 更新了日志消息
	return &Config{
		Address: address,
	}
}

func main() {
	configPath := flag.String("configPath", DefaultConfigPath, "Path to the configuration directory")
	config := LoadConfig(*configPath)
	log.Println("Config is", config)
	fmt.Printf("RPC Server Address: %s\n", config.Address)

	// 无限循环以保持程序运行
	for {
		// 每10秒打印一次RPC地址
		fmt.Printf("RPC Server Address: %s\n", config.Address)
		time.Sleep(10 * time.Second)
	}
}

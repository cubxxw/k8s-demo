package main

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestGetAddressFromFlag(t *testing.T) {
	// 设置命令行参数
	os.Args = []string{"cmd", "-address=127.0.0.1:8080"}
	address := getAddressFromFlag()
	assert.Equal(t, "127.0.0.1:8080", address)
}

func TestGetAddressFromEnv(t *testing.T) {
	// 设置环境变量
	os.Setenv("RPC_ADDRESS", "127.0.0.1:9090")
	address := getAddressFromEnv()
	assert.Equal(t, "127.0.0.1:9090", address)
}

func TestGetAddressFromConfig(t *testing.T) {
	// 使用一个测试配置文件
	viper.SetConfigName("test_config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		t.Fatalf("Error reading config file: %s", err)
	}
	address := getAddressFromConfig(".")
	assert.Equal(t, "127.0.0.1:7070", address)
}

func TestLoadConfig(t *testing.T) {
	// 设置命令行参数
	os.Args = []string{"cmd", "-address=127.0.0.1:8080"}
	config := LoadConfig(".")
	assert.Equal(t, "127.0.0.1:8080", config.Address)

	// 设置环境变量
	os.Setenv("RPC_ADDRESS", "127.0.0.1:9090")
	config = LoadConfig(".")
	assert.Equal(t, "127.0.0.1:9090", config.Address)

	// 使用一个测试配置文件
	viper.SetConfigName("test_config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		t.Fatalf("Error reading config file: %s", err)
	}
	config = LoadConfig(".")
	assert.Equal(t, "127.0.0.1:7070", config.Address)
}

// 注意：为了运行这些测试，您需要一个名为test_config.yaml的测试配置文件在当前目录中，内容如下：
// address: "127.0.0.1:7070"

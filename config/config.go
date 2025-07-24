package config

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	httpkit "github.com/tiamxu/kit/http"
	"github.com/tiamxu/kit/log"
	"gopkg.in/yaml.v3"
)

var (
	cfg        *Config
	name       = "aliops"
	configPath = "config/config.yaml"
)

type Config struct {
	ENV      string                  `yaml:"env"`
	LogLevel string                  `yaml:"log_level"`
	HttpSrv  httpkit.GinServerConfig `yaml:"http_srv"`
	Aliyun   AliyunConfig            `yaml:"aliyun"`
	// dnsClient *client.DNSClient
}

type AliyunConfig struct {
	AccessKeyId     string `yaml:"access_key_id"`
	AccessKeySecret string `yaml:"access_key_secret"`
	RegionId        string `yaml:"region_id"`
}

func (c *Config) Initial() (err error) {

	defer func() {
		if err == nil {
			log.Printf("config initialed, env: %s,name: %s", cfg.ENV, name)
		}
	}()
	//日志
	if level, err := logrus.ParseLevel(c.LogLevel); err != nil {
		return fmt.Errorf("invalid log level: %w", err)
	} else {
		log.DefaultLogger().SetLevel(level)
	}
	// c.dnsClient, err = client.NewDNSClient(&c.Aliyun)
	// if err != nil {
	// 	return fmt.Errorf("初始化DNS客户端失败: %w", err)
	// }
	return nil
}
func LoadConfig() (*Config, error) {
	cfg = new(Config)
	// env := os.Getenv("ENV")
	env := "local"

	switch env {
	case "dev":
		configPath = "config/config-dev.yaml"
	case "test":
		configPath = "config/config-test.yaml"
	case "prod":
		configPath = "config/config-prod.yaml"
	default:
		configPath = "config/config.yaml"
	}
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败:%w", err)
	}
	expanded := os.ExpandEnv(string(data))

	if err := yaml.Unmarshal([]byte(expanded), &cfg); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %w", err)
	}

	// 验证必要配置
	if cfg.Aliyun.AccessKeyId == "" {
		return nil, fmt.Errorf("缺少阿里云AccessKeyId配置")
	}
	if cfg.Aliyun.AccessKeySecret == "" {
		return nil, fmt.Errorf("缺少阿里云AccessKeySecret配置")
	}

	return cfg, nil

}

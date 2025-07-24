package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tiamxu/aliops/api"
	"github.com/tiamxu/aliops/client"
	"github.com/tiamxu/aliops/config"
	"github.com/tiamxu/aliops/service"
	"github.com/tiamxu/kit/log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}
	if err := cfg.Initial(); err != nil {
		log.Fatalf("Config initialization failed: %v", err)
	}
	dnsClient, err := client.NewDNSClient(&cfg.Aliyun)
	if err != nil {
		log.Fatalf("初始化DNS客户端失败: %v", err)

	}
	// 创建服务
	dnsService := service.NewDNSService(dnsClient)
	dnsHandler := api.NewDNSHandler(dnsService)
	r := gin.Default()

	api := r.Group("/api/dns")
	{
		api.GET("/records", dnsHandler.List)
	}
	// routes.InitRoutes(r)
	if err := r.Run(cfg.HttpSrv.Address); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}

}

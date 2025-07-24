package routes

// func InitRoutes(r *gin.Engine) {
// 	dnsClient, err := client.NewDNSClient(&cfg.Aliyun)
// 	if err != nil {
// 		log.Fatalf("初始化DNS客户端失败: %v", err)

// 	}
// 	// 创建服务
// 	dnsService := service.NewDNSService(dnsClient)
// 	dnsHandler := api.NewDNSHandler(dnsService)

// 	api := r.Group("/api/v1/dns")
// 	{
// 		api.GET("/records", dnsHandler.List)
// 	}
// }

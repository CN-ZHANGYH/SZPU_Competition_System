package main

import (
	"Social_Gin/config"
	"Social_Gin/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	config.InitMySQL()

	// 初始化路由组
	r := gin.Default()
	r = config.InitCors(r)
	r = router.CollectVmRoute(r)
	r = router.CollectUserRoute(r)
	r = router.CollectContentRoute(r)

	// 启动HTTP服务，默认在0.0.0.0:8888启动服务
	r.Run(":8888")
}

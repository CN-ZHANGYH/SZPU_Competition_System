package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func InitCors(r *gin.Engine) *gin.Engine {
	// 设置跨域请求的头部信息
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                       // 允许跨域发来请求的网站
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // 允许的请求方法
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool { // 自定义过滤源站的方法
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	return r
}

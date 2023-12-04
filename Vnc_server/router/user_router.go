package router

import (
	"Social_Gin/controller"
	"github.com/gin-gonic/gin"
)

func CollectUserRoute(r *gin.Engine) *gin.Engine {
	userGroup := r.Group("/user")
	{
		userGroup.POST("/login", controller.Login)
		userGroup.POST("/register", controller.Register)
		userGroup.GET("/userInfo", controller.GetUserInfo)
		userGroup.PUT("/update", controller.UpdateUserInfo)
	}
	return r
}

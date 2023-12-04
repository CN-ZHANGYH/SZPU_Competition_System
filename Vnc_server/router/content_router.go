package router

import (
	"Social_Gin/controller"
	"github.com/gin-gonic/gin"
)

func CollectContentRoute(r *gin.Engine) *gin.Engine {
	userGroup := r.Group("/content")
	{
		userGroup.POST("/add", controller.AddContent)
		userGroup.POST("/get", controller.GetContentInfo)
		userGroup.POST("/getList", controller.GetContentList)
		userGroup.GET("/getLabel", controller.GetContentCount)
	}
	return r
}

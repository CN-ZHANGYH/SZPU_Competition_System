package router

import (
	"Social_Gin/controller"
	"github.com/gin-gonic/gin"
)

func CollectVDockerRoute(r *gin.Engine) *gin.Engine {
	userGroup := r.Group("/docker")
	{
		userGroup.GET("/info", controller.GetDockerInfo)
		userGroup.POST("/createVm", controller.CreateVm)
		userGroup.GET("/getContainerList", controller.GetDockerContainerList)
		userGroup.DELETE("/delete", controller.RemoveDockerContainer)
		userGroup.POST("/start", controller.StartDockerContainer)
		userGroup.POST("/stop", controller.StopDockerContainer)
	}
	return r
}

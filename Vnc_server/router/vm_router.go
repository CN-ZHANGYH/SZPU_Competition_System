package router

import (
	"Social_Gin/controller"
	"github.com/gin-gonic/gin"
)

func CollectVmRoute(r *gin.Engine) *gin.Engine {
	userGroup := r.Group("/vm")
	{
		userGroup.POST("/add", controller.AddVm)
		userGroup.GET("/get", controller.GetVm)
		userGroup.GET("/getList", controller.GetVmList)
		userGroup.DELETE("/delete", controller.DeleteVm)
		userGroup.GET("/getHostAndUrl", controller.GetVmHostAndUrl)
		userGroup.POST("/create", controller.CreateVm)
		userGroup.GET("/getContainerList", controller.GetDockerContainerList)
	}
	return r
}

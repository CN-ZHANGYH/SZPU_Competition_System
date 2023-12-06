package controller

import (
	"Social_Gin/docker"
	"Social_Gin/model/response"
	model "Social_Gin/model/vnc"
	"Social_Gin/service"
	"github.com/gin-gonic/gin"
)

func GetDockerContainerList(ctx *gin.Context) {
	service.GetDockerVmList(ctx)
}

func CreateVm(ctx *gin.Context) {
	var dockerVmInfo model.DockerVmInfo
	if err := ctx.ShouldBind(&dockerVmInfo); err == nil {
		service.CreateVmInfo(&dockerVmInfo, ctx)
		return
	}
	response.FailWithMessage("创建失败", ctx)
	return
}

func GetDockerInfo(ctx *gin.Context) {
	docker.GetDockerInfo(ctx)
}

func RemoveDockerContainer(ctx *gin.Context) {
	containerId := ctx.Query("containerId")
	docker.RemoveDockerContainer(containerId, ctx)
}

func StartDockerContainer(ctx *gin.Context) {
	containerId := ctx.Query("containerId")
	docker.StartDockerContainer(containerId, ctx)
}

func StopDockerContainer(ctx *gin.Context) {
	containerId := ctx.Query("containerId")
	docker.StopDockerContainer(containerId, ctx)
}

func SelectDockerContainer(ctx *gin.Context) {
	containerId := ctx.Query("containerId")
	docker.SelectDockerContainer(containerId, ctx)
}

func SelectDockerImagesLabel(ctx *gin.Context) {
	docker.SelectDockerImagesLabel(ctx)
}

func SelectDockerNetworkLabel(ctx *gin.Context) {
	docker.SelectDockerNetworkLabel(ctx)
}

package controller

import (
	"Vnc_Server/docker"
	"Vnc_Server/model/response"
	model "Vnc_Server/model/vnc"
	"Vnc_Server/service"
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

func SearchDockerImage(ctx *gin.Context) {
	imageName := ctx.Query("name")
	docker.SearchDockerImage(imageName, ctx)
}

func PullDockerImage(ctx *gin.Context) {
	imageName := ctx.Query("name")
	docker.PullDockerImage(imageName, ctx)
}

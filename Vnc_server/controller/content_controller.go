package controller

import (
	"Vnc_Server/model/response"
	model "Vnc_Server/model/vnc"
	"Vnc_Server/service"
	"github.com/gin-gonic/gin"
)

func AddContent(ctx *gin.Context) {
	var content model.ContentInfo
	if err := ctx.ShouldBind(&content); err == nil {
		service.AddContent(&content, ctx)
		return
	}

	response.FailWithMessage("参数无效", ctx)
	return
}

func GetContentInfo(ctx *gin.Context) {
	name := ctx.Query("name")
	service.GetContentInfo(name, ctx)
}

func GetContentList(ctx *gin.Context) {
	service.GetContentList(ctx)
}

func GetContentCount(ctx *gin.Context) {
	service.GetContentCount(ctx)
}

func RemoveContentInfo(ctx *gin.Context) {
	name := ctx.Query("name")
	service.RemoveContentInfo(name, ctx)
}

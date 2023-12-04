package controller

import (
	"Social_Gin/model/response"
	model "Social_Gin/model/vnc"
	"Social_Gin/service"
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

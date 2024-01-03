package controller

import (
	"Vnc_Server/model/response"
	model "Vnc_Server/model/vnc"
	"Vnc_Server/model/vo"
	"Vnc_Server/service"
	"github.com/gin-gonic/gin"
)

// Register 用户注册
func Register(ctx *gin.Context) {
	var user model.UserInfo
	if err := ctx.ShouldBind(&user); err == nil {
		service.Register(&user, ctx)
		return
	}
	response.FailWithMessage("参数无效", ctx)
	return
}

// Login 用户登录
func Login(ctx *gin.Context) {
	var loginParam vo.LoginParam
	if err := ctx.ShouldBind(&loginParam); err == nil {
		service.Login(&loginParam, ctx)
		return
	}
	response.FailWithMessage("参数无效", ctx)
	return
}

func GetUserInfo(ctx *gin.Context) {
	username := ctx.Query("username")
	service.GetUserInfo(username, ctx)
}

func UpdateUserInfo(ctx *gin.Context) {
	var userInfo model.UserInfo
	if err := ctx.ShouldBind(&userInfo); err == nil {
		service.UpdateUserInfo(&userInfo, ctx)
		return
	}
	response.FailWithMessage("参数无效", ctx)
	return
}

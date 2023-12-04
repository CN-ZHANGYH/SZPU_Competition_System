package service

import (
	"Social_Gin/config"
	"Social_Gin/model/response"
	model "Social_Gin/model/vnc"
	"Social_Gin/model/vo"
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"strings"
)

const SALT = "123321Vvv#Salt"

func Register(userInfo *model.UserInfo, ctx *gin.Context) {
	if strings.EqualFold(userInfo.Username, "") ||
		strings.EqualFold(userInfo.Password, "") ||
		strings.EqualFold(userInfo.Email, "") ||
		strings.EqualFold(userInfo.Phone, "") {

		response.FailWithMessage("信息不能为空", ctx)
		return
	}
	var selectUserInfo model.UserInfo
	if SelectUserInfo(userInfo.Username, &selectUserInfo) > 0 {
		response.FailWithMessage("当前用户已经注册", ctx)
		return
	}

	hashedPassword := hashPassword(userInfo.Password, SALT)
	userInfo.Password = hashedPassword

	userInfo.Status = 1
	userInfo.Role = 1
	config.DB.Create(&userInfo)
	response.OkWithDetailed(userInfo, "注册成功", ctx)
}

func Login(loginParam *vo.LoginParam, ctx *gin.Context) {
	if strings.EqualFold(loginParam.Username, "") || strings.EqualFold(loginParam.Password, "") {
		response.FailWithMessage("用户密码不能为空", ctx)
		return
	}
	var selectUserInfo model.UserInfo
	if SelectUserInfo(loginParam.Username, &selectUserInfo) > 0 {
		if verifyPassword(loginParam.Password, selectUserInfo.Password, SALT) {
			response.OkWithDetailed(selectUserInfo, "登录成功", ctx)
			return
		}
	}
	response.FailWithMessage("用户或密码错误", ctx)
}

func SelectUserInfo(username string, userInfo *model.UserInfo) int64 {
	return config.DB.Where("username = ?", username).First(&userInfo).RowsAffected
}

// 对密码进行加密
func hashPassword(password string, salt string) string {
	hash := sha256.New()
	hash.Write([]byte(password + salt))
	return hex.EncodeToString(hash.Sum(nil))
}

// 验证密码
func verifyPassword(inputPassword string, storedPassword string, salt string) bool {
	return storedPassword == hashPassword(inputPassword, salt)
}

func GetUserInfo(username string, ctx *gin.Context) {
	var selectUserInfo model.UserInfo
	if SelectUserInfo(username, &selectUserInfo) > 0 {
		response.OkWithDetailed(selectUserInfo, "查询成功", ctx)
		return
	}
	response.FailWithMessage("查询失败", ctx)
	return

}

// UpdateUserInfo 可以更改用户密码和邮箱
func UpdateUserInfo(userInfo *model.UserInfo, c *gin.Context) {
	// 判断当前的用户名是否为空
	if userInfo.Username == "" {
		response.FailWithMessage("当前用户为空", c)
		return
	}
	// 查询一下用户是否存在
	count := config.DB.Model(&userInfo).Updates(userInfo).RowsAffected
	if count > 0 {
		response.OkWithDetailed(userInfo, "更新成功", c)
		return
	}
	response.FailWithMessage("更新失败", c)
	return

}

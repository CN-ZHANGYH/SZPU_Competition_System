package controller

import (
	"Vnc_Server/model/response"
	model "Vnc_Server/model/vnc"
	"Vnc_Server/service"
	"github.com/gin-gonic/gin"
)

func AddVm(c *gin.Context) {
	var vmInfo model.VmInfo
	if err := c.ShouldBind(&vmInfo); err == nil {
		service.AddVmInfo(&vmInfo, c)
		return
	}
	response.FailWithMessage("数据解析失败", c)
	return
}

func GetVm(c *gin.Context) {
	name := c.Query("name")
	service.GetVmInfo(name, c)
}

func GetVmHostAndUrl(c *gin.Context) {
	service.GetVmHostAndUrl(c)
}

func DeleteVm(c *gin.Context) {
	name := c.Query("name")
	service.DeleteVm(name, c)
}

func GetVmList(c *gin.Context) {
	service.GetVmList(c)
}

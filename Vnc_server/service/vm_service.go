package service

import (
	"Vnc_Server/config"
	"Vnc_Server/docker"
	"Vnc_Server/model/response"
	model "Vnc_Server/model/vnc"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

// AddVmInfo 添加虚拟机信息
func AddVmInfo(vmInfo *model.VmInfo, context *gin.Context) {
	if strings.EqualFold(vmInfo.Host, "") ||
		strings.EqualFold(vmInfo.Name, "") ||
		strings.EqualFold(vmInfo.Password, "") ||
		vmInfo.Port == 0 {
		response.FailWithMessage("当前的主机信息不能为空", context)
		return
	}
	var selectVmInfo model.VmInfo
	if SelectVmInfo(vmInfo.Name, vmInfo.Port, &selectVmInfo) > 0 {
		if strings.EqualFold(vmInfo.Name, selectVmInfo.Name) {
			response.FailWithMessage("当前的主机信息已存在", context)
			return
		}
	}
	vmInfo.Url = "http://" + vmInfo.Host + ":" + strconv.Itoa(vmInfo.Port) + "/vnc.html"
	config.DB.Create(&vmInfo)
	response.OkWithMessage("添加成功", context)
	return
}

func SelectVmInfo(name string, port int, vmInfo *model.VmInfo) int64 {
	return config.DB.Where("name = ? or port = ?", name, port).First(&vmInfo).RowsAffected
}

func GetVmInfo(name string, ctx *gin.Context) {
	if name == "" {
		response.FailWithMessage("当前主机名称不能为空", ctx)
		return
	}
	var selectVmInfo model.VmInfo
	if config.DB.Where("name = ?", name).First(&selectVmInfo).RowsAffected > 0 {
		response.OkWithDetailed(selectVmInfo, "查询成功", ctx)
		return
	}
	response.FailWithMessage("查询失败", ctx)
	return
}

func GetVmList(ctx *gin.Context) {
	var vmList []model.VmInfo
	if config.DB.Find(&vmList).RowsAffected > 0 {
		var resultMap = make(map[string]interface{})
		resultMap["data"] = vmList
		resultMap["total"] = len(vmList)

		response.OkWithDetailed(resultMap, "查询成功", ctx)
		return
	}
	response.FailWithMessage("还未添加虚拟机", ctx)
	return
}

func DeleteVm(name string, ctx *gin.Context) {
	if strings.EqualFold(name, "") {
		response.FailWithMessage("参数无效", ctx)
		return
	}
	var selectVmInfo model.VmInfo
	if config.DB.Where("name = ?", name).First(&selectVmInfo).RowsAffected > 0 {
		config.DB.Delete(&selectVmInfo)
		response.OkWithMessage("删除成功", ctx)
		return
	}
	response.FailWithMessage("删除失败", ctx)
	return

}

func GetVmHostAndUrl(ctx *gin.Context) {
	var vmList []model.VmInfo
	if config.DB.Find(&vmList).RowsAffected > 0 {
		var resultList []map[string]interface{}
		for i := 0; i < len(vmList); i++ {
			var resultMap = make(map[string]interface{})
			resultMap["label"] = vmList[i].Name
			resultMap["value"] = vmList[i].Url

			resultList = append(resultList, resultMap)
		}
		response.OkWithDetailed(resultList, "查询成功", ctx)
		return
	}
}

// CreateVmInfo 创建Docker的虚拟机环境
func CreateVmInfo(dockerVmInfo *model.DockerVmInfo, ctx *gin.Context) {
	if strings.EqualFold(dockerVmInfo.ContainerName, "") ||
		strings.EqualFold(dockerVmInfo.ContainerImage, "") {
		response.FailWithMessage("参数无效", ctx)
		return
	}
	if err := docker.CreateVmContainer(dockerVmInfo, ctx); err == nil {
		response.OkWithMessage("创建容器成功", ctx)
		return
	}
	response.FailWithMessage("创建容器失败", ctx)
	return

}

func GetDockerVmList(ctx *gin.Context) {
	docker.GetDockerContainerList(ctx)
}

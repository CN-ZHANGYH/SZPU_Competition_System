package service

import (
	"Social_Gin/config"
	"Social_Gin/model/response"
	model "Social_Gin/model/vnc"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strings"
)

func AddContent(content *model.ContentInfo, ctx *gin.Context) {
	if strings.EqualFold(content.Name, "") ||
		strings.EqualFold(content.Type, "") ||
		strings.EqualFold(content.VmList, "") ||
		strings.EqualFold(content.Content, "") {

		response.FailWithMessage("参数无效", ctx)
		return
	}
	var selectContent model.ContentInfo
	if SelectContentByName(content.Name, &selectContent) > 0 {
		response.FailWithMessage("当前试题已存在", ctx)
		return
	}
	var result map[string]interface{}
	err := json.Unmarshal([]byte(content.Content), &result)
	if err == nil {
		newContent := result["_rawValue"]
		content.Content = newContent.(string)
	}
	config.DB.Create(&content)
	response.OkWithMessage("添加成功", ctx)
	return

}

func SelectContentByName(name string, content *model.ContentInfo) int64 {
	return config.DB.Where("name = ?", name).First(content).RowsAffected
}

func GetContentInfo(name string, ctx *gin.Context) {
	if strings.EqualFold(name, "") {
		response.FailWithMessage("参数无效", ctx)
		return
	}
	var content model.ContentInfo
	if SelectContentByName(name, &content) > 0 {
		response.OkWithDetailed(content, "查询成功", ctx)
		return
	}
	response.FailWithMessage("查询失败", ctx)
	return
}

func GetContentList(ctx *gin.Context) {
	var contentList []model.ContentInfo
	if config.DB.Find(&contentList).RowsAffected > 0 {
		var resultMap = make(map[string]interface{})
		resultMap["data"] = contentList
		resultMap["total"] = len(contentList)

		response.OkWithDetailed(resultMap, "查询成功", ctx)
		return
	}
	response.FailWithMessage("还未添加试题", ctx)
	return
}

func GetContentCount(ctx *gin.Context) {
	var contentList []model.ContentInfo
	if config.DB.Find(&contentList).RowsAffected > 0 {
		var resultList []map[string]interface{}
		for i := 0; i < len(contentList); i++ {
			var resultMap = make(map[string]interface{})
			resultMap["label"] = contentList[i].Name
			resultMap["value"] = contentList[i].Name
			resultList = append(resultList, resultMap)
		}
		response.OkWithDetailed(resultList, "查询成功", ctx)
		return
	}
}

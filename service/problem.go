package service

import (
	"log"
	"net/http"
	"strconv"

	"gin-ojsys.cn/config"
	"gin-ojsys.cn/models"
	"github.com/gin-gonic/gin"
)

// GetProblemList 获取题目列表
// @Tags 公共方法
// @Summary 获取题目列表
// @Param page query int false "页码"
// @Param size query int false "每页数量"
// @Param keyword query string false "关键字"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /problems [get]
func GetProblemList(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", config.AppSetting.DefaultPage))
	size, _ := strconv.Atoi(c.DefaultQuery("size", config.AppSetting.DefaultPageSize))
	keyword := c.Query("keyword")

	data := make([]*models.Problem, 0)

	tx := models.GetProblemList(keyword)
	offset := (page - 1) * size
	var count int64
	err := tx.Count(&count).Offset(offset).Limit(size).Find(&data).Error
	if err != nil {
		log.Println("GetProblemList Error : ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
		"page_info": map[string]interface{}{
			"page":  page,
			"size":  size,
			"total": count,
		},
	})
}

// GetProblemDetail 获取题目详情
// @Tags 公共方法
// @Summary 获取题目详情
// @Param id path int true "题目ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /problems/{id} [get]
func GetProblemDetail(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	data := &models.Problem{}

	tx := models.GetProblemDetail(id)
	err := tx.First(data).Error
	if err != nil {
		log.Println("GetProblemDetail Error : ", err)
		data = nil
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})
}

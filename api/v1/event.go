package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wejectchen/ginblog/model"
	"github.com/wejectchen/ginblog/utils/errmsg"
)

// 添加分类
func AddEvent(c *gin.Context) {
	var data model.Event
	_ = c.ShouldBindJSON(&data)
	if data.Name == "" || data.Host == "" || data.When == "" {
		c.JSON(
			http.StatusOK, gin.H{
				"status": 404,
				// "data":    data,
				"message": "資料填寫不完整",
			},
		)
		return
	}

	code := model.CheckEvent(data.Name)
	if code == errmsg.SUCCSE {
		model.CreateEvent(&data)
	} else {
		data = model.Event{}
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// 查询分类信息
func GetEventInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data, code := model.GetEventInfo(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		},
	)

}

// 查询分类列表
func GetEvent(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, total := model.GetEvent(pageSize, pageNum)
	code := errmsg.SUCCSE
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// EditCate 编辑分类名
func EditEvent(c *gin.Context) {
	var data model.Event
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code := model.CheckEvent(data.Name)
	if code == errmsg.SUCCSE {
		model.EditEvent(id, &data)
	}
	if code == errmsg.ERROR_CATENAME_USED {
		c.Abort()
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// DeleteCate 删除用户
func DeleteEvent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := model.DeleteEvent(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

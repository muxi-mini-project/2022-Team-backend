package handler

import (
	"team/model"

	"github.com/gin-gonic/gin"
)

// @Summary “任务待办”
// @Description “获取未完成任务”
// @Tags tasklist
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200
// @Failure 401 "身份验证失败"
// @Failure 400 "获取失败"
// @Router /info/todolist [get]
func ToDoList(c *gin.Context) {
	temp, ok := c.Get("id")
	id := temp.(int)
	if !ok {
		c.JSON(401, gin.H{
			"code":    "401",
			"message": "身份验证失败",
		})
	}
	Userinfo, err := model.GetUserInfo(id)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "获取失败",
		})
	}
	ToDoList, team := model.GenToDoList(Userinfo.UserId)
	c.JSON(200, gin.H{
		"code": 200,
		"task": ToDoList,
		"team": team,
	})
}

// @Summary “已完成任务”
// @Description “获取已完成任务”
// @Tags tasklist
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200
// @Failure 401 "验证失败"
// @Failure 400 "获取失败"
// @Router /info/donelist [get]
func DoneList(c *gin.Context) {
	temp, ok := c.Get("id")
	id := temp.(int)
	if !ok {
		c.JSON(401, gin.H{
			"code":    "401",
			"message": "身份验证失败",
		})
	}
	Userinfo, err := model.GetUserInfo(id)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "获取失败",
		})
	}
	DoneList, team := model.GenDoneList(Userinfo.UserId)
	c.JSON(200, gin.H{
		"code": 200,
		"task": DoneList,
		"team": team,
	})
}

// @Summary “完成任务”
// @Description “打对钩完成任务”
// @Tags tasklist
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param task_id path string true "任务id"
// @Success 200 "任务完成"
// @Failure 401 "任务完成失败"
// @Failure 401 "验证失败"
// @Router /info/donetask/:task_id [put]
func DoneTask(c *gin.Context) {
	temp, ok := c.Get("id")
	uId := temp.(int)
	if !ok {
		c.JSON(401, gin.H{
			"code":    "401",
			"message": "身份验证失败",
		})
	}
	id := c.Param("task_id")
	if err := model.CompleteTask(id, uId); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "任务完成失败",
		})
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "任务完成",
	})
}

// @Summary “取消任务的完成”
// @Description “取消对钩”
// @Tags tasklist
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param task_id path string true "任务id"
// @Success 200
// @Failure 401 "验证失败"
// @Router /info/donetask [put]
func CancelDone(c *gin.Context) {
	temp, ok := c.Get("id")
	uId := temp.(int)
	if !ok {
		c.JSON(401, gin.H{
			"code":    "401",
			"message": "身份验证失败",
		})
	}

	id := c.Param("task_id")
	if err := model.CancelComplete(id, uId); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "任务取消完成失败",
		})
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "任务取消完成成功",
	})
}

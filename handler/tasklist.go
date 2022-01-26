package handler

import (
	"fmt"
	"team/model"

	"github.com/gin-gonic/gin"
)

// @Summary “任务待办”
// @Description “获取未完成任务”
// @Tags my
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200
// @Failure 401 "验证失败"
// @Failure 404 "获取失败"
// @Router /info/todolist [get]
func ToDoList(c *gin.Context) {
	token := c.Request.Header.Get("token")
	phone, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": "验证失败"})
	}
	Userinfo, err := model.GetUserInfo(phone)
	if err != nil {
		c.JSON(404, gin.H{"message": "获取失败"})
	}
	ToDoList := model.GenToDoList(Userinfo.UserId)
	c.JSON(200, gin.H{"data": ToDoList})
}

// @Summary “已完成任务”
// @Description “获取已完成任务”
// @Tags my
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200
// @Failure 401 "验证失败"
// @Failure 404 "获取失败"
// @Router /info/donelist [get]
func DoneList(c *gin.Context) {
	token := c.Request.Header.Get("token")
	phone, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": "验证失败"})
	}
	Userinfo, err := model.GetUserInfo(phone)
	if err != nil {
		c.JSON(404, gin.H{"message": "获取失败"})
	}
	DoneList := model.GenToDoList(Userinfo.UserId)
	c.JSON(200, gin.H{"data": DoneList})
}

// @Summary “完成任务”
// @Description “打对钩完成任务”
// @Tags my
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param tid Path string true "tid"
// @Success 200
// @Failure 401 "验证失败"
// @Router /info/donetask [put]
func DoneTask(c *gin.Context) {
	token := c.Request.Header.Get("token")
	_, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": "验证失败"})
	}
	id := c.Param("tid")
	if err := model.CompleteTask(id); err != nil {
		fmt.Println(err)
	}
}

// @Summary “取消任务的完成”
// @Description “取消对钩”
// @Tags my
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param tid Path string true "tid"
// @Success 200
// @Failure 401 "验证失败"
// @Router /info/donetask [put]
func CancelDone(c *gin.Context) {
	token := c.Request.Header.Get("token")
	_, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": "验证失败"})
	}
	id := c.Param("tid")
	if err := model.CancelComplete(id); err != nil {
		fmt.Println(err)
	}
}

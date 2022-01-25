package handler

import (
	"2022-TEAM-BACKEND/model"
	"fmt"

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
	c.JSON(200, gin.H{"message": ToDoList})
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
	c.JSON(200, gin.H{"message": DoneList})
}

// @Summary “已完成任务”
// @Description “获取已完成任务”
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
	if err := model.DB.Table("user_task").Where("id=?", id).Updates(map[string]interface{}{"performance": true}).Error; err != nil {
		fmt.Println(err)
	}
}

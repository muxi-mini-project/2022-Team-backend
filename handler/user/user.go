package handler

import (
	"2022-TEAM-BACKEND/model"

	"github.com/gin-gonic/gin"
)

// @Summary "注册"
// @Description "注册一个新用户"
// @tags user
// @Accept json
// @Produce json
// @Param user body model.Users "true"
// @Success 200 "注册成功"
// @Failure 400 "输入有误，格式错误"
// @Router /user [post]

func User(c *gin.Context) {

	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"message": "输入有误，格式错误"})
		return
	}

}

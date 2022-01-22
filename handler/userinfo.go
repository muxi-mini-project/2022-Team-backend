package handler

import (
	"2022-TEAM-BACKEND/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

// @Summary "用户界面"
// @Description "获取用户的的基本信息"
// @Tags my
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} model.Userinfo
// @Failure 404 "获取失败"
// Router /info [get]

func Userinfo(c *gin.Context) {
	//注意这里token要手写到header里（因为是客户端的工作）
	token := c.Request.Header.Get("token")
	// fmt.Println(token)
	phone, err := model.VerifyToken(token)
	fmt.Println(phone)
	if err != nil {
		c.JSON(401, gin.H{"message": "验证失败"})
	}
	Userinformation, err := model.GetUserInfo(phone)
	if err != nil {
		c.JSON(404, gin.H{"message": "获取失败"})
	}
	c.JSON(200, Userinformation)
}

// @Summary “修改用户的信息”
// @Description “修改用户的基本信息”
// @Tags my
// @Accept json
// @Produce json
// @Param user body model.Uer true "user"
// @Param token header string true "token"
// @Success 200 "修改成功"
// @Failure 401 "验证失败"
// @Failure 400 "修改失败"
// @Router /info [put]

func ChangeInfomation(c *gin.Context) {
	var user model.User
	token := c.Request.Header.Get("token")
	phone, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": "验证失败"})
	}
	if err1 := c.BindJSON(&user); err1 != nil {
		c.JSON(400, gin.H{"message": "输入格式有误"})
		return
	}
	user.Phone = phone
	if user.NickName == "" {
		c.JSON(400, gin.H{"message": "用户名不可为空!"})
		return
	}
	//for range 键值循环
	for _, char := range user.NickName {
		if string(char) == " " {
			c.JSON(400, gin.H{"message": "用户名中不可含有空格"})
			return
		}
	}
	// if user.Password == "" {
	// 	c.JSON(400, gin.H{"message": "密码不可为空!"})
	// 	return
	// }
	// //for range 键值循环
	// for _, num := range user.Password {
	// 	if string(num) == " " {
	// 		c.JSON(400, gin.H{"message": "密码中不可含有空格"})
	// 		return
	// 	}
	// }
	if err2 := model.ChangeUserInfo(user); err2 != nil {
		c.JSON(400, gin.H{"message": "修改失败"})
		return
	}
	c.JSON(200, gin.H{"message": "修改成功"})
}

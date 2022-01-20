package handler

import (
	"2022-TEAM-BACKEND/model"
	"fmt"

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
	// if _,a := model.IfExistUserPhone(user.Phone);
	//电话位数问题交给前端
	if _, a := model.IfExistUserPhone(user.Phone); a != 1 {
		c.JSON(200, gin.H{
			"message": "对不起，该电话号码已经被绑定",
		})
	}
	user_id := model.Register(user.Phone, user.Password)
	fmt.Println(user.Phone)
	c.JSON(200, gin.H{
		"user_id": user_id,
	})
	// 	//注册后的下一个页面输入信息
	// 	if user_id != " " {
	// 		if err := c.BindJSON(&user); err != nil {
	// 			c.JSON(400, gin.H{
	// 				"message": "输入有误，格式错误"})
	// 			return
	// 		}
	// 		err1 := model.InitInfo(user.NickName, user.Password)
	// 		fmt.Println(err1)
	// 	}
}

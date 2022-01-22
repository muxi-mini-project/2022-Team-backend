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
	fmt.Println(user.Phone)
	if _, a := model.IfExistUserPhone(user.Phone); a != 1 {
		// user.Phone
		c.JSON(200, gin.H{
			"message": "对不起，该电话号码已经被绑定",
		})
		return
	}
	user_id := model.Register(user.Phone, user.Password)
	fmt.Println(user.Phone)
	c.JSON(200, gin.H{
		"user_id": user_id,
	})
}

// 注册后的下一个页面输入信息,这应该也要token了吧

// if user_id != " " {
// 	if err := c.BindJSON(&user); err != nil {
// 		c.JSON(400, gin.H{
// 			"message": "输入有误，格式错误"})
// 		return
// 	} else {
// 		err1 := model.InitInfo(user.NickName, user.Avatar)
// 		fmt.Println(err1)
// 	}
// }

// @Summary "登入"
// @Describtion 验证用户信息实现登入
// @Tags login
// @Accept json
// @Producer json
// @Param user body model.Userinfo true "user"
// @Success 200 {object} model.Token "登陆成功"
// @Failure 400 "输入格式错误"
// @Failure 404 "用户不存在"
// @Failure 401 "密码错误"
// @Router /login[post]

func Login(c *gin.Context) {
	var user model.User
	//BindJSON把前端的数据写到user里
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"message": "输入格式有误"})
		return
	}

	fmt.Println(user.Phone, user.Password)
	//验证用户是否存在（电话是否已经注册）
	if model.VerifyPhone(user.Phone) != false {
		c.JSON(404, gin.H{"message": "用户不存在"})
		return
	}
	//验证密码,密码可能重复所以还要电话
	if model.VerifyPassword(user.Phone, user.Password) == false {
		c.JSON(401, gin.H{"message": "密码错误"})
		return
	} else {
		// user.UserId = model.GetId{user.Phone}
		c.JSON(200, gin.H{
			"message": "登陆成功",
			"token":   model.GenerateToken(user.Phone),
		})
		token := model.GenerateToken(user.Phone)
		fmt.Println(token)
		return
	}
}

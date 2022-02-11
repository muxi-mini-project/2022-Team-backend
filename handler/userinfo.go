package handler

import (
	"fmt"
	"team/model"

	"github.com/gin-gonic/gin"
)

// @Summary "用户界面"
// @Description "获取用户的的基本信息"
// @Tags userinfo
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 "信息获取成功"
// @Failure 401 "身份验证失败"
// @Failure 400 "获取失败"
// Router /user/info [get]
func Userinfo(c *gin.Context) {
	//注意这里token要手写到header里（客户端）
	temp, ok := c.Get("id")
	id := temp.(int)
	if !ok {
		c.JSON(401, gin.H{
			"code":    401,
			"message": "身份验证失败",
		})
	}
	Userinformation, err := model.GetUserInfo(id)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "获取失败",
		})
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    Userinformation,
	})
}

// @Summary “修改用户的信息”
// @Description “修改用户的基本信息”
// @Tags userinfo
// @Accept json
// @Produce json
// @Param user body model.User true "user"
// @Param token header string true "token"
// @Param user body model.User true "输入昵称"
// @Success 200 "修改成功"
// @Failure 401 "验证失败"
// @Failure 400 "修改失败"
// @Router /user/info [put]
func ChangeNickname(c *gin.Context) {
	var user model.User
	temp, ok := c.Get("id")
	id := temp.(int)
	if !ok {
		c.JSON(401, gin.H{
			"code":    401,
			"message": "身份验证失败",
		})
	}
	if err1 := c.BindJSON(&user); err1 != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "输入有误",
		})
		return
	}
	oldInfo, _ := model.GetUserInfo(id)
	user.Phone = oldInfo.Phone
	if user.NickName == "" {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "用户名不可为空!",
		})
		return
	}
	//for range 键值循环
	for _, char := range user.NickName {
		if string(char) == " " {
			c.JSON(400, gin.H{
				"code":    400,
				"message": "用户名中不可含有空格",
			})
			return
		}
	}
	if err2 := model.ChangeUserInfo(user); err2 != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "修改失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "修改成功",
	})
}

// @Summary “验证用户密码”
// @Description “修改密码前对密码的验证功能”
// @Tags userinfo
// @Accept json
// @Produce json
// @Param user body model.User true "输入密码"
// @Param token header string true "token"
// @Success 200 "验证成功"
// @Failure 401 "验证失败"
// @Failure 400 "输入有误"
// @Router /user/change_password/verify [get]
func VerifyPassword(c *gin.Context) {
	var user model.User
	temp, ok := c.Get("id")
	id := temp.(int)
	if !ok {
		c.JSON(401, gin.H{
			"code":    401,
			"message": "身份验证失败",
		})
	}
	Userinformation, err := model.GetUserInfo(id)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "获取失败",
		})
	}
	if err1 := c.BindJSON(&user); err1 != nil {
		fmt.Println(err1)
		c.JSON(400, gin.H{
			"code":    400,
			"message": "输入格式有误",
		})
		return
	}
	if user.Password != Userinformation.Password {
		c.JSON(401, gin.H{
			"code":    401,
			"message": "验证失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "验证成功",
	})
}

// @Summary “修改用户密码”
// @Description “修改密码”
// @Tags userinfo
// @Accept json
// @Produce json
// @Param password body model.Password true "输入两次新密码"
// @Param token header string true "token"
// @Success 200 "修改成功"
// @Failure 401 "验证失败"
// @Failure 400 "修改失败"
// @Router /user/change_password/change [post]
func ChangePassword(c *gin.Context) {
	// var user model.User
	temp, ok := c.Get("id")
	id := temp.(int)
	if !ok {
		c.JSON(401, gin.H{
			"code":    401,
			"message": "身份验证失败",
		})
	}
	var password model.Password
	if err1 := c.BindJSON(&password); err1 != nil {
		c.JSON(401, gin.H{
			"code":    401,
			"message": "输入格式有误",
		})
		return
	}
	if password.ConfirmPassword != password.NewPassword {
		c.JSON(401, gin.H{
			"code":    401,
			"message": "两次输入不一致",
		})
		return
	} else {
		err := model.ModifyPassword(id, password.ConfirmPassword)
		if err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"message": "修改失败",
			})
		}
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "修改成功",
	})
}

// @Summary “用户反馈”
// @Description “用户反馈”
// @Tags userinfo
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param user body model.User true "输入反馈信息"
// @Success 200 "反馈成功"
// @Failure 401 "验证失败"
// @Failure 400 "输入格式有误"
// @Router /user/feedback [put]
func Feedback(c *gin.Context) {
	// ID := c.MustGet("student_id").(string)
	temp, ok := c.Get("id")
	id := temp.(int)
	if !ok {
		c.JSON(401, gin.H{
			"code":    401,
			"message": "身份验证失败",
		})
	}
	var user model.User
	user.UserId = id
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "输入有误,格式错误",
		})
		return
	}
	err := model.UserFeedback(user.UserId, user.Feedback)
	if err == nil {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "反馈成功",
		})
	}
}

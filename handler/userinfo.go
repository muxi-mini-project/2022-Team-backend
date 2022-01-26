package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"team/model"
	"team/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary "用户界面"
// @Description "获取用户的的基本信息"
// @Tags my
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} model.Userinfo
// @Failure 400 "获取失败"
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
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": "验证失败"})
	}
	if err1 := c.BindJSON(&user); err1 != nil {
		c.JSON(400, gin.H{"message": "输入格式有误"})
		return
	}
	oldInfo, _ := model.GetUserInfo(id)
	user.Phone = oldInfo.Phone
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
	if err2 := model.ChangeUserInfo(user); err2 != nil {
		c.JSON(400, gin.H{"message": "修改失败"})
		return
	}
	c.JSON(200, gin.H{"message": "修改成功"})
}

// @Summary “验证用户密码”
// @Description “修改密码前对密码的验证功能”
// @Tags my
// @Accept json
// @Produce json
// @Param user.Password body model.Uer true "user.Password"
// @Param token header string true "token"
// @Success 200 "验证成功"
// @Failure 401 "验证失败"
// @Failure 400  "输入格式有误"
// @Failure 404 "获取失败"
// @Router /change_password/verify [get]
func VerifyPassword(c *gin.Context) {
	var user model.User
	token := c.Request.Header.Get("token")
	phone, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": "验证失败"})
	}
	Userinformation, err := model.GetUserInfo(phone)
	if err != nil {
		c.JSON(404, gin.H{"message": "获取失败"})
	}
	if err1 := c.BindJSON(&user.Password); err1 != nil {
		c.JSON(400, gin.H{"message": "输入格式有误"})
		return
	}
	if user.Password != Userinformation.Password {
		c.JSON(401, gin.H{"message": "验证失败"})
		return
	}
	c.JSON(200, gin.H{"message": "验证成功"})
}

// @Summary “修改用户密码”
// @Description “修改密码”
// @Tags my
// @Accept json
// @Produce json
// @Param password body model.Password true "password"
// @Param token header string true "token"
// @Success 200 "修改成功"
// @Failure 401 "验证失败"
// @Failure 402 "输入格式有误"
// @Failure 403 "两次输入不一致"
// @Failure 400 "修改失败"
// @Router /change_password/change [post]
func ChangePassword(c *gin.Context) {
	// var user model.User
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": "验证失败"})
	}
	var password model.Password
	if err1 := c.BindJSON(&password); err1 != nil {
		c.JSON(402, gin.H{"message": "输入格式有误"})
		return
	}
	if password.ConfirmPassword != password.NewPassword {
		c.JSON(403, gin.H{"message": "两次输入不一致"})
		return
	} else {
		err := model.ModifyPassword(id, password.ConfirmPassword)
		if err != nil {
			c.JSON(400, gin.H{"message": "修改失败"})
		}
	}
	c.JSON(200, gin.H{"message": "修改成功"})
}

func ModifyProfile(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": err.Error(),
		})
		log.Panicln("绑定发生错误", err.Error())
	}
	file, err := c.FormFile("avatar-file")
	if err != nil {
		c.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": err.Error(),
		})
		log.Panicln("文件上传错误", err.Error())
	}
	path := utils.RootPath()
	path = filepath.Join(path, "avatar")
	fmt.Println("path = >", path)
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		c.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": err,
		})
		log.Panicln("无法创建文件夹", err.Error())
	}
	fileName := strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	err = c.SaveUploadedFile(file, filepath.Join(path, fileName))
	if err != nil {
		c.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": err,
		})
		log.Panicln("无法保存文件", err.Error())
	}
	avatarUrl := "/avatar/" + fileName
	user.Avatar = sql.NullString{String: avatarUrl}
	err = model.UpdateAvator(user.UserId)
	if err != nil {
		c.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": err,
		})
		log.Panicln("数据无法更新", err.Error())
	}
	id := strconv.Itoa(user.UserId)
	c.Redirect(http.StatusMovedPermanently, "/user/profile?id="+id)
}

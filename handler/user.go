package handler

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"path"
	"strconv"
	"team/model"
	"team/services"
	"team/services/connector"

	"github.com/gin-gonic/gin"
)

// @Summary "登录"
// @Tags user
// @Description "一站式登录"
// @Accept json
// @Produce json
// @Param user body model.User true "输入学号，密码进行登录"
// @Success 200  "将用户id作为token保留"
// @Failure 401  "身份认证失败 重新登录"
// @Failure 400  "输入有误"
// @Router /login [post]
func Login(c *gin.Context) {
	var u model.User
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Lack Param Or Param Not Satisfiable."})
		return
	}
	if u.StudentId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Lack Param Or Param Not Satisfiable."})
		return
	}
	pwd := u.Password
	//首次登录，验证一站式
	//判断是否首次登录
	result := model.DB.Table("user").Where("student_id = ?", u.StudentId).First(&u)
	if result.Error != nil {
		I, err := model.GetUserInfoFormOne(u.StudentId, pwd)
		fmt.Println(I)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "Password or account is wrong.")
			return
		}
		//对用户信息初始化
		u.NickName = " "
		//对密码进行base64加密
		u.Password = base64.StdEncoding.EncodeToString([]byte(u.Password))
		model.DB.Table("user").Create(&u)
		model.DB.Table("user").Where("student_id = ?", u.StudentId).Select("id").Find(&u.UserId)
	} else {
		//在数据库中解密比较
		password, _ := base64.StdEncoding.DecodeString(u.Password)
		model.DB.Table("user").Where("student_id = ?", u.StudentId).Select("id").Find(&u.UserId)
		if string(password) != pwd {
			c.JSON(http.StatusUnauthorized, "password or account is wrong.")
			return
		}
	}
	fmt.Println(u.UserId)
	signedToken := model.GenerateToken(u.UserId)
	c.JSON(200, gin.H{
		"code":    200,
		"message": "将用户id作为token保留",
		"data":    signedToken,
	})
}

// @Summary "修改头像"
// @Tags user
// @Description "修改用户头像"
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param file formData file true "文件"
// @Success 200 "上传成功"
// @Failure 401 "身份验证失败"
// @Failure 400 "上传失败"
// @Router /user/pupup/avatar [post]
// @Router /user/avatar [put]
func ModifyProfile(c *gin.Context) {

	var user model.User
	id := c.MustGet("id").(int)
	user, err := model.GetUserInfo(id)
	if err != nil {
		fmt.Println(err)
	}
	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(400, gin.H{
			"code":    401,
			"message": "上传失败!",
		})
		return
	}

	filepath := "./"
	if _, err := os.Stat(filepath); err != nil {
		if !os.IsExist(err) {
			os.MkdirAll(filepath, os.ModePerm)
		}
	}

	fileExt := path.Ext(filepath + file.Filename)

	id1 := strconv.Itoa(id)

	file.Filename = id1 + user.StudentId + fileExt

	filename := filepath + file.Filename

	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.JSON(400, gin.H{
			"code":    401,
			"message": "上传失败!",
		})
		return
	}

	// 删除原头像
	// user, _ := model.GetUserInfo(id)
	if user.Path != "" && user.Sha != "" {
		connector.RepoCreate().Del(user.Path, user.Sha)
	}

	// 上传新头像
	Base64 := services.ImagesToBase64(filename)
	picUrl, picPath, picSha := connector.RepoCreate().Push(file.Filename, Base64)

	os.Remove(filename)
	var avatar model.User
	avatar.UserId = id
	avatar.Avatar = picUrl
	avatar.Path = picPath
	avatar.Sha = picSha
	err0 := model.UpdateAvator(avatar)
	if picUrl == "" || err0 != nil {
		c.JSON(401, gin.H{
			"message": "上传失败",
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": "上传成功",
		"url":     picUrl,
		"sha":     picSha,
		"path":    picPath,
	})

}

// @Summary "初始化用户信息"
// @Description "再点击“完成设置”之前头像已经设置完，注册后弹窗里输入昵称"
// @tags user
// @Accept json
// @Produce json
// @Param user body model.User true "输入昵称"
// @Param token header string true "token"
// @Success 200 "注册成功"
// @Failure 400 "输入有误"
// @Failure 401 "身份验证失败"
// @Router /user/pupup [post]
func InitUserInfo(c *gin.Context) {
	// id := c.MustGet("id").(int)
	// temp := c.Request.Header.Get("id")
	// id, _ := strconv.Atoi(temp)
	id := c.MustGet("id").(int)

	var user model.User
	user.UserId = id
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"message": "输入有误，格式错误"})
		return
	}
	if user.NickName == "" {
		c.JSON(400, gin.H{"message": "用户名不可为空!"})
		return
	}
	if _, a := model.IfExistNickname(user.NickName); a != 1 {
		// user.Phone
		c.JSON(200, gin.H{
			"code":    400,
			"message": "对不起，该用户名已被注册",
		})
		return
	}

	err1 := model.InitInfo(user.UserId, user.NickName)

	if err1 == nil {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "注册成功",
		})
		return
	}
}

//下面这一段是将头像上传至本地，先留存以备日后类似需求
// func ModifyProfile(c *gin.Context) {
// file, err := c.FormFile("imgfile")
// if err != nil {
// 	c.JSON(200, gin.H{
// 		"code": 400,
// 		"msg":  "上传失败!",
// 	})
// 	return
// } else {
// 	fileExt := strings.ToLower(path.Ext(file.Filename))
// 	if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".gif" && fileExt != ".jpeg" {
// 		c.JSON(200, gin.H{
// 			"code": 400,
// 			"msg":  "上传失败!只允许png,jpg,gif,jpeg文件",
// 		})
// 		return
// 	}
// 	dst := path.Join("./", file.Filename)
// 	c.SaveUploadedFile(file, dst)
// 	c.JSON(200, gin.H{
// 		"code": 200,
// 		"msg":  "上传成功!",
// 		"result": gin.H{
// 			"path": dst,
// 		},
// 	})
// 	// id := c.MustGet("id").(int)
// 	temp := c.Request.Header.Get("id")
// 	id, _ := strconv.Atoi(temp)
// 	if err := model.UpdateAvator(id, dst); err != nil {
// 		fmt.Println(err)
// 	}
// }

// // @Summary "注册"
// // @Description "注册一个新用户"
// // @tags user
// // @Accept json
// // @Produce json
// // @Param user body model.User true "user"
// // @Success 200 "用户创建成功"
// // @Failure 400 "输入有误，格式错误"
// // @Failure 401 "电话号码重复"
// // @Router /user [post]
// func User(c *gin.Context) {
// 	var user model.User
// 	if err := c.BindJSON(&user); err != nil {
// 		c.JSON(400, gin.H{
// 			"code":    400,
// 			"message": "输入有误，格式错误"})
// 		return
// 	}
// 	//电话位数问题前端处理
// 	fmt.Println(user.Phone)
// 	if _, a := model.IfExistUserPhone(user.Phone); a != 1 {
// 		c.JSON(401, gin.H{
// 			"code":    401,
// 			"message": "对不起，该电话号码已经被绑定",
// 		})
// 		return
// 	}
// 	user_id := model.Register(user.Phone, user.Password)
// 	fmt.Println(user.Phone)
// 	c.JSON(200, gin.H{
// 		"code":    200,
// 		"message": "用户创建成功",
// 		"user_id": user_id,
// 	})
// }

// // @Summary "登录"
// // @Describtion "输入电话密码验证用户信息实现登入"
// // @Tags user
// // @Accept json
// // @Producer json
// // @Param token header string true "token"
// // @Param user body model.User true "user"
// // @Success 200  "登陆成功"
// // @Failure 400 "输入格式错误"
// // @Failure 404 "用户不存在"
// // @Failure 401 "密码错误"
// // @Router /login [post]
// func Login(c *gin.Context) {
// 	var user model.User
// 	//BindJSON把前端的数据写到user里
// 	if err := c.BindJSON(&user); err != nil {
// 		c.JSON(400, gin.H{
// 			"code":    400,
// 			"message": "输入格式有误",
// 		})
// 		return
// 	}

// 	fmt.Println(user.Phone, user.Password)
// 	//验证用户是否存在（电话是否已经注册）
// 	if model.VerifyPhone(user.Phone) != false {
// 		c.JSON(404, gin.H{
// 			"code":    404,
// 			"message": "用户不存在",
// 		})
// 		return
// 	}
// 	//验证密码,密码可能重复所以还要电话（用这两个验证是否有这条数据在）
// 	if model.VerifyPassword(user.Phone, user.Password) == false {
// 		c.JSON(401, gin.H{
// 			"code":    401,
// 			"message": "密码错误",
// 		})
// 		return
// 	} else {
// 		user0, _ := model.GetUserId(user.Phone)
// 		c.JSON(200, gin.H{
// 			"code":    200,
// 			"message": "登陆成功,请将token放到请求头中",
// 			"token":   model.GenerateToken(user0.UserId),
// 		})
// 		token := model.GenerateToken(user0.UserId)
// 		fmt.Println(token)
// 		return
// 	}

// }

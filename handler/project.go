//需要在路径里输入team名
package handler

import (
	"2022-TEAM-BACKEND/model"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// @Summary "创建项目"
// @Description "创建一个新的项目"
// @Tags project
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param team path string true "team"
// @Success 200
// @Failure 404 "身份验证失败"
// @Failure 400 "格式错误"
// @Failure 401 "创建失败"
// @Router /create_project

func CreateProject(c *gin.Context) {
	var project model.Project
	token := c.Request.Header.Get("token")
	phone, err0 := model.VerifyToken(token)
	if err0 != nil {
		c.JSON(404, gin.H{"message": "身份验证失败"})
		return
	}
	err2 := c.BindJSON(&project)
	if err2 != nil {
		c.JSON(401, gin.H{"message": "格式错误"})
	}
	user, err3 := model.GetUserInfo(phone)
	if err3 != nil {
		log.Println(err3)
	}
	project.Creator = user.NickName
	team := c.Param("team")
	project.TeamName = team
	fmt.Println(project.TeamName)
	// if _, a := model.IfExistTeamname(teamInfo.TeamName); a != 1 {
	// 	c.JSON(200, gin.H{
	// 		"message": "对不起，该团队名已被注册",
	// 	})
	// 	return
	// }
	if err := model.DB.Table("project").Create(&project).Error; err != nil {
		fmt.Println("项目创建出错" + err.Error()) //err.Error打印错误
		return
	}
	c.JSON(200, gin.H{"message": "项目创建成功"})
}

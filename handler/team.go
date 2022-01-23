package handler

import (
	"2022-TEAM-BACKEND/model"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// @Summary "创建团队"
// @Description "创建一个新的团队"
// @Tags team
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param teamInfo body model.Team true "teamInfo"
// @Success 200
// @Failure 404 "身份验证失败"
// @Failure 400 "格式错误"
// @Failure 401 "创建失败"
// @Router /create_team
func CreateTeam(c *gin.Context) {
	var teamInfo model.Team
	token := c.Request.Header.Get("token")
	phone, err0 := model.VerifyToken(token)
	if err0 != nil {
		c.JSON(404, gin.H{"message": "身份验证失败"})
		return
	}
	err2 := c.BindJSON(&teamInfo)
	if err2 != nil {
		c.JSON(401, gin.H{"message": "格式错误"})
		return
	}
	user, err3 := model.GetUserInfo(phone)
	if err3 != nil {
		log.Println(err3)
	}
	if _, a := model.IfExistTeamname(teamInfo.TeamName); a != 1 {
		c.JSON(200, gin.H{
			"message": "对不起，该团队名已被注册",
		})
		return
	}
	if err := model.RegisterTeam(teamInfo.TeamName, teamInfo.Avatar, user.NickName, teamInfo.TeamCoding); err != nil {
		c.JSON(401, gin.H{"message": "创建失败"})
		return
	}
	c.JSON(200, gin.H{"message": "团队创建成功"})
}

// @Summary "加入团队"
// @Description "加入一个新的团队"
// @Tags team
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param team body model.UserTeam true "team"
// @Success 200 "加入成功"
// @Failure 404 "身份验证失败"
// @Failure 400 "格式错误"
// @Failure 401 "加入失败"
// @Router /create_team
func JoinTeam(c *gin.Context) {
	//获取用户信息
	var team model.UserTeam
	token := c.Request.Header.Get("token")
	phone, err0 := model.VerifyToken(token)
	if err0 != nil {
		c.JSON(404, gin.H{"message": "身份验证失败"})
		return
	}
	user, err1 := model.GetUserInfo(phone)
	if err1 != nil {
		log.Println(err1)
	}
	err2 := c.BindJSON(&team)
	if err2 != nil {
		c.JSON(401, gin.H{"message": "格式错误"})
		return
	}
	fmt.Println(team.TeamName)
	if err := model.JoinTeam(user.NickName, team.TeamName); err != nil {
		c.JSON(401, gin.H{"message": "加入失败"})
		return
	}
	c.JSON(200, gin.H{"message": "加入成功"})
}

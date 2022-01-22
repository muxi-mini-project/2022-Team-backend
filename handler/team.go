package handler

import (
	"2022-TEAM-BACKEND/model"
	"log"

	"github.com/gin-gonic/gin"
)

// @Summary "创建团队"
// @Description "创建一个新的团队"
// @Tags team
// @Accept json
// @Produce json
// @Param
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
	}
	user, err3 := model.GetUserInfo(phone)
	if err3 != nil {
		log.Println(err3)
	}
	if err := model.RegisterTeam(teamInfo.TeamName, teamInfo.Avatar, user.NickName, teamInfo.TeamCoding); err != nil {
		c.JSON(401, gin.H{"message": "创建失败"})
		return
	}
	return
}

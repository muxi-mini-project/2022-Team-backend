package handler

import (
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"team/model"
	"team/services"
	"team/services/connector"

	"github.com/gin-gonic/gin"
)

// @Summary "加入团队"
// @Description "加入一个新的团队"
// @Tags team
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param team body model.UserTeam true "team"
// @Success 200 "加入成功"
// @Failure 401 "身份验证失败"
// @Failure 400 "加入失败"
// @Router /team/paticipation [post]
func JoinTeam(c *gin.Context) {
	//扫码加团队：通过扫码我应该可以获得对应的团队id
	//获取用户信息
	var team model.UserTeam
	temp, ok := c.Get("id")
	id := temp.(int)
	if !ok {
		c.JSON(401, gin.H{
			"code":    200,
			"message": "身份验证失败",
		})
	}
	user, err1 := model.GetUserInfo(id)
	if err1 != nil {
		log.Println(err1)
	}
	err2 := c.BindJSON(&team)
	if err2 != nil {
		c.JSON(400, gin.H{
			"401":     400,
			"message": "格式错误",
		})
		return
	}
	fmt.Println(team.TeamId)
	if err := model.JoinTeam(user.UserId, team.TeamId); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "加入失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "加入成功",
		"team_id": team.TeamId,
	})
}

// @Summary "创建团队"
// @Description "创建一个新的团队"
// @Tags team
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param teamInfo body model.Team true "teamInfo"
// @Success 200 "团队创建成功"
// @Failure 401 "身份验证失败"
// @Failure 400 "创建失败"
// @Router /team [post]
func CreateTeam(c *gin.Context) {
	var teamInfo model.Team

	temp, ok := c.Get("id")
	id := temp.(int)
	if !ok {
		c.JSON(401, gin.H{
			"code":    401,
			"message": "身份验证失败",
		})
	}

	file, err := c.FormFile("file")
	teamInfo.TeamName = c.Request.FormValue("name")
	teamInfo.TeamCoding = c.Request.FormValue("coding")

	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
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

	file.Filename = id1 + teamInfo.TeamName + fileExt

	filename := filepath + file.Filename

	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.JSON(400, gin.H{
			"msg": "上传失败!",
		})
		return
	}

	// // 删除原头像
	// user, _ := model.GetUserInfo(id)
	// if user.Path != "" && user.Sha != "" {
	// 	connector.RepoCreate().Del(user.Path, user.Sha)
	// }

	// 上传新头像
	Base64 := services.ImagesToBase64(filename)
	picUrl, picPath, picSha := connector.RepoCreate().Push(file.Filename, Base64)

	os.Remove(filename)

	var avatar model.Team
	avatar.Avatar = picUrl

	team_id := model.RegisterTeam(teamInfo.TeamName, id, teamInfo.TeamCoding, avatar.Avatar)

	if picUrl == "" {
		c.JSON(400, gin.H{
			"message": "上传失败",
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": "团队创建成功",
		"url":     picUrl,
		"sha":     picSha,
		"path":    picPath,
		"team_id": team_id,
	})

}

// @Summary "查看团队"
// @Description "单击团队名查看团队信息"
// @Tags team
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param team_id path string true "team_id"
// @Success 200
// @Failure 401 "身份验证失败"
// @Failure 400 "获取失败"
// @Router /team/:team_id [get]
func ViewTeamInfo(c *gin.Context) {
	_, ok := c.Get("id")
	if !ok {
		c.JSON(401, gin.H{
			"code":    401,
			"message": "身份验证失败",
		})
	}

	team_id := c.Param("team_id")

	team, userInfo, err := model.GetTeamInfo(team_id)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "获取失败",
		})
	}

	c.JSON(200, gin.H{
		"code": 200,
		"team": team,
		"user": userInfo,
	})
}

// @Summary "查看用户加入的所有团队"
// @Description "刚登陆后的第一个界面"
// @Tags team
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200
// @Failure 401 "身份验证失败"
// @Failure 400 "获取失败"
// @Router /team [get]
func ViewAllTeam(c *gin.Context) {
	temp, ok := c.Get("id")
	id := temp.(int)
	if !ok {
		c.JSON(401, gin.H{
			"code":    401,
			"message": "验证失败",
		})
	}

	team, err := model.GetAllTeamInfo(id)
	if err != nil {
		c.JSON(404, gin.H{
			"code":    400,
			"message": "获取失败",
		})
	}

	c.JSON(200, gin.H{"team": team})
}

// // @Summary "创建团队"
// // @Description "创建一个新的团队"
// // @Tags team
// // @Accept json
// // @Produce json
// // @Param token header string true "token"
// // @Param teamInfo body model.Team true "teamInfo"
// // @Success 200
// // @Failure 401 "身份验证失败"
// // @Failure 404 "格式错误"
// // @Failure 400 "创建失败"
// // @Router /create_team
// //团队码中可以存入团队id
// func CreateTeam(c *gin.Context) {
// 	var teamInfo model.Team
// 	temp, ok := c.Get("id")
// 	id := temp.(int)
// 	if !ok {
// 		c.JSON(401, gin.H{"message": "验证失败"})
// 	}
// 	err2 := c.BindJSON(&teamInfo)
// 	if err2 != nil {
// 		c.JSON(401, gin.H{"message": "格式错误"})
// 		return
// 	}
// 	user, err3 := model.GetUserInfo(id)
// 	if err3 != nil {
// 		log.Println(err3)
// 	}
// 	team_id := model.RegisterTeam(teamInfo.TeamName, user.UserId, teamInfo.TeamCoding)
// 	c.JSON(200, gin.H{
// 		"message": "团队创建成功",
// 		"team_id": team_id,
// 	})
// }

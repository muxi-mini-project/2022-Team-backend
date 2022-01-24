package handler

//需要在路径里输入team名
import (
	"2022-TEAM-BACKEND/model"
	"fmt"

	// "log"

	"github.com/gin-gonic/gin"
)

// @Summary "创建任务并分配"
// @Description "项目任务新建"
// @Tags project
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param team path string true "team"
// @Success 200
// @Failure 400 "格式错误"
// @Failure 401 "创建失败"
// @Router //create_task/step [post]

func CreateTask(c *gin.Context) {
	var task model.Task
	// token := c.Request.Header.Get("token")
	// phone, err0 := model.VerifyToken(token)
	// if err0 != nil {
	// 	c.JSON(404, gin.H{"message": "身份验证失败"})
	// 	return
	// }
	err2 := c.BindJSON(&task)
	if err2 != nil {
		c.JSON(401, gin.H{"message": "格式错误"})
	}
	// user, err3 := model.GetUserInfo(phone)
	// if err3 != nil {
	// 	log.Println(err3)
	// }
	// project.Creator = user.NickName
	// team := c.Param("team")
	// project.TeamName = team
	// fmt.Println(project.TeamName)
	// if _, a := model.IfExistTeamname(teamInfo.TeamName); a != 1 {
	// 	c.JSON(200, gin.H{
	// 		"message": "对不起，该团队名已被注册",
	// 	})
	// 	return
	// }
	if err := model.DB.Table("task").Create(&task).Error; err != nil {
		fmt.Println("任务创建出错" + err.Error()) //err.Error打印错误
		return
	}
	c.JSON(200, gin.H{"message": "任务创建成功"})
}

// @Summary "团队成员"
// @Description "通过团队的id获得该团队成员昵称"
// @Tags team
// @Accept json
// @Produce json
// @Param toekn header string true "token"
// @Param Pid path string true "Pid"
// @Success 200 {object} model.BooksInfo"{"msg":"success"}"
// @Failure 404 "获取失败"
// @Router /homepage/shelf [get]
//显示团队成员昵称
func AssignTasks(c *gin.Context) {
	id := c.Param("Pid")
	MemberId := model.GetTeamMenberId(id)
	MemberInfo, err := model.GetTeamMenberName(MemberId)
	if err != nil {
		c.JSON(404, gin.H{"message": "获取失败"})
	}
	c.JSON(200, MemberInfo)
}

//分配任务

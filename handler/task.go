package handler

//需要在路径里输入team名
import (
	"2022-TEAM-BACKEND/model"
	"fmt"
	"strconv"

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

	err2 := c.BindJSON(&task)
	if err2 != nil {
		c.JSON(401, gin.H{"message": "格式错误"})
	}

	if err := model.DB.Table("task").Create(&task).Error; err != nil {
		fmt.Println("任务创建出错" + err.Error()) //err.Error打印错误
		return
	}
	c.JSON(200, gin.H{
		"message": "任务创建成功",
		"task_id": task.TaskId,
	})
}

// @Summary "团队成员"
// @Description "通过团队的id获得该团队成员昵称"
// @Tags team
// @Accept json
// @Produce json
// @Param toekn header string true "token"
// @Param Pid path string true "Pid"
// @Param TaskId path string true "TaskId"
// @Success 200
// @Failure 404 "获取失败"
// @Router /Assign_task/:Pid [get]
// @Router /Assign_task/:Pid/:TaskId [post]
//这个函数得改一下
//显示团队成员昵称(打对钩选人那个)
func AssignTasks(c *gin.Context) {
	Pid := c.Param("Pid")
	MemberId := model.GetTeamMenberId(Pid)
	MemberInfo, err := model.GetTeamMenberName(MemberId)
	if err != nil {
		c.JSON(404, gin.H{"message": "获取失败"})
		return
	}
	c.JSON(200, MemberInfo)
	//获取对应姓名的id
	var name []string
	c.BindJSON(&name)
	var AssignId []string
	for _, Name := range name {
		for id, MInfo := range MemberInfo {
			if Name == MInfo {
				AssignId = append(AssignId, string(MemberId[id]))
			}
		}
	}
	//分配任务
	temp := c.Param("TaskId")
	Tid, _ := strconv.Atoi(temp)
	for _, AId := range AssignId {
		temp2, _ := strconv.Atoi(AId)
		err := model.AssginIntoTable(temp2, Tid, false)
		if err != nil {
			fmt.Println(err)
		}
	}
}

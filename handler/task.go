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
// @Failure 404 "格式错误"
// @Failure 400 "创建失败"
// @Router //create_task/step [post]

func CreateTask(c *gin.Context) {
	token := c.Request.Header.Get("token")
	phone, err0 := model.VerifyToken(token)
	if err0 != nil {
		c.JSON(401, gin.H{"message": "身份验证失败"})
		return
	}
	userInfo, err := model.GetUserInfo(phone)
	if err != nil {
		c.JSON(404, gin.H{"message": "获取失败"})
	}
	var task model.Task
	task.CreatorId = userInfo.UserId
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
//局部变量用小驼峰
func AssignTasks(c *gin.Context) {
	pId := c.Param("Pid")
	memberId := model.GetTeamMenberId(pId)
	memberInfo, err := model.GetTeamMenberName(memberId)
	if err != nil {
		c.JSON(404, gin.H{"message": "获取失败"})
		return
	}
	c.JSON(200, memberInfo)
	//获取对应姓名的id
	var names []string
	c.BindJSON(&names)
	var assignId []string
	for _, name := range names {
		for id, MInfo := range memberInfo {
			if name == MInfo {
				assignId = append(assignId, string(memberId[id]))
			}
		}
	}
	//分配任务
	temp := c.Param("TaskId")
	tId, _ := strconv.Atoi(temp)
	for _, aId := range assignId {
		temp2, _ := strconv.Atoi(aId)
		err := model.AssginIntoTable(temp2, tId, false)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// @Summary "删除任务"
// @Description "删除一个任务"
// @Tags project
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param task_id path string true "task_id"
// @Success 200
// @Failure 401 "身份验证失败"
// @Failure 404 "获取失败"
// @Failure 403 "无权删除"
// @Failure 400 "删除失败"
// @Router /delete_task/:task_id [put]
func DeleteTask(c *gin.Context) {
	token := c.Request.Header.Get("token")
	phone, err0 := model.VerifyToken(token)
	if err0 != nil {
		c.JSON(401, gin.H{"message": "身份验证失败"})
		return
	}
	userInfo, err := model.GetUserInfo(phone)
	if err != nil {
		c.JSON(404, gin.H{"message": "获取失败"})
	}
	taskId := c.Param("task_id")
	// pId,_ := strconv.Atoi(temp)
	task, _ := model.GetTaskInfo(taskId)
	if userInfo.UserId != task.CreatorId {
		c.JSON(403, gin.H{"message": "对不起，非创建人无法删除任务"})
	} else {
		if err := model.RemoveTask(taskId); err != nil {
			c.JSON(400, gin.H{"message": "删除失败"})
			return
		}
		c.JSON(200, gin.H{"message": "删除成功"})
	}

}

// @Summary "修改任务"
// @Description "创建人修改一个任务"
// @Tags project
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param task_id path string true "task_id"
// @Success 200
// @Failure 401 "身份验证失败"
// @Failure 404 "获取失败"
// @Failure 403 "无权删除"
// @Failure 400 "删除失败"
// @Router /modify_project/:pId [put]
//分配任务就按之前分配任务的那个接口去搞
func ModifyTask(c *gin.Context) {
	// var project model.Project
	token := c.Request.Header.Get("token")
	phone, err0 := model.VerifyToken(token)
	if err0 != nil {
		c.JSON(401, gin.H{"message": "身份验证失败"})
		return
	}
	userInfo, _ := model.GetUserInfo(phone)
	taskId := c.Param("task_id")
	// pId,_ := strconv.Atoi(temp)
	task, _ := model.GetTaskInfo(taskId)
	if userInfo.UserId != task.CreatorId {
		c.JSON(403, gin.H{"message": "对不起，非创建人无法修改任务"})
	} else {
		err := c.BindJSON(task)
		if err != nil {
			fmt.Println(err)
		}
		if err := model.ChangeTaskInfo(task); err != nil {
			c.JSON(400, gin.H{"message": "修改失败"})
			return
		}
		c.JSON(200, gin.H{"message": "修改成功"})

	}
}

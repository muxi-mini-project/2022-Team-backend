package handler

//需要在路径里输入team名
import (
	"fmt"
	"strconv"
	"team/model"

	"github.com/gin-gonic/gin"
)

// @Summary "获取任务编辑页面需要的信息"
// @Description "获取项目任务新建所需的团队成员和项目名称"
// @Tags task
// @Accept json
// @Produce json
// @Param team_id path string true "team_id"
// @Success 200	"获取成功"
// @Failure 400 "获取失败"
// @Router /task/team_info/:team_id [get]
//局部变量用小驼峰
func GetChoice(c *gin.Context) {
	team_id := c.Param("team_id")
	memberId := model.GetTeamMenberId(team_id)
	memberInfo, err := model.GetTeamMenberName(memberId)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    401,
			"message": "获取失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    memberInfo,
	})

	teamPro, err1 := model.GetTeamPro(team_id)
	if err1 != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "获取失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    teamPro,
	})
}

// @Summary "项目对应的步骤"
// @Description "在新建任务界面选择项目后(填入项目名称)返回步骤"
// @Tags task
// @Accept json
// @Produce json
// @Param pro body model.Project true "项目名称"
// @Param team_id path string true "team_id"
// @Success 200
// @Failure 400 "获取失败"
// @Router /task/pro_step_info/:team_id [get]
func GetProStepName(c *gin.Context) {
	//局部变量用小驼峰
	var pro model.Project
	if err := c.BindJSON(&pro); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "输入有误，格式错误",
		})
		return
	}
	temp := c.Param("team_id")
	tId, _ := strconv.Atoi(temp)
	name, err := model.GetProStep(pro.ProjectName, tId)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "获取成功",
		"user_id": name,
	})
}

// @Summary "创建任务并分配"
// @Description "项目任务新建"
// @Tags task
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param step_id path string true "team_id"
// @Param task body model.Task true "注意填入的成员是结构体"
// @Success 200
// @Failure 401 "身份验证失败"
// @Failure 404 "格式错误"
// @Failure 400 "创建失败"
// @Router /team/task/:team_id [post]
func CreateTask(c *gin.Context) {
	id := c.MustGet("id").(int)
	userInfo, err := model.GetUserInfo(id)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "获取失败",
		})
	}
	var task model.Task
	task.CreatorId = userInfo.UserId
	err2 := c.BindJSON(&task)
	if err2 != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "输入格式错误",
		})
	}

	temp2 := c.Param("team_id")
	tId, _ := strconv.Atoi(temp2)
	// fmt.Println(task.ProName, tId)
	task.StepId, err = model.GetProStepId(tId, task.ProName, task.StepName)
	fmt.Println(task.StepId)
	if err != nil {
		fmt.Println(err, "hen")
	}
	task.TeamId = tId
	if task.TaskId = model.CreateTask(task.TaskName, task.CreatorId, task.StartTime, task.Deadline, task.Remark, task.StepId, task.ProName, task.StepName, task.TeamId); task.TaskId == 0 {
		fmt.Println(task.TaskId, "cao")
		c.JSON(400, gin.H{
			"code":    400,
			"message": "任务创建出错",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "任务创建成功",
		"task_id": task.TaskId,
	})
	//分配任务
	team_id := c.Param("team_id")
	memberId := model.GetTeamMenberId(team_id)
	memberInfo, err := model.GetTeamMenberName(memberId)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "获取信息失败",
		})
		return
	}
	// c.JSON(200, memberInfo)
	//获取对应姓名的id
	var assignId []string
	for _, name := range task.Member {
		for id, MInfo := range memberInfo {
			if name == MInfo {
				assignId = append(assignId, string(memberId[id]))
			}
		}
	}
	for _, aId := range assignId {
		temp2, _ := strconv.Atoi(aId)
		u, err0 := model.GetUserInfo(temp2)
		err := model.AssginIntoTable(temp2, task.TaskId, u.NickName, false)
		if err != nil || err0 != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"message": "分配任务失败",
			})
			return
		}
	}

}

// @Summary "删除任务"
// @Description "删除一个任务"
// @Tags task
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param task_id path string true "task_id"
// @Success 200 "删除成功"
// @Failure 401 "身份验证失败"
// @Failure 400 "删除失败"
// @Router /task/:task_id [delete]
func DeleteTask(c *gin.Context) {
	id := c.MustGet("id").(int)
	userInfo, err := model.GetUserInfo(id)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "获取失败",
		})
	}
	taskId := c.Param("task_id")
	// project_id,_ := strconv.Atoi(temp)
	task, _ := model.GetTaskInfo(taskId)
	if userInfo.UserId != task.CreatorId {
		c.JSON(401, gin.H{
			"code":    401,
			"message": "对不起，非创建人无法删除任务",
		})
	} else {
		if err := model.RemoveTask(taskId); err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"message": "删除失败",
			})
			return
		}

		c.JSON(200, gin.H{
			"code":    "200",
			"message": "删除成功",
		})

	}

}

// @Summary "查看任务"
// @Description "编辑前查看任务信息"
// @Tags project
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param task_id path string true "task_id"
// @Success 200
// @Failure 401 "身份验证失败"
// @Failure 400 "获取失败"
// @Router /team/task/:task_id [get]
func ViewTask(c *gin.Context) {

	task_id := c.Param("task_id")

	task, err := model.GetTaskInfo(task_id)

	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "获取失败",
		})
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    task,
	})

}

// @Summary "修改任务"
// @Description "创建人修改一个任务"
// @Tags project
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param task_id path string true "task_id"
// @Param task body model.Task true "填入需要修改的信息"
// @Success 200	"修改成功"
// @Failure 401 "身份验证失败"
// @Failure 400 "删除失败"
// @Router /team/task/:task_id/:team_id [put]
func ModifyTask(c *gin.Context) {
	// var project model.Project
	id := c.MustGet("id").(int)
	userInfo, _ := model.GetUserInfo(id)
	taskId := c.Param("task_id")
	// project_id,_ := strconv.Atoi(temp)
	task, _ := model.GetTaskInfo(taskId)
	if userInfo.UserId != task.CreatorId {
		c.JSON(401, gin.H{
			"code":    400,
			"message": "对不起，非创建人无法修改任务",
		})
	} else {
		err := c.BindJSON(&task)
		if err != nil {
			fmt.Println(err)
		}
		if err := model.ChangeTaskInfo(task); err != nil {
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
	fmt.Println(task.Member)
	//分配任务
	team_id := c.Param("team_id")
	memberId := model.GetTeamMenberId(team_id)
	memberInfo, err := model.GetTeamMenberName(memberId)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "获取信息失败",
		})
		return
	}
	// c.JSON(200, memberInfo)
	//获取对应姓名的id
	var assignId []string
	for _, name := range task.Member {
		for id, MInfo := range memberInfo {
			if name == MInfo {
				assignId = append(assignId, string(memberId[id]))
			}
		}
	}
	for _, aId := range assignId {
		temp2, _ := strconv.Atoi(aId)
		u, err0 := model.GetUserInfo(temp2)
		err := model.AssginIntoTable(temp2, task.TaskId, u.NickName, false)
		if err != nil || err0 != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"message": "分配任务失败",
			})
			return
		}
	}

}

// //局部变量用小驼峰
// func AssignTasks(c *gin.Context) {
// 	project_id := c.Param("project_id")
// 	memberId := model.GetTeamMenberId(project_id)
// 	memberInfo, err := model.GetTeamMenberName(memberId)
// 	if err != nil {
// 		c.JSON(404, gin.H{"message": "获取失败"})
// 		return
// 	}
// 	c.JSON(200, memberInfo)
// 	//获取对应姓名的id
// 	var names []string
// 	c.BindJSON(&names)
// 	var assignId []string
// 	for _, name := range names {
// 		for id, MInfo := range memberInfo {
// 			if name == MInfo {
// 				assignId = append(assignId, string(memberId[id]))
// 			}
// 		}
// 	}
// 	//分配任务
// 	temp := c.Param("task_id")
// 	tId, _ := strconv.Atoi(temp)
// 	for _, aId := range assignId {
// 		temp2, _ := strconv.Atoi(aId)
// 		err := model.AssginIntoTable(temp2, tId, false)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 	}
// }

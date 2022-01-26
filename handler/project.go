//需要在路径里输入team名
package handler

import (
	"fmt"
	"log"
	"strconv"
	"team/model"

	"github.com/gin-gonic/gin"
)

// @Summary "创建项目"
// @Description "创建一个新的项目"
// @Tags project
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param team_id path string true "team_id"
// @Success 200
// @Failure 401 "身份验证失败"
// @Failure 404 "格式错误"
// @Failure 400 "创建失败"
// @Router /create_project

func CreateProject(c *gin.Context) {
	var project model.Project
	// token := c.Request.Header.Get("token")
	// id, err0 := model.VerifyToken(token)
	Id, ok := c.Get("user_id")
	id := Id.(int)
	if !ok {
		c.JSON(404, gin.H{"message": "身份验证失败"})
		return
	}
	err2 := c.BindJSON(&project)
	if err2 != nil {
		c.JSON(401, gin.H{"message": "格式错误"})
	}
	user, err3 := model.GetUserInfo(id)
	if err3 != nil {
		log.Println(err3)
	}
	project.CreatorId = user.UserId
	temp := c.Param("team")
	teamId, _ := strconv.Atoi(temp)
	project.TeamId = teamId
	fmt.Println(project.TeamId)
	// if _, a := model.IfExistTeamname(teamInfo.TeamName); a != 1 {
	// 	c.JSON(200, gin.H{
	// 		"message": "对不起，该项目名已被使用",
	// 	})
	// 	return
	// }
	if err := model.DB.Table("project").Create(&project).Error; err != nil {
		fmt.Println("项目创建出错" + err.Error()) //err.Error打印错误
		return
	}
	c.JSON(200, gin.H{"message": "项目创建成功"})
}

//点一个新建执行一次
// @Summary "新建步骤"
// @Description "单击新建创建步骤"
// @Tags step
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param project_id path string true "project_id"
// @Success 200
// @Failure 401 "身份验证失败"
// @Failure 404 "格式错误"
// @Failure 400 "创建失败"
// @Router /create_step/:project_id
func CreateStep(c *gin.Context) {
	var step model.Step
	err := c.BindJSON(&step)
	if err != nil {
		c.JSON(400, gin.H{"message": "格式错误"})
		return
	}
	//if里声明的变量是局部变量
	step.ProjectId, _ = strconv.Atoi(c.Param("project_id"))
	if err := model.AddStep(step.StepName, step.ProjectId); err != nil {
		c.JSON(401, gin.H{"message": "新建步骤失败"})
		return
	}
	c.JSON(200, gin.H{
		"message": "步骤创建成功",
		"step_id": step.StepId,
	})
}

// @Summary "删除项目"
// @Description "删除一个项目"
// @Tags project
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param project_id path string true "project_id"
// @Success 200
// @Failure 401 "身份验证失败"
// @Failure 404 "获取失败"
// @Failure 403 "无权删除"
// @Failure 400 "删除失败"
// @Router /delete_project/:project_id [put]
func DeleteProject(c *gin.Context) {
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
	project_id := c.Param("project_id")
	// project_id,_ := strconv.Atoi(temp)
	project, _ := model.GetProjectInfo(project_id)
	if userInfo.UserId != project.CreatorId {
		c.JSON(403, gin.H{"message": "对不起，非创建人无法删除项目"})
	} else {
		if err := model.RemoveProject(project_id); err != nil {
			c.JSON(400, gin.H{"message": "删除失败"})
			return
		}
		c.JSON(200, gin.H{"message": "删除成功"})
	}

}

// @Summary "修改项目"
// @Description "创建人修改一个项目"
// @Tags project
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param project_id path string true "project_id"
// @Success 200
// @Failure 401 "身份验证失败"
// @Failure 404 "获取失败"
// @Failure 403 "无权删除"
// @Failure 400 "删除失败"
// @Router /modify_project/:project_id [put]
func ModifyProject(c *gin.Context) {
	// var project model.Project
	token := c.Request.Header.Get("token")
	phone, err0 := model.VerifyToken(token)
	if err0 != nil {
		c.JSON(401, gin.H{"message": "身份验证失败"})
		return
	}
	userInfo, _ := model.GetUserInfo(phone)
	project_id := c.Param("project_id")
	// project_id,_ := strconv.Atoi(temp)
	project2, _ := model.GetProjectInfo(project_id)
	if userInfo.UserId != project2.CreatorId {
		c.JSON(403, gin.H{"message": "对不起，非创建人无法修改项目"})
	} else {
		err := c.BindJSON(project2)
		if err != nil {
			fmt.Println(err)
		}
		if err := model.ChangeProjectInfo(project2); err != nil {
			c.JSON(400, gin.H{"message": "修改失败"})
			return
		}
		c.JSON(200, gin.H{"message": "修改成功"})

	}
}

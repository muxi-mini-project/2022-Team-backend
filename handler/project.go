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
// @Success 200 "项目创建成功"
// @Failure 401 "身份验证失败"
// @Failure 400 "创建失败"
// @Router /team/project/:team_id [post]

func CreateProject(c *gin.Context) {
	var project model.Project
	Id, ok := c.Get("id")
	id := Id.(int)
	if !ok {
		c.JSON(401, gin.H{
			"code":    401,
			"message": "身份验证失败",
		})
		return
	}

	err2 := c.BindJSON(&project)
	if err2 != nil {
		c.JSON(400, gin.H{
			"code":    "400",
			"message": "输入格式错误",
		})
	}
	fmt.Println(project)
	user, err3 := model.GetUserInfo(id)
	if err3 != nil {
		log.Println(err3)
	}

	project.CreatorId = user.UserId
	temp := c.Param("team_id")
	teamId, _ := strconv.Atoi(temp)
	project.TeamId = teamId
	fmt.Println(project.TeamId)

	if project.ProjectId = model.CreatePro(project.ProjectName, project.CreatorId, project.StartTime, project.Deadline, project.Remark, project.TeamId); project.ProjectId == 0 {
		fmt.Println(project.ProjectId, "cao")
		c.JSON(400, gin.H{
			"code":    400,
			"message": "项目创建出错",
		})
		return
	}

	var step string
	for _, step = range project.Step {
		if err := model.AddStep(step, project.ProjectId); err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"message": "新建步骤失败",
			})
			return
		}
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "项目创建成功",
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
// @Failure 400 "删除失败"
// @Router /team/project/:project_id [delete]
func DeleteProject(c *gin.Context) {
	temp, ok := c.Get("id")
	id := temp.(int)
	if !ok {
		c.JSON(401, gin.H{
			"code":    401,
			"message": "身份验证失败",
		})
	}
	userInfo, err := model.GetUserInfo(id)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "获取失败",
		})
	}
	project_id := c.Param("project_id")
	// project_id,_ := strconv.Atoi(temp)
	project, _ := model.GetProjectInfo(project_id)
	if userInfo.UserId != project.CreatorId {
		c.JSON(401, gin.H{
			"code":    401,
			"message": "对不起，非创建人无法删除项目",
		})
	} else {
		if err := model.RemoveProject(project_id); err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"message": "删除失败",
			})
			return
		}
		c.JSON(200, gin.H{
			"code":    200,
			"message": "删除成功",
		})
	}

}

// @Summary "查看项目"
// @Description "单击编辑查看项目信息"
// @Tags project
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param project_id path string true "project_id"
// @Success 200 "获取成功"
// @Failure 401 "身份验证失败"
// @Failure 404 "获取失败"
// @Router /team/project/:project_id [get]
func ViewProject(c *gin.Context) {
	_, ok := c.Get("id")
	if !ok {
		c.JSON(401, gin.H{
			"code":    200,
			"message": "身份验证失败",
		})
	}

	project_id := c.Param("project_id")

	project, err := model.GetProjectInfo(project_id)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "获取失败",
		})
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    project,
	})

}

// @Summary "修改项目"
// @Description "创建人修改一个项目"
// @Tags project
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param project_id path string true "project_id"
// @Success 200 "修改成功"
// @Failure 401 "身份验证失败"
// @Failure 400 "删除失败"
// @Router /team/project/:project_id [put]
func ModifyProject(c *gin.Context) {
	temp, ok := c.Get("id")
	id := temp.(int)
	if !ok {
		c.JSON(401, gin.H{
			"code":    "401",
			"message": "身份验证失败",
		})
	}

	userInfo, _ := model.GetUserInfo(id)
	project_id := c.Param("project_id")

	project1, _ := model.GetProjectInfo(project_id)
	var project model.Project
	if userInfo.UserId != project1.CreatorId {
		c.JSON(401, gin.H{
			"code":    401,
			"message": "对不起，非创建人无法修改项目",
		})
	} else {
		//注意bindjson这里有“&”符号
		err := c.BindJSON(&project)
		if err != nil {
			fmt.Println(err, "hhh")
		}
		if err := model.ChangeProjectInfo(project, project1.ProjectId); err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"message": "修改失败",
			})
			return
		}

		var step string
		for _, step = range project.Step {
			if err := model.AddStep(step, project1.ProjectId); err != nil {
				c.JSON(400, gin.H{
					"code":    400,
					"message": "新建步骤失败",
				})
				return
			}
		}

		c.JSON(200, gin.H{
			"code":    200,
			"message": "修改成功",
		})

	}
}

// //点一个新建执行一次(这一版本先留存)
// // @Summary "新建步骤"
// // @Description "单击新建创建步骤"
// // @Tags step
// // @Accept json
// // @Produce json
// // @Param token header string true "token"
// // @Param project_id path string true "project_id"
// // @Success 200
// // @Failure 401 "身份验证失败"
// // @Failure 404 "格式错误"
// // @Failure 400 "创建失败"
// // @Router /team/create_step/:project_id
// func CreateStep(c *gin.Context) {
// 	var step model.Step
// 	err := c.BindJSON(&step)
// 	if err != nil {
// 		c.JSON(400, gin.H{"message": "格式错误"})
// 		return
// 	}
// 	//if里声明的变量是局部变量
// 	step.ProjectId, _ = strconv.Atoi(c.Param("project_id"))
// 	if err := model.AddStep(step.StepName, step.ProjectId); err != nil {
// 		c.JSON(401, gin.H{"message": "新建步骤失败"})
// 		return
// 	}
// 	c.JSON(200, gin.H{
// 		"message": "步骤创建成功",
// 		"step_id": step.StepId,
// 	})
// }

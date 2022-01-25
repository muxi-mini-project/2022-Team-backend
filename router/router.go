package router

import (
	"2022-TEAM-BACKEND/handler"

	"github.com/gin-gonic/gin"
)

//POST新建，put更新
func Router(r *gin.Engine) {
	r.POST("/user", handler.User)
	r.POST("/user/pupup", handler.InitUserInfo)
	r.POST("/login", handler.Login)
	r.GET("/info", handler.Userinfo)
	r.PUT("/info", handler.ChangeInfomation)
	r.GET("/change_password/verify", handler.VerifyPassword)
	r.POST("/change_password/change", handler.ChangePassword)
	r.POST("/create_team", handler.CreateTeam)
	r.POST("/join_team", handler.JoinTeam)
	r.POST("/create_project/:team_id", handler.CreateProject)
	r.POST("/create_project/create_step/:Pid", handler.CreateStep)
	r.POST("/create_task/:Pid", handler.CreateTask)
	r.GET("/Assign_task/:Pid", handler.AssignTasks)
	r.POST("/Assign_task/:Pid/:TaskId", handler.AssignTasks)
	r.GET("/info/todolist", handler.ToDoList)
	r.GET("/info/donelist", handler.DoneList)
	r.PUT("/info/donelist", handler.DoneTask)
}

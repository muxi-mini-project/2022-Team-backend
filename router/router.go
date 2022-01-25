package router

import (
	"2022-TEAM-BACKEND/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

//POST新建，put更新
//刚开始写(ง •_•)ง
func Router(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API router.")
	})
	x1 := r.Group("/user")
	{
		x1.POST("", handler.User)
		x1.POST("/pupup", handler.InitUserInfo)
	}
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
	r.POST("/create_project/create_step/:pId", handler.CreateStep)
	r.POST("/create_task/:Pid", handler.CreateTask)
	r.GET("/Assign_task/:Pid", handler.AssignTasks)
	r.POST("/Assign_task/:Pid/:TaskId", handler.AssignTasks)
	r.GET("/info/todolist", handler.ToDoList)
	r.GET("/info/donelist", handler.DoneList)
	r.PUT("/info/donetask", handler.DoneTask)
	r.PUT("/info/cancel_done", handler.CancelDone)
	r.PUT("/delete_project/:pId", handler.DeleteProject)
	r.PUT("/modify_project/:pId", handler.ModifyProject)
	r.PUT("/delete_task/:task_id", handler.DeleteTask)
}

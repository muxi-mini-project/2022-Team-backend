package router

import (
	"net/http"
	"team/handler"

	"github.com/gin-gonic/gin"
)

//POST新建，put更新，delete？
//路由尽量搞小写
func Router(r *gin.Engine) {
	//没路由时的统一回复

	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API router.")
	})
	//注册新用户
	r.POST("", handler.User)

	//注册后弹窗中初始化昵称和头像
	r.POST("/pupup", handler.InitUserInfo)
	r.POST("/login", handler.Login)
	v1 := r.Group("/user")
	{
		//注册新用户
		v1.POST("", handler.User)

		//注册后弹窗中初始化昵称和头像
		v1.POST("/pupup", handler.InitUserInfo)

		//获取用户信息
		v1.GET("/info", handler.Userinfo)

		//修改用户信息
		v1.PUT("/info", handler.ChangeInfomation)

		//验证密码
		v1.GET("/change_password/verify", handler.VerifyPassword)

		//修改密码
		v1.POST("/change_password/change", handler.ChangePassword)

	}

	v2 := r.Group("/team")
	{
		//创建团队
		v2.POST("/create_team", handler.CreateTeam)

		//加入团队
		v2.POST("/join_team", handler.JoinTeam)

		//创建项目
		v2.POST("/create_project/:team_id", handler.CreateProject)

		//删除项目
		v2.DELETE("/delete_project/:project_id", handler.DeleteProject)

		//修改项目
		v2.PUT("/modify_project/:project_id", handler.ModifyProject)

		//创建步骤
		v2.POST("/create_project/create_step/:project_id", handler.CreateStep)

		//创建任务
		v2.POST("/create_task/:step_id", handler.CreateTask)

		//查看成员
		v2.GET("/Assign_task/:project_id", handler.AssignTasks)

		//布置任务
		v2.POST("/Assign_task/:project_id/:task_id", handler.AssignTasks)

		//修改任务
		v2.PUT("/modify_task/:task_id", handler.ModifyTask)

		//删除任务
		v2.DELETE("/delete_task/:task_id", handler.DeleteTask)

	}

	v3 := r.Group("/mytask")
	{

		//查看今日待办
		v3.GET("/info/todolist", handler.ToDoList)

		//查看已完成任务
		v3.GET("/info/donelist", handler.DoneList)

		//在今日待办中单击完成任务
		v3.PUT("/info/donetask", handler.DoneTask)

		//在已完成任务中单击取消完成
		v3.PUT("/info/cancel_done", handler.CancelDone)
	}
}

// r.POST("/user", handler.User)
// r.POST("/user/pupup", handler.InitUserInfo)
// r.POST("/login", handler.Login)
// r.GET("/info", handler.Userinfo)
// r.PUT("/info", handler.ChangeInfomation)
// r.GET("/change_password/verify", handler.VerifyPassword)
// r.POST("/change_password/change", handler.ChangePassword)
// r.POST("/create_team", handler.CreateTeam)
// r.POST("/join_team", handler.JoinTeam)
// r.POST("/create_project/:team_id", handler.CreateProject)
// r.POST("/create_project/create_step/:pId", handler.CreateStep)
// r.POST("/create_task/:Pid", handler.CreateTask)
// r.GET("/Assign_task/:Pid", handler.AssignTasks)
// r.POST("/Assign_task/:Pid/:TaskId", handler.AssignTasks)
// r.PUT("/delete_project/:pId", handler.DeleteProject)
// r.PUT("/modify_project/:pId", handler.ModifyProject)
// r.PUT("/delete_task/:task_id", handler.DeleteTask)
// r.GET("/info/todolist", handler.ToDoList)
// r.GET("/info/donelist", handler.DoneList)
// r.PUT("/info/donetask", handler.DoneTask)
// r.PUT("/info/cancel_done", handler.CancelDone)

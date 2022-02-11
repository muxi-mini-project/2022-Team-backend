package router

import (
	"net/http"
	"team/handler"

	"github.com/gin-gonic/gin"
)

//POST新建，put更新，delete
//路由尽量搞小写/
func Router(r *gin.Engine) {
	//没路由时的统一回复

	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API router.")
	})
	//注册新用户
	r.POST("/user", handler.User) //1

	//登录(完成注册设置时自动登录)
	r.POST("/login", handler.Login) //1

	v1 := r.Group("/user").Use(Auth()) //1
	{

		//注册后弹窗中初始化昵称
		v1.POST("/pupup", handler.InitUserInfo) //1

		//初始化头像
		v1.POST("/pupup/avatar", handler.ModifyProfile) //1

		//获取用户信息
		v1.GET("/info", handler.Userinfo) //1

		//修改用户昵称
		v1.PUT("/info", handler.ChangeNickname) //1

		//修改用户头像
		v1.PUT("/avatar", handler.ModifyProfile) //1

		//验证密码
		v1.POST("/change_password/verify", handler.VerifyPassword) //1

		//修改密码
		v1.POST("/change_password/change", handler.ChangePassword) //1

		//用户反馈
		v1.PUT("/feedback", handler.Feedback) //1

	}

	v2 := r.Group("/team").Use(Auth())
	{
		//查看加入的所有团队
		v2.GET("", handler.ViewAllTeam) //1

		//创建团队
		v2.POST("", handler.CreateTeam) //团队码先放一放

		//加入团队
		v2.POST("/paticipation", handler.JoinTeam) //团队码的问题

		//查看团队信息(creator_id+user_id-->creator)
		v2.GET("/:team_id", handler.ViewTeamInfo) //1

		//创建项目
		v2.POST("/project/:team_id", handler.CreateProject) //1

		//删除项目
		v2.DELETE("/project/:project_id", handler.DeleteProject) //1

		//查看项目信息(先查看信息后修改)
		v2.GET("/project/:project_id", handler.ViewProject) //1

		//修改项目
		v2.PUT("/project/:project_id", handler.ModifyProject) //1

		//创建任务（获取信息--获取步骤--完成新建）
		v2.GET("/task/team_info/:team_id", handler.GetChoice) //1

		//获取所选项目对应步骤
		v2.GET("/task/pro_step_info/:team_id", handler.GetProStepName) //1

		//布置任务
		v2.POST("/task/:team_id", handler.CreateTask) //1

		//查看任务信息(先查看后修改)
		v2.GET("/task/:task_id", handler.ViewTask) //1

		//修改任务
		v2.PUT("/task/:task_id/:team_id", handler.ModifyTask) //1

		//删除任务
		v2.DELETE("/task/:task_id", handler.DeleteTask) //1

	}

	v3 := r.Group("/mytask").Use(Auth())
	{

		//查看今日待办(所有信息一次传上去，完了前端一部分写在列表里，一部分写在具体内容里)
		v3.GET("/info/todolist", handler.ToDoList) //1

		//查看已完成任务
		v3.GET("/info/donelist", handler.DoneList) //1

		//在今日待办中单击完成任务
		v3.PUT("/info/donetask/:task_id", handler.DoneTask) //1

		//在已完成任务中单击取消完成
		v3.PUT("/info/cancel_done/:task_id", handler.CancelDone) //1
	}
}

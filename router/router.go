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
	r.POST("/create_team", handler.CreateTeam)
	r.POST("/join_team", handler.JoinTeam)
	r.POST("/create_project/:team", handler.CreateProject)
}

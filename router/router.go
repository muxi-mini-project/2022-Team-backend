package router

import (
	"2022-TEAM-BACKEND/handler"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.POST("/user", handler.User)
	r.POST("/login", handler.Login)
	r.GET("/info", handler.Userinfo)
	r.PUT("/info", handler.ChangeInfomation)
	r.PUT("/create_team", handler.CreateTeam)
}

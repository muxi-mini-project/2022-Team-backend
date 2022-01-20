package router

import (
	"2022-TEAM-BACKEND/handler"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.POST("/user", handler.User)

}

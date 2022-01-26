package router

import (
	"team/model"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": "验证失败"})
	}
	c.Set("user_id", id)
}

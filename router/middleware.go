package router

import (
	"fmt"
	"team/model"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		id, err := model.VerifyToken(token)
		fmt.Println(token)

		fmt.Println(id, "s")
		c.Set("id", id)
		if err != nil || id == 0 {
			c.JSON(401, gin.H{
				"code":    401,
				"message": "身份验证失败"})
			c.Abort()
		}
	}
}

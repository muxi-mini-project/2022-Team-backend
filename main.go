package main

import (
	"2022-TEAM-BACKEND/config"
	"2022-TEAM-BACKEND/model"
	_ "fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/spf13/viper"
	_ "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Team
// @version 1.0.0
// @description 一款面向小型团队的任务进度共享软件
// @termsOfService http://swagger.io/terrms
// @host localhost:8918
// @BasePath:/api/v1
// @Schemes http

func main() {
	config.ConfigInit()
	model.Initdb()
	r := gin.Default()
	r.Run(":9918")
}

package main

import (
	"2022-TEAM-BACKEND/config"
	"2022-TEAM-BACKEND/model"
	"2022-TEAM-BACKEND/router"
	"fmt"

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
	r := gin.Default() //创建带有默认中间件的路由
	config.ConfigInit()
	//注意大写规范
	model.DB = model.Initdb()

	router.Router(r)
	if err := r.Run(":9918"); err != nil {
		fmt.Println(err)
	}

}

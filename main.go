package main

import (
	"flag"
	"fmt"
	"team/config"
	"team/model"
	"team/router"
	"team/services/flag_handle"

	"github.com/gin-gonic/gin"
	_ "github.com/spf13/viper"
	_ "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Team
// @version 1.0.0
// @description 一款面向小型团队的任务进度共享软件
// @termsOfService http://swagger.io/terrms
// @host localhost:9918
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

func init() {
	port := flag.String("port", "9981", "本地监听的端口")
	platform := flag.String("platform", "gitee", "平台名称，支持gitee/github")
	token := flag.String("token", "0cc7adde90fd9c09ebe2f8a8326dd943", "Gitee/Github 的用户授权码")
	owner := flag.String("owner", "Alen-H", "仓库所属空间地址(企业、组织或个人的地址path)")
	repo := flag.String("repo", "imagebed", "仓库路径(path)")
	path := flag.String("path", "", "文件的路径")
	branch := flag.String("branch", "master", "分支")
	flag.Parse()
	flag_handle.PORT = *port
	flag_handle.OWNER = *owner
	flag_handle.REPO = *repo
	flag_handle.PATH = *path
	flag_handle.TOKEN = *token
	flag_handle.PLATFORM = *platform
	flag_handle.BRANCH = *branch
	if flag_handle.TOKEN == "" {
		panic("token 必须！")
	}
}

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
// @contact.name Eternal-Faith
// @contact.email 2295616516@qq.com
// @host 122.112.236.36:9918
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
	flag.Parse()
	flag_handle.PORT = *flag.String("port", "9918", "本地监听的端口")
	flag_handle.OWNER = *flag.String("owner", "Alen-H", "仓库所属空间地址(企业、组织或个人的地址path)")
	flag_handle.REPO = *flag.String("repo", "imagebed", "仓库路径(path)")
	flag_handle.PATH = *flag.String("path", "", "文件的路径")
	flag_handle.TOKEN = *flag.String("token", "0cc7adde90fd9c09ebe2f8a8326dd943", "Gitee/Github 的用户授权码")
	flag_handle.PLATFORM = *flag.String("platform", "gitee", "平台名称，支持gitee/github")
	flag_handle.BRANCH = *flag.String("branch", "master", "分支")
	if flag_handle.TOKEN == "" {
		panic("token 必须！")
	}
}

package config

import (
	"log"

	"github.com/spf13/viper"
)

func ConfigInit() {
	viper.SetConfigFile("./conf/config.yaml") //指定配置文件路径
	viper.SetConfigName("config")             //配置文件的文件名，没有扩展名
	viper.AddConfigPath("./conf")             //查找配置文件所在路径，“.”在工作目录中搜索配置文件
	err := viper.ReadInConfig()               //搜索并读取配置文件

	if err != nil {
		log.Fatal(err)
	}
}

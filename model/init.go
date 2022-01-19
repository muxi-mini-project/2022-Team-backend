package model

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//DB 全局变量

func Initdb() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True", viper.Get("mysql.user"), viper.Get("mysql.password"), viper.Get("mysql.host"), viper.Get("mysql.port"), viper.Get("mysql.db"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) //连接数据库的固定格式
	if err != nil {
		log.Fatal(err) //1.打印输出内容 2.退出应用程序 3.defer函数不会执行
	}
	// 注意：port的端口号和main函数的端口号不能相同，不然会按占用处理
	return db
}

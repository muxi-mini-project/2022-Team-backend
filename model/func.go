package model

import (
	// "errors"
	"fmt"
	"log"
	"strconv"
	// "time"
	// "github.com/dgrijalva/jwt-go"
)

func Register(phone string, password string) string {
	user := User{Phone: phone, Password: password} //结构体里的值不一定都要用，一个包里的东西用的时候就当一个go文件就行了
	if err := DB.Table("user").Create(&user).Error; err != nil {
		fmt.Println("注册出错" + err.Error()) //err.Error打印错误
		return " "
	}
	Id := strconv.Itoa(user.UserId)
	return Id
}
func InitInfo(nickname string, avatar string) error {
	user := User{NickName: nickname, Avatar: avatar}
	if err := DB.Table("user").Where("id = ?", user.UserId).Updates(map[string]interface{}{"nickname": user.NickName, "avatar": user.Avatar}).Error; err != nil {
		return err
	}
	return nil
}

//防止电话重复绑定
func IfExistUserPhone(Phone string) (error, int) {
	var temp User
	if err := DB.Table("user").Where("phone = ?", Phone).Find(&temp).Error; err != nil {
		log.Println(err) //比fmt.Println多时间戳
		return err, 1
	}
	return nil, 0
}

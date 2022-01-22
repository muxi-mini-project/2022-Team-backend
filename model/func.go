package model

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func Register(phone string, password string) string {
	user := User{Phone: phone, Password: password} //结构体里的值不一定都要用，一个包里的东西用的时候就当一个go文件就行了
	//if里声明的变量只能在if里用
	if err := DB.Table("user").Create(&user).Error; err != nil {
		fmt.Println("注册出错" + err.Error()) //err.Error打印错误
		return " "
	}
	Id := strconv.Itoa(user.UserId)
	return Id
}

// func InitInfo(nickname string, avatar string) error {
// 	user := User{NickName: nickname, Avatar: avatar}
// 	if err := DB.Table("user").Where("id = ?", user.UserId).Updates(map[string]interface{}{"nickname": user.NickName, "avatar": user.Avatar}).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

//防止电话重复绑定331,如果有这条数据则说明该电话号码已被注册
func IfExistUserPhone(phone string) (error, int) {
	var temp User
	if err := DB.Table("user").Where("phone = ?", phone).Find(&temp).Error; temp.Phone == "" {
		log.Println(err) //比fmt.Println多时间戳
		// fmt.Println("hh", err)
		return err, 1
	}
	fmt.Println(temp)
	return nil, 0
}

//验证用户是否存在29
func VerifyPhone(phone string) bool {
	var user = make([]User, 1) //分配一个结构体
	if err := DB.Table("user").Where("phone=?", phone).Find(&user).Error; err != nil {
		log.Println(err.Error())
		return false
	}
	//结构体的长度？
	if len(user) != 1 {
		fmt.Println(len(user))
		return true
	}
	return false
}

//验证密码
func VerifyPassword(phone string, password string) bool {
	var user User
	if err := DB.Table("user").Where("phone = ? and password = ?", phone, password).Find(&user).Error; err != nil {
		log.Println(err.Error())
		return false
	}
	//觉得这里不需要再验证密码和电话了
	return true
}

//phone唯一对应用户了，不需要获取用户id
//生成token与验证

type jwtClaims struct {
	jwt.StandardClaims        //jwt-go包预定义的一些字段
	Phone              string `json:"phone"`
}

var (
	key        = "miniProject"
	ExpireTime = 604800 //token过期时间
)

//我自己往token里写进去的只有phone
func GenerateToken(phone string) string {
	claims := &jwtClaims{
		Phone: phone,
	}
	//签发者和过期时间
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
	singedToken, err := genToken(*claims)
	if err != nil {
		log.Print("produceToken err:")
		fmt.Println(err)
		return ""
	}
	return singedToken
}

func genToken(claims jwtClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	singedToken, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return singedToken, nil
}

//验证token
func VerifyToken(token string) (string, error) {
	TempToken, err := jwt.ParseWithClaims(token, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return "", errors.New("token解析失败")
	}
	claims, ok := TempToken.Claims.(*jwtClaims)
	if !ok {
		return "", errors.New("发生错误")
	}
	if err := TempToken.Claims.Valid(); err != nil {
		return "", errors.New("发生错误")
	}
	fmt.Println(claims.Phone)
	return claims.Phone, nil
}

//获取用户信息
func GetUserInfo(phone string) (User, error) {
	var user User
	if err := DB.Table("user").Where("phone=?", phone).Find(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func ChangeUserInfo(user User) error {
	fmt.Println(user.Phone)
	if err := DB.Table("user").Where("phone=?", user.Phone).Updates(map[string]interface{}{"nickname": user.NickName, "avatar": user.Avatar}).Error; err != nil {
		return err
	}
	return nil
}

func RegisterTeam(teamName string, avatar string, creator string, teamCoding string) error {
	team := Team{TeamName: teamName, Avatar: avatar, Creator: creator, TeamCoding: teamCoding}
	if err := DB.Table("team").Create(&team).Error; err != nil {
		fmt.Println("注册团队出错" + err.Error()) //err.Error打印错误
		return err
	}
	return nil
}

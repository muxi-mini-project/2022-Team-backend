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

//获取用户信息
func GetUserId(phone string) (User, error) {
	var user User
	if err := DB.Table("user").Where("phone=?", phone).Find(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

//初始化信息
func InitInfo(id int, nickname string) error {
	user := User{UserId: id, NickName: nickname}
	if err := DB.Table("user").Where("id = ?", user.UserId).Updates(map[string]interface{}{"nickname": user.NickName}).Error; err != nil {
		return err
	}
	return nil
}

//用户反馈
func UserFeedback(id int, feedback string) error {
	user := User{UserId: id, Feedback: feedback}
	if err := DB.Table("user").Where("id = ?", user.UserId).Updates(map[string]interface{}{"feedback": user.Feedback}).Error; err != nil {
		return err
	}
	return nil
}

//防止电话重复绑定331,如果有这条数据则说明该电话号码已被注册
func IfExistUserPhone(phone string) (error, int) {
	var temp User
	if err := DB.Table("user").Where("phone = ?", phone).Find(&temp).Error; temp.Phone == "" { //电话为空说明数据库中没有这个电话
		log.Println(err) //比fmt.Println多时间戳
		// fmt.Println("hh", err)
		return err, 1
	}
	fmt.Println(temp)
	return nil, 0
}

//防止用户名重复
func IfExistNickname(nickname string) (error, int) {
	var temp User
	if err := DB.Table("user").Where("nickname = ?", nickname).Find(&temp).Error; temp.NickName == "" { //电话为空说明数据库中没有这个电话
		log.Println(err) //比fmt.Println多时间戳
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

//生成token与验证

type jwtClaims struct {
	jwt.StandardClaims     //jwt-go包预定义的一些字段
	Id                 int `json:"id"`
}

var (
	key        = "miniProject"
	ExpireTime = 604800 //token过期时间
)

//我自己往token里写进去的只有id
func GenerateToken(id int) string {
	claims := &jwtClaims{
		Id: id,
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
func VerifyToken(token string) (int, error) {
	TempToken, err := jwt.ParseWithClaims(token, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return 0, errors.New("token解析失败")
	}
	claims, ok := TempToken.Claims.(*jwtClaims)
	if !ok {
		return 0, errors.New("发生错误")
	}
	if err := TempToken.Claims.Valid(); err != nil {
		return 0, errors.New("发生错误")
	}
	fmt.Println(claims.Id, "hhh")
	return claims.Id, nil
}

//获取用户信息
func GetUserInfo(uid int) (User, error) {
	var user User
	if err := DB.Table("user").Where("id=?", uid).Find(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

//修改个人信息
func ChangeUserInfo(user User) error {
	fmt.Println(user.Phone)
	if err := DB.Table("user").Where("phone=?", user.Phone).Updates(map[string]interface{}{"nickname": user.NickName, "avatar": user.Avatar}).Error; err != nil {
		return err
	}
	return nil
}

//注册团队
func RegisterTeam(teamName string, creator_id int, teamCoding string, avatar string) int {
	team := Team{TeamName: teamName, CreatorId: creator_id, TeamCoding: teamCoding, Avatar: avatar}
	if err := DB.Table("team").Create(&team).Error; err != nil {
		fmt.Println("注册团队出错" + err.Error()) //err.Error打印错误
		return 0
	}
	return team.TeamId
}

//防止团队名重复
func IfExistTeamname(teamname string) (error, int) {
	var temp Team
	if err := DB.Table("team").Where("name = ?", teamname).Find(&temp).Error; temp.TeamName == "" {
		log.Println(err) //比fmt.Println多时间戳
		return err, 1
	}
	fmt.Println(temp)
	return nil, 0
}

//加入团队
func JoinTeam(userId int, teamId int) error {
	team := UserTeam{UserId: userId, TeamId: teamId}
	if err := DB.Table("user_team").Create(&team).Error; err != nil {
		fmt.Println("加入团队出错" + err.Error()) //err.Error打印错误
		return err
	}
	return nil
}

//获取团队信息GetTeamInfo
func GetTeamInfo(id string) (Team, []User, error) {
	var team Team
	if err := DB.Table("team").Where("id=?", id).Find(&team).Error; err != nil {
		return Team{}, []User{}, err
	}
	var uId []string
	if err := DB.Table("user_team").Where("team_id=?", id).Select("user_id").Find(&uId).Error; err != nil {
		return Team{}, []User{}, err
	}
	var userInfo []User
	if err := DB.Table("user").Where("id in (?)", uId).Find(&userInfo).Error; err != nil {
		return Team{}, []User{}, err
	}
	return team, userInfo, nil
}

//获取用户加入的所有团队
func GetAllTeamInfo(id int) ([]Team, error) {
	var tId []string
	if err := DB.Table("user_team").Where("user_id=?", id).Select("team_id").Find(&tId).Error; err != nil {
		return []Team{}, err
	}
	var teamInfo []Team
	if err := DB.Table("team").Where("id in (?)", tId).Find(&teamInfo).Error; err != nil {
		return []Team{}, err
	}
	return teamInfo, nil
}

//创建项目
func CreatePro(proName string, creator_id int, startTime string, deadline string, remark string, teamId int) int {
	project := Project{ProjectName: proName, CreatorId: creator_id, CreateTime: time.Now().Format("2006-01-02 15:04:00"), StartTime: startTime, Deadline: deadline, Remark: remark, TeamId: teamId}
	if err := DB.Table("project").Create(&project).Error; err != nil {
		fmt.Println("项目创建出错" + err.Error()) //err.Error打印错误
		return 0
	}
	return project.ProjectId
}

//添加步骤信息
func AddStep(name string, pId int) error {
	Step := Step{StepName: name, ProjectId: pId}
	if err := DB.Table("step").Create(&Step).Error; err != nil {
		return err
	}
	return nil
}

//获取团队成员id
//没查到return nil
func GetTeamMenberId(tId string) []string {
	var Id []string
	var userTeam []UserTeam
	var temp string
	if err := DB.Table("user_team").Where("team_id=?", tId).Find(&userTeam).Error; err != nil {
		log.Println(err)
		return nil
	} else {
		fmt.Println(userTeam)
		for _, id := range userTeam {
			temp = strconv.Itoa(id.UserId)
			Id = append(Id, string(temp))
		}
		fmt.Println(Id)
		return Id
	}
}

//获取团队成员名字
func GetTeamMenberName(UsersId []string) ([]string, error) {
	var name []string
	if err := DB.Table("user").Where("id in (?)", UsersId).Select(("nickname")).Find(&name).Error; err != nil {
		return nil, err
	}
	return name, nil
}

//获取团队所有项目名
func GetTeamPro(tId string) ([]string, error) {
	fmt.Println(tId, "test")
	var name []string
	if err := DB.Table("project").Where("team_id =? ", tId).Select("name").Find(&name).Error; err != nil {
		return nil, err
	}
	return name, nil
}

//在新建任务界面选择项目后返回步骤(是不是把team_id加进来更好些？)
func GetProStep(proName string, tId int) ([]string, error) {
	var pId int
	if err := DB.Table("project").Where("name = ? and team_id=? ", proName, tId).Select("id").Find(&pId).Error; err != nil {
		return nil, err
	}
	var name []string
	if err := DB.Table("step").Where("project_id =? ", pId).Select("name").Find(&name).Error; err != nil {
		return nil, err
	}
	return name, nil
}

//获取项目步骤的id
func GetProStepId(tId int, proName string, stepName string) (int, error) {
	var pId int
	if err := DB.Table("project").Where("name = ? and team_id = ?", proName, tId).Select("id").Find(&pId).Error; err != nil {
		return 0, err
	}
	fmt.Println("274", pId)
	var id int
	if err := DB.Table("step").Where("project_id = ? and  name = ?", pId, stepName).Select("id").Find(&id).Error; err != nil {
		return 0, err
	}
	fmt.Println("279", id)
	return id, nil
}

//获取未完成的任务
func GenToDoList(Uid int) ([]Task, []Team) {
	var Id []string
	var userTask []UserTask
	var Task []Task
	var team_id []string
	var team []Team
	if err := DB.Table("user_task").Where("principal_id=? and performance=?", Uid, false).Find(&userTask).Error; err != nil {
		log.Println(err)
	} else {
		fmt.Println(userTask)
		for _, id := range userTask {
			temp := strconv.Itoa(id.TaskId)
			Id = append(Id, temp)
		}
		if err := DB.Table("task").Where("id in (?)", Id).Find(&Task).Error; err != nil {
			log.Println(err)
			return nil, nil
		}
		if err := DB.Table("task").Where("id in (?)", Id).Select("team_id").Find(&team_id).Error; err != nil {
			log.Println(err)
			return nil, nil
		}
		if err := DB.Table("team").Where("id in (?)", team_id).Find(&team).Error; err != nil {
			log.Println(err)
			return nil, nil
		}
	}
	return Task, team
}

//获取完成的任务
func GenDoneList(Uid int) ([]Task, []Team) {
	var Id []string
	var userTask []UserTask
	var Task []Task
	var team_id []string
	var team []Team
	if err := DB.Table("user_task").Where("principal_id=? and performance=?", Uid, true).Find(&userTask).Error; err != nil {
		log.Println(err)
	} else {
		fmt.Println(userTask)
		for _, id := range userTask {
			temp := strconv.Itoa(id.TaskId)
			Id = append(Id, temp)
		}
		if err2 := DB.Table("task").Where("id in (?)", Id).Find(&Task).Error; err != nil {
			log.Println(err2)
			return nil, nil
		}
		if err := DB.Table("task").Where("id in (?)", Id).Select("team_id").Find(&team_id).Error; err != nil {
			log.Println(err)
			return nil, nil
		}
		if err := DB.Table("team").Where("id in (?)", team_id).Find(&team).Error; err != nil {
			log.Println(err)
			return nil, nil
		}
	}
	return Task, team
}

//修改密码
func ModifyPassword(id int, newPassword string) error {
	err := DB.Table("user").Where("id=?", id).Updates(map[string]interface{}{"password": newPassword}).Error
	if err != nil {
		return err
	}
	return nil
}

//完成任务
func CompleteTask(id string, uId int) error {
	if err := DB.Table("user_task").Where("task_id = ? and principal_id = ?", id, uId).Updates(map[string]interface{}{"performance": true}).Error; err != nil {
		return err
	}
	return nil
}

//取消完成状态
func CancelComplete(id string, uId int) error {
	if err := DB.Table("user_task").Where("task_id = ? and principal_id = ?", id, uId).Updates(map[string]interface{}{"performance": false}).Error; err != nil {
		return err
	}
	return nil
}

//删除项目
//先删子表后删父表
func RemoveProject(id string) error {
	var step Step
	if err := DB.Table("step").Where("project_id=?", id).Delete(&step).Error; err != nil {
		return err
	}
	var project Project
	if err := DB.Table("project").Where("id=?", id).Delete(&project).Error; err != nil {
		return err
	}
	return nil
}

//获取项目信息
func GetProjectInfo(id string) (Project, error) {
	var project Project
	if err := DB.Table("project").Where("id=?", id).Find(&project).Error; err != nil {
		return Project{}, err
	}
	if err := DB.Table("step").Where("project_id=?", id).Select("name").Find(&project.Step).Error; err != nil {
		return Project{}, err
	}
	return project, nil
}

//修改项目信息
func ChangeProjectInfo(project Project, id int) error {
	if err := DB.Table("project").Where("id=?", id).Updates(map[string]interface{}{"name": project.ProjectName, "start_time": project.StartTime, "deadline": project.Deadline, "remark": project.Remark}).Error; err != nil {
		return err
	}
	return nil
}

//创建任务
func CreateTask(taskName string, creator_id int, startTime string, deadline string, remark string, stepId int, project string, step string, teamId int) int {
	task := Task{TaskName: taskName, CreatorId: creator_id, CreateTime: time.Now().Format("2006-01-02 15:04:00"), StartTime: startTime, Deadline: deadline, Remark: remark, StepId: stepId, ProName: project, StepName: step, TeamId: teamId}
	if err := DB.Table("task").Create(&task).Error; err != nil {
		fmt.Println("任务创建出错(库)" + err.Error()) //err.Error打印错误
		return 0
	}
	return task.TaskId
}

//布置任务
func AssginIntoTable(UId int, Tid int, uName string, performance bool) error {
	UT := UserTask{UserId: UId, TaskId: Tid, PrincipalName: uName, Performance: performance}
	if err := DB.Table("user_task").Create(&UT).Error; err != nil {
		return err
	}
	return nil
}

//获取任务信息(包括成员名字)
func GetTaskInfo(id string) (Task, error) {
	var task Task
	if err := DB.Table("task").Where("id=?", id).Find(&task).Error; err != nil {
		return Task{}, err
	}
	if err := DB.Table("user_task").Where("task_id=?", id).Select("principal_name").Find(&task.Member).Error; err != nil {
		return Task{}, err
	}
	return task, nil
}

//修改任务信息???
func ChangeTaskInfo(task Task) error {
	var userTask UserTask
	if err := DB.Table("user_task").Where("task_id=?", task.TaskId).Delete(&userTask).Error; err != nil {
		return err
	}
	if err := DB.Table("task").Where("id=?", task.TaskId).Updates(map[string]interface{}{"name": task.TaskName, "start_time": task.StartTime, "deadline": task.Deadline, "remark": task.Remark, "step_id": task.StepId, "project": task.ProName, "step": task.StepName}).Error; err != nil {
		return err
	}
	return nil
}

// //创建任务
// func CreateTask(taskName string, creator_id int, startTime string, deadline string, remark string, stepId int, project string, step string) int {
// 	task := Task{TaskName: taskName, CreatorId: creator_id, CreateTime: time.Now().Format("2006-01-02 15:04:00"), StartTime: startTime, Deadline: deadline, Remark: remark, StepId: stepId, ProName: project, StepName: step}
// 	if err := DB.Table("task").Create(&task).Error; err != nil {
// 		fmt.Println("任务创建出错(库)" + err.Error()) //err.Error打印错误
// 		return 0
// 	}
// 	return task.TaskId
// }

//删除任务
func RemoveTask(id string) error {
	var userTask UserTask
	if err := DB.Table("user_task").Where("task_id=?", id).Delete(&userTask).Error; err != nil {
		return err
	}
	var task Task
	if err := DB.Table("task").Where("id=?", id).Delete(&task).Error; err != nil {
		return err
	}
	return nil
}

//修改用户头像
func UpdateAvator(avatar User) error {
	if err := DB.Table("user").Where("id = ?", avatar.UserId).Updates(map[string]interface{}{"avatar": avatar.Avatar, "sha": avatar.Sha, "path": avatar.Path}).Error; err != nil {
		return err
	}
	return nil
}

package model

//gorm:"column:id",是为为对应的sql语句中的字段起名字。gorm自动生成sql时大写字母转小写还会生成下划线
//id应该都设为int不然输不进去数据
//前端传数据时一定要传有json的tag的数据
type User struct {
	UserId   int    `json:"user_id" gorm:"column:id"`
	Phone    string `json:"phone"`
	NickName string `json:"nickname" gorm:"column:nickname"`
	Avatar   string `json:"avatar"`
	Password string `json:"password"`
	Feedback string `json:"feedback"`
}

type Team struct {
	TeamId     int    `json:"team_id" gorm:"column:id"`
	TeamName   string `json:"teamname" gorm:"column:name"`
	Avatar     string `json:"avatar"`
	Creator    string `json:"creator"`
	TeamCoding string `json:"team_coding"`
}

type UserTeam struct {
	Id       int    `json:"id"`
	UserName string `json:"username" gorm:"column:username"`
	TeamName string `json:"teamname" gorm:"column:teamname"`
}

type Project struct {
	ProjectId   int    `json:"project_id" gorm:"column:id"`
	ProjectName string `json:"project_name" gorm:"column:name"`
	Creator     string `json:"creator"`
	CreateTime  string `json:"create_time"`
	StartTime   string `json:"start_time"`
	Deadline    string `json:"deadline"`
	Remark      string `json:"remark"`
	TeamName    string `json:"teamname" gorm:"column:team"`
}

type Step struct {
	StepId      int    `json:"step_id" gorm:"column:id"`
	StepName    string `json:"step_name" gorm:"column:name"`
	CreateTime  string `json:"createtime" gorm:"column:create_team"`
	StartTime   string `json:"start_time" gorm:"column:start_team"`
	Deadline    string `json:"deadline"`
	Remark      string `json:"remark"`
	ProjectName string `json:"project_name" gorm:"column:project"`
}

type Task struct {
	TaskId  int    `json:"task_id" gorm:"column:id"`
	Content string `json:"content" `
	Step    string `json:"step"`
}

type UserTask struct {
	Id          int    `json:"id"`
	UserName    string `json:"username" gorm:"column:principal"`
	TaskName    string `json:"taskname" gorm:"column:task_id"`
	Performance bool   `json:"performance"`
}

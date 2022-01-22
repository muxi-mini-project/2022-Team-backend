package model

//gorm:"column:id",是为为对应的sql语句中的字段起名字。gorm自动生成sql时大写字母转小写还会生成下划线
type User struct {
	UserId   int    `json:"user_id" gorm:"column:id"`
	Phone    string `json:"phone"`
	NickName string `json:"nickname" gorm:"column:nickname"`
	Avatar   string `json:"avatar"`
	Password string `json:"password"`
	Feedback string `json:"feedback"`
}

type Team struct {
	TeamId     string `json:"team_id" gorm:"column:id`
	TeamName   string `json:"teamname" gorm:"column:name`
	Avatar     string `json:"avatar"`
	Creator    string `json:"creator"`
	TeamCoding string `json:"team_coding"`
}

type UserTeam struct {
	Id       string `json:"id"`
	UserName string `json:"username"`
	TeamName string `json:"teamname"`
}

type Project struct {
	ProjectId     string `json:"project_id"`
	ProjectName   string `json:"project_name"`
	Creator       string `json:"creator"`
	CreateTime    string `json:"createtime"`
	StartTime     string `json:"starttime"`
	Deadline      string `json:"deadline"`
	ProjectRemark string `json:"projectremark"`
	TeamName      string `json:"teamname"`
}

type Step struct {
	StepId     string `json:"step_id"`
	StepName   string `json:"step_name"`
	CreateTime string `json:"createtime"`
	StartTime  string `json:"starttime"`
	Deadline   string `json:"deadline"`
	StepRemark string `json:"stepremark"`
	TeamName   string `json:"teamname"`
}

type Task struct {
	TaskId  string `json:"task_id"`
	Content string `json:"content"`
	Step    string `json:"step"`
}

type UserTask struct {
	Id          string `json:"id"`
	UserName    string `json:"username"`
	TaskName    string `json:"taskname"`
	Performance bool   `json:"performance"`
}

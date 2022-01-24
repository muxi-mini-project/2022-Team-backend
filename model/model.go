package model

//注意go语言驼峰式格式
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
	CreatorId  int    `json:"creator_id" `
	TeamCoding string `json:"team_coding"`
}

type UserTeam struct {
	Id     int `json:"id"`
	UserId int `json:"user_id" `
	TeamId int `json:"team_id" `
}

type Project struct {
	ProjectId   int    `json:"project_id" gorm:"column:id"`
	ProjectName string `json:"project_name" gorm:"column:name"`
	CreatorId   int    `json:"creator_id"`
	CreateTime  string `json:"create_time"`
	StartTime   string `json:"start_time"`
	Deadline    string `json:"deadline"`
	Remark      string `json:"remark"`
	TeamId      int    `json:"team_id"`
}

type Step struct {
	StepId    int    `json:"step_id" gorm:"column:id"`
	StepName  string `json:"step_name" gorm:"column:name"`
	ProjectId int    `json:"project_id"`
}

type Task struct {
	TaskId     int    `json:"task_id" gorm:"column:id"`
	TaskName   string `json:"name" `
	CreateTime string `json:"createtime" gorm:"column:create_time"`
	Date       string `json:"date"`
	StartTime  string `json:"start_time" gorm:"column:start_time"`
	Deadline   string `json:"deadline"`
	Remark     string `json:"remark"`
	StepId     int    `json:"step_id"`
}

type UserTask struct {
	Id          int  `json:"id"`
	UserId      int  `json:"principal_id" gorm:"column:principal_id"`
	TaskID      int  `json:"task_id" gorm:"column:task_id"`
	Performance bool `json:"performance"`
}

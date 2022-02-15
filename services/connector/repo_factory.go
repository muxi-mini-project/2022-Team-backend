package connector

import (
	"team/services"
	"team/services/flag_handle"
	"team/services/gitee"
)

//定义serve的映射关系
var serveMap = map[string]services.RepoInterface{
	"gitee": &gitee.GiteeServe{},
}

func RepoCreate() services.RepoInterface {
	return serveMap[flag_handle.PLATFORM]
}

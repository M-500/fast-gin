package api

import (
	"fast-gin/app/api/admin"
	"fast-gin/app/api/na"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/12/5 10:15
//

type ApiGroups struct {
	NaApiGroups    na.NaApis
	AdminApiGroups admin.AdminApis
}

func NewApiGroups() *ApiGroups {
	return &ApiGroups{}
}

var APIGroupsAPP = NewApiGroups()

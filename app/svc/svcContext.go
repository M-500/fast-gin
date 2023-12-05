package svc

import (
	"fast-gin/app/config"
	"fast-gin/pkg/comm/validators"
	"fast-gin/pkg/utils/db_helper"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"xorm.io/xorm"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/12/4 12:27
//

var svc *ServiceContext

type ServiceContext struct {
	Config *config.Config
	Server *gin.Engine
	DbConn *xorm.Engine
	Trans  ut.Translator
}

func SetUpServiceContext(c *config.Config) *ServiceContext {
	if svc != nil {
		return svc
	}
	svc = &ServiceContext{
		Config: c,
		DbConn: db_helper.GetDb(),
		Trans:  validators.SetUpTrans(c.Local),
	}
	return svc
}

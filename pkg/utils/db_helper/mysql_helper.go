package db_helper

import (
	"fast-gin/app/config"
	"fmt"
	"log"
	"time"
	"xorm.io/xorm"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/12/4 12:12
//

var dbEngine *xorm.Engine

func SetUpDB() {
	if dbEngine != nil {
		return
	}
	sourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		config.ConfigInstance.DB.Username,
		config.ConfigInstance.DB.Password,
		config.ConfigInstance.DB.Host,
		config.ConfigInstance.DB.Port,
		config.ConfigInstance.DB.Database,
		config.ConfigInstance.DB.Charset,
	)
	if engine, err := xorm.NewEngine(config.ConfigInstance.DB.Engine, sourceName); err != nil {
		log.Fatalf("db_helper.InitDb(%s),%v\n", sourceName, err)
	} else {
		dbEngine = engine
	}

	if config.ConfigInstance.DB.MaxIdleConns > 0 {
		dbEngine.SetMaxIdleConns(config.ConfigInstance.DB.MaxIdleConns)
	}

	if config.ConfigInstance.DB.MaxOpenConns > 0 {
		dbEngine.SetMaxOpenConns(config.ConfigInstance.DB.MaxOpenConns)
	}
	if config.ConfigInstance.DB.ConnMaxLifetime > 0 {
		dbEngine.SetConnMaxLifetime(time.Minute * time.Duration(config.ConfigInstance.DB.ConnMaxLifetime))
	}
}

func GetDb() *xorm.Engine {
	if dbEngine == nil {
		SetUpDB()
	}
	return dbEngine
}

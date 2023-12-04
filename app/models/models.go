package models

import (
	"time"
)

type ExampleUser struct {
	Id         int        `xorm:"not null pk autoincr comment('主键ID') INT"`
	Username   string     `xorm:"not null comment('用户名') VARCHAR(32)"`
	Phone      string     `xorm:"comment('手机号') unique CHAR(11)"`
	Gender     int        `xorm:"comment('1男，0女，2保密') TINYINT"`
	Addr       string     `xorm:"comment('地址') VARCHAR(255)"`
	Pwd        string     `xorm:"comment('密码(密文)') VARCHAR(255)"`
	CreateTime *time.Time `xorm:"not null comment('创建时间') DATETIME"`
	UpdateTime *time.Time `xorm:"not null comment('更新时间') DATETIME"`
	DeleteTime *time.Time `xorm:"comment('删除时间') DATETIME"`
}

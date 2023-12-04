package dao

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/12/4 11:51
//
import (
	"context"
	"xorm.io/xorm"
)

type ExampleDao struct {
	db  *xorm.Engine
	ctx context.Context
}

func NewExampleDao(ctx context.Context) *ExampleDao {
	return &ExampleDao{
		//db: ,
		ctx: ctx,
	}
}

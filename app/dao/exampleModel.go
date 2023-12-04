package dao

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/12/4 11:51
//
import (
	"context"
	"fast-gin/app/models"
	"fast-gin/pkg/utils/db_helper"
	"xorm.io/xorm"
)

type ExampleDao struct {
	db  *xorm.Engine
	ctx context.Context
}

func NewExampleDao(ctx context.Context) *ExampleDao {
	return &ExampleDao{
		db:  db_helper.GetDb(),
		ctx: ctx,
	}
}

func (dao *ExampleDao) GetById(id int) (*models.ExampleUser, error) {
	data := &models.ExampleUser{}
	_, err := dao.db.ID(id).Get(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (dao *ExampleDao) GetByUsername(username string) (*models.ExampleUser, error) {
	data := &models.ExampleUser{}
	_, err := dao.db.Where("`username` = ?", username).Get(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (dao *ExampleDao) GetByPhone(phone string) (*models.ExampleUser, error) {
	data := &models.ExampleUser{}
	_, err := dao.db.Where("`phone` = ?", phone).Get(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

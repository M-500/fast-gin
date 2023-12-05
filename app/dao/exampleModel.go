package dao

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/12/4 11:51
//
import (
	"context"
	"fast-gin/app/models"
	"fast-gin/pkg/utils"
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

func (dao *ExampleDao) Get(id int) (*models.ExampleUser, error) {
	data := &models.ExampleUser{}
	_, err := dao.db.ID(id).Get(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (dao *ExampleDao) FindByUsername(username string) (*models.ExampleUser, error) {
	data := &models.ExampleUser{}
	_, err := dao.db.Where("`username` = ?", username).Get(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (dao *ExampleDao) FindByPhone(phone string) (*models.ExampleUser, error) {
	data := &models.ExampleUser{}
	_, err := dao.db.Where("`phone` = ?", phone).Get(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (dao *ExampleDao) FindAllPager(username string, page int, size int) ([]models.ExampleUser, int64, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	start := (page - 1) * size
	dataList := make([]models.ExampleUser, 0)
	total, err := dao.db.Where("`username` like ?", username).Desc("id").Limit(size, start).FindAndCount(&dataList)
	return dataList, total, err
}

func (dao *ExampleDao) Insert(data *models.ExampleUser) error {
	data.CreateTime = utils.Now()
	data.UpdateTime = utils.Now()
	_, err := dao.db.Insert(data)
	return err
}

func (dao *ExampleDao) Update(data *models.ExampleUser, musColumns ...string) error {
	sess := dao.db.ID(data.Id)
	if len(musColumns) > 0 {
		sess.MustCols(musColumns...)
	}
	data.UpdateTime = utils.Now()
	_, err := sess.Update(data)
	return err
}

// 更高一层的封装
func (dao *ExampleDao) Save(data *models.ExampleUser, musColumns ...string) error {
	if data.Id > 0 {
		return dao.Update(data, musColumns...)
	}
	return dao.Insert(data)
}

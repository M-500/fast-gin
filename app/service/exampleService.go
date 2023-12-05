package service

import (
	"context"
	"fast-gin/app/dao"
	"fast-gin/app/models"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/12/4 11:50
//

type ExampleService struct {
	ctx           context.Context
	daoExampleDao *dao.ExampleDao
}

func NewExampleService(ctx context.Context) *ExampleService {
	return &ExampleService{
		ctx:           ctx,
		daoExampleDao: dao.NewExampleDao(ctx),
	}
}

func (s *ExampleService) Get(id int) (*models.ExampleUser, error) {
	return s.daoExampleDao.Get(id)
}

func (s *ExampleService) FindByUsername(username string) (*models.ExampleUser, error) {
	return s.daoExampleDao.FindByUsername(username)
}

func (s *ExampleService) FindByPhone(phone string) (*models.ExampleUser, error) {
	return s.daoExampleDao.FindByPhone(phone)
}

func (s *ExampleService) FindAllPager(username string, page int, size int) ([]models.ExampleUser, int64, error) {
	return s.daoExampleDao.FindAllPager(username, page, size)
}

func (s *ExampleService) Save(data *models.ExampleUser, musColumns ...string) error {
	return s.daoExampleDao.Save(data, musColumns...)
}

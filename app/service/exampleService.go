package service

import (
	"context"
	"fast-gin/app/dao"
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

func (s *ExampleService) Get(id int) {

}

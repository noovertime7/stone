package services

import (
	"context"
	"fmt"
	"github.com/noovertime7/stone/dao"
	"github.com/noovertime7/stone/dao/common"
	"github.com/noovertime7/stone/dto"
	"github.com/noovertime7/stone/runtime"
)

type StoneTypeService struct {
}

func (s *StoneTypeService) Get(ctx context.Context, id int) (dao.StoneTypes, error) {
	model := &dao.StoneTypes{}
	data, ok, err := model.Find(ctx, dao.StoneTypes{Id: id})
	if !ok {
		return dao.StoneTypes{}, fmt.Errorf("分类不存在")
	}
	return data, err
}

func (s *StoneTypeService) List(ctx context.Context) ([]*dao.StoneTypes, error) {
	model := dao.StoneTypes{}
	return model.FindList(ctx, dao.StoneTypes{})
}

func (s *StoneTypeService) Delete(ctx context.Context, id int) error {
	model := dao.StoneTypes{}
	return model.Delete(ctx, dao.StoneTypes{Id: id}, false)
}

func (s *StoneTypeService) Page(ctx context.Context, params runtime.DataBasePager) (dto.PageStoneTypesOut, error) {
	model := dao.StoneTypes{}
	data, total, err := model.PageList(ctx, params)
	return dto.PageStoneTypesOut{
		Total:    total,
		List:     data,
		Page:     params.GetPage(),
		PageSize: params.GetPageSize(),
	}, err
}

func (s *StoneTypeService) Update(ctx context.Context, id int, name string) error {
	model := dao.StoneTypes{}
	return model.Updates(ctx, common.WithIDOption(id, common.Equal), &dao.StoneTypes{Name: name})
}

func (s *StoneTypeService) Save(ctx context.Context, name string) error {
	model := dao.StoneTypes{}
	return model.Save(ctx, &dao.StoneTypes{Name: name})
}

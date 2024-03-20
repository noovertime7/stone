package services

import (
	"context"
	"fmt"
	"github.com/noovertime7/stone/dao"
	"github.com/noovertime7/stone/dao/common"
	"github.com/noovertime7/stone/dto"
	"github.com/noovertime7/stone/runtime"
)

type StoneService struct {
}

func (s *StoneService) Get(ctx context.Context, id int) (dao.Stone, error) {
	model := &dao.Stone{}
	data, ok, err := model.Find(ctx, dao.Stone{Id: id})
	if !ok {
		return dao.Stone{}, fmt.Errorf("分类不存在")
	}
	return data, err
}

func (s *StoneService) List(ctx context.Context) ([]*dao.Stone, error) {
	model := dao.Stone{}
	return model.FindList(ctx, dao.Stone{})
}

func (s *StoneService) Delete(ctx context.Context, id int) error {
	model := dao.Stone{}
	return model.Delete(ctx, dao.Stone{Id: id}, false)
}

func (s *StoneService) Page(ctx context.Context, params runtime.DataBasePager) (dto.PageStoneOut, error) {
	model := dao.Stone{}
	data, total, err := model.PageList(ctx, params)
	return dto.PageStoneOut{
		Total:    total,
		List:     data,
		Page:     params.GetPage(),
		PageSize: params.GetPageSize(),
	}, err
}

func (s *StoneService) Update(ctx context.Context, id int, name string) error {
	model := dao.Stone{}
	return model.Updates(ctx, common.WithIDOption(id, common.Equal), &dao.Stone{Name: name})
}

func (s *StoneService) Save(ctx context.Context, stone *dto.CreateStone) error {
	model := dao.Stone{
		Name:         stone.Name,
		CoverImages:  stone.CoverImages,
		DetailImages: stone.DetailImages,
		Description:  stone.Description,
		Hot:          dao.NotHot,
		StoneTypeId:  stone.StoneTypeId,
	}
	return model.Save(ctx, &model)
}

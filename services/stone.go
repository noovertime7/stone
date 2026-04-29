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
		return dao.Stone{}, fmt.Errorf("大理石不存在")
	}
	data.ViewCount = data.ViewCount + 1
	_ = model.Updates(ctx, common.WithIDOption(data.Id, common.Equal), &dao.Stone{ViewCount: data.ViewCount})
	return data, err
}

func (s *StoneService) List(ctx context.Context) ([]*dao.Stone, error) {
	model := dao.Stone{}
	return model.FindList(ctx, dao.Stone{})
}

func (s *StoneService) FindHotList(ctx context.Context) ([]*dao.Stone, error) {
	model := dao.Stone{}
	return model.FindList(ctx, dao.Stone{Hot: dao.Hot})
}

func (s *StoneService) FindStonesByTypeID(ctx context.Context, tid int) ([]*dao.Stone, error) {
	model := dao.Stone{}

	return model.FindList(ctx, dao.Stone{StoneTypeId: tid})
}

func (s *StoneService) FindSameTypeStones(ctx context.Context, id int) ([]*dao.Stone, error) {
	model := dao.Stone{Id: id}
	stone, ok, err := model.Find(ctx, model)
	if !ok {
		return nil, fmt.Errorf("大理石不存在")
	}
	if err != nil {
		return nil, err
	}
	return model.FindList(ctx, dao.Stone{StoneTypeId: stone.StoneTypeId})
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

func (s *StoneService) Update(ctx context.Context, id int, stone *dto.CreateStone) error {
	model := &dao.Stone{
		StoneTypeId:  stone.StoneTypeId,
		Name:         stone.Name,
		CoverImages:  stone.CoverImages,
		DetailImages: stone.DetailImages,
		Description:  stone.Description,
		Hot:          stone.Hot,
		Color:        stone.Color,
		Origin:       stone.Origin,
		Texture:      stone.Texture,
	}
	return model.Updates(ctx, common.WithIDOption(id, common.Equal), model)
}

func (s *StoneService) Save(ctx context.Context, stone *dto.CreateStone) error {
	model := dao.Stone{
		Name:         stone.Name,
		CoverImages:  stone.CoverImages,
		DetailImages: stone.DetailImages,
		Description:  stone.Description,
		Hot:          stone.Hot,
		StoneTypeId:  stone.StoneTypeId,
		Color:        stone.Color,
		Origin:       stone.Origin,
		Texture:      stone.Texture,
	}
	return model.Save(ctx, &model)
}

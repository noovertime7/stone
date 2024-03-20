package services

import (
	"context"
	"fmt"
	"github.com/e421083458/golang_common/log"
	"github.com/noovertime7/stone/dao"
	"github.com/noovertime7/stone/dao/common"
	"github.com/noovertime7/stone/dto"
	"github.com/noovertime7/stone/runtime"
	"math/rand"
	"time"
)

type StoneService struct {
}

func (s *StoneService) Get(ctx context.Context, id int) (dao.Stone, error) {
	model := &dao.Stone{}
	data, ok, err := model.Find(ctx, dao.Stone{Id: id})
	if !ok {
		return dao.Stone{}, fmt.Errorf("大理石不存在")
	}
	// 使用时间作为种子，确保每次运行生成的随机数都不同
	rand.Seed(time.Now().UnixNano())
	// 生成200到400之间的随机数
	randomNumber := rand.Intn(301) + 100
	if randomNumber > 300 {
		data.BuyNum = data.BuyNum + 1
		_ = model.Updates(ctx, common.WithIDOption(data.Id, common.Equal), &dao.Stone{BuyNum: data.BuyNum})
		log.Info(" 用户购买了大理石，购买次数加一")
	}
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

func (s *StoneService) Update(ctx context.Context, id int, name string) error {
	model := dao.Stone{}
	return model.Updates(ctx, common.WithIDOption(id, common.Equal), &dao.Stone{Name: name})
}

func (s *StoneService) Save(ctx context.Context, stone *dto.CreateStone) error {
	// 使用时间作为种子，确保每次运行生成的随机数都不同
	rand.Seed(time.Now().UnixNano())
	// 生成200到500之间的随机数
	randomNumber := rand.Intn(301) + 200
	model := dao.Stone{
		Name:         stone.Name,
		CoverImages:  stone.CoverImages,
		DetailImages: stone.DetailImages,
		Description:  stone.Description,
		Hot:          dao.NotHot,
		StoneTypeId:  stone.StoneTypeId,
		BuyNum:       randomNumber,
	}
	return model.Save(ctx, &model)
}

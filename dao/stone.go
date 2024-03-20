package dao

import (
	"context"
	"github.com/noovertime7/stone/dao/common"
	"github.com/noovertime7/stone/dao/gormdatatypes"
	"github.com/noovertime7/stone/runtime"
)

type HotType int

const (
	Hot    = 1
	NotHot = 2
)

type Stone struct {
	Id           int                     `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	StoneTypeId  int                     `json:"stoneTypeId"`
	Name         string                  `json:"name"`
	CoverImages  gormdatatypes.JSONSlice `json:"coverImages"`
	DetailImages gormdatatypes.JSONSlice `json:"detailImages"`
	Description  string                  `json:"description"`
	Hot          HotType                 `json:"hot"`
	BuyNum       int                     `json:"buyNum"`
	common.CommonModel
}

func (s *Stone) TableName() string {
	return "t_stone"
}

func (s *Stone) PageList(ctx context.Context, params runtime.DataBasePager, ops ...common.Option) ([]*Stone, int64, error) {
	var total int64
	query := GetDB().WithContext(ctx).Where("")
	limit := params.GetPageSize()
	offset := limit * (params.GetPage() - 1)

	var list []*Stone
	// 如果有条件搜索 下方会自动创建搜索语句
	if params.IsFitter() {
		params.Do(query)
	}

	for _, op := range ops {
		query = op(query)
	}

	if err := query.Find(&list).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Order("id desc").Limit(limit).Offset(offset).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (s *Stone) Save(ctx context.Context, obj *Stone) error {
	return GetDB().WithContext(ctx).Save(obj).Error
}

func (s *Stone) Updates(ctx context.Context, opt common.Option, search *Stone) error {
	query := opt(GetDB())
	return query.WithContext(ctx).Updates(&search).Error
}

func (s *Stone) Find(ctx context.Context, search Stone) (Stone, bool, error) {
	var out Stone
	return out, out.Id != 0, GetDB().WithContext(ctx).Where(&search).Find(&out).Error
}

func (s *Stone) FindList(ctx context.Context, search Stone, opts ...common.Option) ([]*Stone, error) {
	query := GetDB().WithContext(ctx)
	for _, opt := range opts {
		query = opt(query)
	}

	var out []*Stone
	return out, query.Where(&search).Find(&out).Error
}

func (s *Stone) Delete(ctx context.Context, search Stone, isDelete bool) error {
	if isDelete {
		return GetDB().WithContext(ctx).Where("id = ?", search.Id).Unscoped().Delete(&search).Error
	}
	return GetDB().WithContext(ctx).Where("id = ?", search.Id).Delete(&search).Error
}

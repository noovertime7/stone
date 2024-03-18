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
	Id           int                     `json:"id" gorm:"column:id"`
	StoneTypeId  string                  `json:"stoneTypeId"`
	Name         string                  `json:"name"`
	HomeImage    string                  `json:"homeImage"`
	Price        int64                   `json:"price" gorm:"comment:价格"`
	DetailImages gormdatatypes.JSONSlice `json:"detailImages"`
	Description  string                  `json:"description"`
	Hot          HotType                 `json:"hot"`
	common.CommonModel
}

func (f *Stone) TableName() string {
	return "t_stone"
}

func (a *Stone) PageList(ctx context.Context, params runtime.DataBasePager, ops ...common.Option) ([]*Stone, int64, error) {
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

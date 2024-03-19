package dao

import (
	"context"
	"github.com/noovertime7/stone/dao/common"
	"github.com/noovertime7/stone/runtime"
)

type StoneTypes struct {
	Id     int     `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	Name   string  `json:"name"`
	Stones []Stone `json:"stones" gorm:"-"`
	common.CommonModel
}

func (f *StoneTypes) TableName() string {
	return "t_StoneTypes_types"
}

func (s *StoneTypes) PageList(ctx context.Context, params runtime.DataBasePager, ops ...common.Option) ([]*StoneTypes, int64, error) {
	var total int64
	query := GetDB().WithContext(ctx).Where("")
	limit := params.GetPageSize()
	offset := limit * (params.GetPage() - 1)

	var list []*StoneTypes
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

func (s *StoneTypes) Save(ctx context.Context, obj *StoneTypes) error {
	return GetDB().WithContext(ctx).Save(obj).Error
}

func (s *StoneTypes) Updates(ctx context.Context, opt common.Option, search *StoneTypes) error {
	query := opt(GetDB())
	return query.WithContext(ctx).Updates(&search).Error
}

func (s *StoneTypes) Find(ctx context.Context, search StoneTypes) (StoneTypes, bool, error) {
	var out StoneTypes
	return out, out.Id != 0, GetDB().WithContext(ctx).Where(&search).Find(&out).Error
}

func (s *StoneTypes) FindList(ctx context.Context, search StoneTypes, opts ...common.Option) ([]*StoneTypes, error) {
	query := GetDB().WithContext(ctx)
	for _, opt := range opts {
		query = opt(query)
	}

	var out []*StoneTypes
	return out, query.Where(&search).Find(&out).Error
}

func (s *StoneTypes) Delete(ctx context.Context, search StoneTypes, isDelete bool) error {
	if isDelete {
		return GetDB().WithContext(ctx).Where("id = ?", search.Id).Unscoped().Delete(&search).Error
	}
	return GetDB().WithContext(ctx).Where("id = ?", search.Id).Delete(&search).Error
}

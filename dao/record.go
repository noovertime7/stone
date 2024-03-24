package dao

import (
	"context"
	"github.com/noovertime7/stone/dao/common"
	"github.com/noovertime7/stone/dao/gormdatatypes"
	"github.com/noovertime7/stone/runtime"
)

type Record struct {
	Id               int                     `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	StoneName        string                  `json:"stoneName"`
	StoneId          int                     `json:"stoneId"`
	Video            string                  `json:"video"`
	Images           gormdatatypes.JSONSlice `json:"images"`
	Location         string                  `json:"location"`
	Description      string                  `json:"description"`
	Longitude        float64                 `json:"longitude"`
	Latitude         float64                 `json:"latitude"`
	DetailedLocation string                  `json:"detailedlocation"`
	common.CommonModel
}

func (s *Record) TableName() string {
	return "t_record"
}

func (s *Record) PageList(ctx context.Context, params runtime.DataBasePager, ops ...common.Option) ([]*Record, int64, error) {
	var total int64
	query := GetDB().WithContext(ctx).Where("")
	limit := params.GetPageSize()
	offset := limit * (params.GetPage() - 1)

	var list []*Record
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

func (s *Record) Save(ctx context.Context, obj *Record) error {
	return GetDB().WithContext(ctx).Save(obj).Error
}

func (s *Record) Updates(ctx context.Context, opt common.Option, search *Record) error {
	query := opt(GetDB())
	return query.WithContext(ctx).Updates(&search).Error
}

func (s *Record) Find(ctx context.Context, search Record) (Record, bool, error) {
	var out Record
	return out, out.Id != 0, GetDB().WithContext(ctx).Where(&search).Find(&out).Error
}

func (s *Record) FindList(ctx context.Context, search Record, opts ...common.Option) ([]*Record, error) {
	query := GetDB().WithContext(ctx)
	for _, opt := range opts {
		query = opt(query)
	}

	var out []*Record
	return out, query.Where(&search).Find(&out).Error
}

func (s *Record) Delete(ctx context.Context, search Record, isDelete bool) error {
	if isDelete {
		return GetDB().WithContext(ctx).Where("id = ?", search.Id).Unscoped().Delete(&search).Error
	}
	return GetDB().WithContext(ctx).Where("id = ?", search.Id).Delete(&search).Error
}

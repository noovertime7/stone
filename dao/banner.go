package dao

import (
	"context"
	"github.com/noovertime7/stone/dao/common"
)

type Banner struct {
	Id        int    `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	Image     string `json:"image"`
	Url       string `json:"url"`
	SortOrder int    `json:"sortOrder" gorm:"default:0"`
	common.CommonModel
}

func (b *Banner) TableName() string {
	return "t_banner"
}

func (b *Banner) FindList(ctx context.Context) ([]*Banner, error) {
	var out []*Banner
	return out, GetDB().WithContext(ctx).Order("sort_order asc").Find(&out).Error
}

func (b *Banner) Save(ctx context.Context, obj *Banner) error {
	return GetDB().WithContext(ctx).Save(obj).Error
}

func (b *Banner) Delete(ctx context.Context, id int) error {
	return GetDB().WithContext(ctx).Where("id = ?", id).Delete(&Banner{}).Error
}

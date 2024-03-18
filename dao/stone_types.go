package dao

import "github.com/noovertime7/stone/dao/common"

type StoneTypes struct {
	Id     int     `json:"id" gorm:"column:id"`
	Name   string  `json:"name"`
	Stones []Stone `json:"stones" gorm:"-"`
	common.CommonModel
}

func (f *StoneTypes) TableName() string {
	return "t_stone_types"
}

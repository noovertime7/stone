package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/stone/dao"
	"github.com/noovertime7/stone/pkg"
	"gorm.io/gorm"
)

type CreateStoneTypes struct {
	Name string `json:"name" form:"name" validate:"required"`
}

func (p *CreateStoneTypes) BindingValidParams(ctx *gin.Context) error {
	return pkg.DefaultGetValidParams(ctx, p)
}

type PageStoneTypesInput struct {
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword"`   //关键字
}

func (p *PageStoneTypesInput) BindingValidParams(ctx *gin.Context) error {
	return pkg.DefaultGetValidParams(ctx, p)
}

func (p *PageStoneTypesInput) GetPage() int {
	return p.Page
}

func (p *PageStoneTypesInput) GetPageSize() int {
	return p.PageSize
}

func (p *PageStoneTypesInput) IsFitter() bool {
	return p.Keyword != ""
}

func (p *PageStoneTypesInput) Do(tx *gorm.DB) {
	return
	//tx.Where("name like ? or content like ?", "%"+p.Keyword+"%", "%"+p.Keyword+"%")
}

type PageStoneTypesOut struct {
	Total    int64             `json:"total"`
	List     []*dao.StoneTypes `json:"list"`
	Page     int               `json:"page" form:"page"`         // 页码
	PageSize int               `json:"pageSize" form:"pageSize"` // 每页大小
}

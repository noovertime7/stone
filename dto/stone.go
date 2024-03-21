package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/stone/dao"
	"github.com/noovertime7/stone/pkg"
	"gorm.io/gorm"
)

type CreateStone struct {
	StoneTypeId  int      `json:"stoneTypeId"`
	Name         string   `json:"name"`
	CoverImages  []string `json:"coverImages"`
	DetailImages []string `json:"detailImages"`
	Description  string   `json:"description"`
}

func (p *CreateStone) BindingValidParams(ctx *gin.Context) error {
	return pkg.DefaultGetValidParams(ctx, p)
}

type PageStoneInput struct {
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword"`   //关键字
}

func (p *PageStoneInput) BindingValidParams(ctx *gin.Context) error {
	return pkg.DefaultGetValidParams(ctx, p)
}

func (p *PageStoneInput) GetPage() int {
	return p.Page
}

func (p *PageStoneInput) GetPageSize() int {
	return p.PageSize
}

func (p *PageStoneInput) IsFitter() bool {
	return p.Keyword != ""
}

func (p *PageStoneInput) Do(tx *gorm.DB) {
	//return
	tx.Where("name like ?", "%"+p.Keyword+"%")
}

type PageStoneOut struct {
	Total    int64        `json:"total"`
	List     []*dao.Stone `json:"list"`
	Page     int          `json:"page" form:"page"`         // 页码
	PageSize int          `json:"pageSize" form:"pageSize"` // 每页大小
}

package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/stone/dao"
	"github.com/noovertime7/stone/pkg"
	"gorm.io/gorm"
)

type CreateRecordInput struct {
	StoneId          int      `json:"stoneId"`
	Video            string   `json:"video"`
	Images           []string `json:"images"`
	Location         string   `json:"location"`
	Description      string   `json:"description"`
	Longitude        float64  `json:"longitude"`
	Latitude         float64  `json:"latitude"`
	DetailedLocation string   `json:"detailedlocation"`
}

func (p *CreateRecordInput) BindingValidParams(ctx *gin.Context) error {
	return pkg.DefaultGetValidParams(ctx, p)
}

type PageRecordInput struct {
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword"`   //关键字
}

func (p *PageRecordInput) BindingValidParams(ctx *gin.Context) error {
	return pkg.DefaultGetValidParams(ctx, p)
}

func (p *PageRecordInput) GetPage() int {
	return p.Page
}

func (p *PageRecordInput) GetPageSize() int {
	return p.PageSize
}

func (p *PageRecordInput) IsFitter() bool {
	return p.Keyword != ""
}

func (p *PageRecordInput) Do(tx *gorm.DB) {
	//return
	tx.Where("name like ?", "%"+p.Keyword+"%")
}

type PagerRecordOut struct {
	Total    int64         `json:"total"`
	List     []*dao.Record `json:"list"`
	Page     int           `json:"page" form:"page"`         // 页码
	PageSize int           `json:"pageSize" form:"pageSize"` // 每页大小
}

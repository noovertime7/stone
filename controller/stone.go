package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/stone/dto"
	"github.com/noovertime7/stone/middleware"
	"github.com/noovertime7/stone/services"
	"net/http"
)

func StoneApiRegister(router *gin.RouterGroup) {
	c := stoneAPI{}
	router.POST("/stone", c.Create)
	router.GET("/stone/page", c.Page)
}

type stoneAPI struct {
	service services.StoneService
}

func (s *stoneAPI) Create(ctx *gin.Context) {
	params := &dto.CreateStone{}
	if err := params.BindingValidParams(ctx); err != nil {
		middleware.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}
	err := s.service.Save(ctx, params)
	if err != nil {
		middleware.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}
	middleware.ResponseSuccessNoData(ctx)
}

func (s *stoneAPI) Page(ctx *gin.Context) {
	params := &dto.PageStoneInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		middleware.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}
	data, err := s.service.Page(ctx, params)
	if err != nil {
		middleware.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

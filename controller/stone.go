package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/stone/dto"
	"github.com/noovertime7/stone/middleware"
	"github.com/noovertime7/stone/services"
	"net/http"
	"strconv"
)

func StoneApiRegister(router *gin.RouterGroup) {
	c := stoneAPI{}
	router.POST("/stone", c.Create)
	router.POST("/stone/:id/update", middleware.ParamGet(middleware.IDParam), c.Update)
	router.GET("/stone/page", c.Page)
	router.GET("/stone/:id", middleware.ParamGet(middleware.IDParam), c.Get)
	router.GET("/stones/:id/same", middleware.ParamGet(middleware.IDParam), c.FindSameTypeStones)
	router.GET("/stones/:id/bytype", middleware.ParamGet(middleware.IDParam), c.FindStonesByType)
	router.GET("/hotStones", c.FindHotList)
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

func (s *stoneAPI) Update(ctx *gin.Context) {
	stoneTypeId := middleware.GetStringFromCtx(ctx, middleware.IDParam)
	intID, _ := strconv.Atoi(stoneTypeId)
	params := &dto.CreateStone{}
	if err := params.BindingValidParams(ctx); err != nil {
		middleware.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}
	err := s.service.Update(ctx, intID, params)
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

func (s *stoneAPI) FindHotList(ctx *gin.Context) {
	data, err := s.service.FindHotList(ctx.Request.Context())
	if err != nil {
		middleware.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

func (s *stoneAPI) FindSameTypeStones(ctx *gin.Context) {
	stoneTypeId := middleware.GetStringFromCtx(ctx, middleware.IDParam)
	intID, _ := strconv.Atoi(stoneTypeId)
	data, err := s.service.FindSameTypeStones(ctx, intID)
	if err != nil {
		middleware.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

func (s *stoneAPI) FindStonesByType(ctx *gin.Context) {
	stoneTypeId := middleware.GetStringFromCtx(ctx, middleware.IDParam)
	intID, _ := strconv.Atoi(stoneTypeId)
	data, err := s.service.FindStonesByTypeID(ctx, intID)
	if err != nil {
		middleware.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

func (s *stoneAPI) Get(ctx *gin.Context) {
	id := middleware.GetStringFromCtx(ctx, middleware.IDParam)

	intID, _ := strconv.Atoi(id)
	data, err := s.service.Get(ctx, intID)
	if err != nil {
		middleware.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

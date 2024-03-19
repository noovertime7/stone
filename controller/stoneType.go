package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/stone/dto"
	"github.com/noovertime7/stone/middleware"
	"github.com/noovertime7/stone/services"
	"strconv"
)

func StoneTypeRegister(router *gin.RouterGroup) {
	curd := StoneTypeController{
		service: &services.StoneTypeService{},
	}
	router.POST("/stoneType", curd.Save)
	router.GET("/stoneType/:id", middleware.ParamGet(middleware.IDParam), curd.Get)
	router.GET("/stoneTypes", curd.List)
	router.DELETE("/stoneType/:id", middleware.ParamGet(middleware.IDParam), curd.Delete)
	router.PUT("/stoneType/:id/:name", middleware.ParamGet(middleware.IDParam, middleware.NameParam), curd.Update)
}

type StoneTypeController struct {
	service *services.StoneTypeService
}

func (s *StoneTypeController) Get(ctx *gin.Context) {
	id := middleware.GetStringFromCtx(ctx, middleware.IDParam)
	intId, _ := strconv.Atoi(id)
	data, err := s.service.Get(ctx.Request.Context(), intId)
	if err != nil {
		middleware.ResponseError(ctx, middleware.InternalErrorCode, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

func (s *StoneTypeController) Save(ctx *gin.Context) {
	params := &dto.CreateStoneTypes{}
	if err := params.BindingValidParams(ctx); err != nil {
		middleware.ResponseError(ctx, middleware.InvalidRequestErrorCode, err)
		return
	}
	err := s.service.Save(ctx, params.Name)
	if err != nil {
		middleware.ResponseError(ctx, middleware.InternalErrorCode, err)
		return
	}
	middleware.ResponseSuccessNoData(ctx)
}

func (s *StoneTypeController) List(ctx *gin.Context) {
	data, err := s.service.List(ctx.Request.Context())
	if err != nil {
		middleware.ResponseError(ctx, middleware.InternalErrorCode, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

func (s *StoneTypeController) Delete(ctx *gin.Context) {
	id := middleware.GetStringFromCtx(ctx, middleware.IDParam)
	intId, _ := strconv.Atoi(id)
	err := s.service.Delete(ctx, intId)
	if err != nil {
		middleware.ResponseError(ctx, middleware.InternalErrorCode, err)
		return
	}
	middleware.ResponseSuccessNoData(ctx)
}

func (s *StoneTypeController) Update(ctx *gin.Context) {
	id := middleware.GetStringFromCtx(ctx, middleware.IDParam)
	name := middleware.GetStringFromCtx(ctx, middleware.NameParam)

	intId, _ := strconv.Atoi(id)
	err := s.service.Update(ctx, intId, name)
	if err != nil {
		middleware.ResponseError(ctx, middleware.InternalErrorCode, err)
		return
	}
	middleware.ResponseSuccessNoData(ctx)
}

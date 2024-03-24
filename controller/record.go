package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/stone/dto"
	"github.com/noovertime7/stone/middleware"
	"github.com/noovertime7/stone/services"
	"strconv"
)

func RecordeRegister(router *gin.RouterGroup) {
	curd := recordeRegister{
		service: &services.RecordService{},
	}
	router.POST("/record", curd.Save)
	router.GET("/record/:id", middleware.ParamGet(middleware.IDParam), curd.Get)
	router.GET("/record/page", curd.Page)
	router.GET("/record", curd.List)
	router.DELETE("/record/:id", middleware.ParamGet(middleware.IDParam), curd.Delete)
}

type recordeRegister struct {
	service *services.RecordService
}

func (s *recordeRegister) Get(ctx *gin.Context) {
	id := middleware.GetStringFromCtx(ctx, middleware.IDParam)
	intId, _ := strconv.Atoi(id)
	data, err := s.service.Get(ctx.Request.Context(), intId)
	if err != nil {
		middleware.ResponseError(ctx, middleware.InternalErrorCode, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

func (s *recordeRegister) Save(ctx *gin.Context) {
	params := &dto.CreateRecordInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		middleware.ResponseError(ctx, middleware.InvalidRequestErrorCode, err)
		return
	}
	err := s.service.Save(ctx, params)
	if err != nil {
		middleware.ResponseError(ctx, middleware.InternalErrorCode, err)
		return
	}
	middleware.ResponseSuccessNoData(ctx)
}

func (s *recordeRegister) Page(ctx *gin.Context) {
	params := &dto.PageRecordInput{}
	if err := params.BindingValidParams(ctx); err != nil {
		middleware.ResponseError(ctx, middleware.InvalidRequestErrorCode, err)
		return
	}
	data, err := s.service.Page(ctx, params)
	if err != nil {
		middleware.ResponseError(ctx, middleware.InternalErrorCode, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

func (s *recordeRegister) List(ctx *gin.Context) {
	data, err := s.service.List(ctx.Request.Context())
	if err != nil {
		middleware.ResponseError(ctx, middleware.InternalErrorCode, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

func (s *recordeRegister) Delete(ctx *gin.Context) {
	id := middleware.GetStringFromCtx(ctx, middleware.IDParam)
	intId, _ := strconv.Atoi(id)
	err := s.service.Delete(ctx, intId)
	if err != nil {
		middleware.ResponseError(ctx, middleware.InternalErrorCode, err)
		return
	}
	middleware.ResponseSuccessNoData(ctx)
}

//func (s *recordeRegister) Update(ctx *gin.Context) {
//	id := middleware.GetStringFromCtx(ctx, middleware.IDParam)
//	name := middleware.GetStringFromCtx(ctx, middleware.NameParam)
//
//	intId, _ := strconv.Atoi(id)
//	err := s.service.Update(ctx, intId, name)
//	if err != nil {
//		middleware.ResponseError(ctx, middleware.InternalErrorCode, err)
//		return
//	}
//	middleware.ResponseSuccessNoData(ctx)
//}

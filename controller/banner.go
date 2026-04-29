package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/stone/middleware"
	"github.com/noovertime7/stone/services"
	"net/http"
	"strconv"
)

func BannerRegister(router *gin.RouterGroup) {
	c := bannerController{
		service: &services.BannerService{},
	}
	router.GET("/home/banner", c.List)
	router.POST("/banner", c.Create)
	router.DELETE("/banner/:id", middleware.ParamGet(middleware.IDParam), c.Delete)
}

type bannerController struct {
	service *services.BannerService
}

func (b *bannerController) List(ctx *gin.Context) {
	data, err := b.service.List(ctx.Request.Context())
	if err != nil {
		middleware.ResponseError(ctx, middleware.InternalErrorCode, err)
		return
	}
	middleware.ResponseSuccess(ctx, data)
}

type createBannerInput struct {
	Image     string `json:"image" form:"image"`
	Url       string `json:"url" form:"url"`
	SortOrder int    `json:"sortOrder" form:"sortOrder"`
}

func (b *bannerController) Create(ctx *gin.Context) {
	var params createBannerInput
	if err := ctx.ShouldBind(&params); err != nil {
		middleware.ResponseError(ctx, http.StatusBadRequest, err)
		return
	}
	err := b.service.Save(ctx.Request.Context(), params.Image, params.Url, params.SortOrder)
	if err != nil {
		middleware.ResponseError(ctx, middleware.InternalErrorCode, err)
		return
	}
	middleware.ResponseSuccessNoData(ctx)
}

func (b *bannerController) Delete(ctx *gin.Context) {
	id := middleware.GetStringFromCtx(ctx, middleware.IDParam)
	intId, _ := strconv.Atoi(id)
	err := b.service.Delete(ctx.Request.Context(), intId)
	if err != nil {
		middleware.ResponseError(ctx, middleware.InternalErrorCode, err)
		return
	}
	middleware.ResponseSuccessNoData(ctx)
}

package controller

import (
	"github.com/e421083458/golang_common/log"
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/stone/middleware"
	"github.com/noovertime7/stone/utils"
	"net/http"
)

func SystemApiRegister(router *gin.RouterGroup) {
	c := systemApiController{}
	router.POST("/upload", c.Upload)
}

type systemApiController struct {
}

func (s *systemApiController) Upload(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		middleware.ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}
	files := form.File["file"]

	links := []string{}
	for _, file := range files {
		if err := utils.UploadFile(ctx, file); err != nil {
			log.Error("上传失败: %v", err)
			middleware.ResponseError(ctx, http.StatusInternalServerError, err)
			return
		}

		links = append(links, utils.BuildFileLink(file))
	}

	middleware.ResponseSuccess(ctx, links)
}

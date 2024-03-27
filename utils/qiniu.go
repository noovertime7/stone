package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
	"os"
)

func UploadFile(ctx *gin.Context, file *multipart.FileHeader) error {
	err := ctx.SaveUploadedFile(file, file.Filename)
	if err != nil {
		return fmt.Errorf("保存到临时文件失败:%w", err)
	}
	ak := "6mJpijuYjeCqbCeLEE7XI7LFtPLEc4lr2j18nz8Z"
	sk := "mB6kb0EKo1rLIhSo7447M0CI0GufyyRJwKmShysM"
	if IsStrEmpty(sk) || IsStrEmpty(ak) {
		return fmt.Errorf("七牛云配置错误")
	}

	// 生成上传凭证
	mac := qbox.NewMac(ak, sk)
	putPolicy := storage.PutPolicy{
		Scope: buildFilePath(file), // 指定文件夹路径
	}
	upToken := putPolicy.UploadToken(mac)
	// 构建上传配置
	cfg := storage.Config{}
	// 空间对应的机房
	// 使用指定的区域
	cfg.Zone = &storage.Zone{
		SrcUpHosts: []string{"up-z1.qiniup.com"},
		RsHost:     "rs-z1.qiniu.com",
	}
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	// 可选配置
	putExtra := storage.PutExtra{}

	err = formUploader.PutFile(ctx.Request.Context(), nil, upToken, buildKey(file), file.Filename, &putExtra)
	if err != nil {
		clean(file)
		return err
	}
	return clean(file)
}

func buildFilePath(file *multipart.FileHeader) string {
	// 上传到七牛后保存的文件名
	return fmt.Sprintf("%s:%s", "chentengspace", "stone/"+file.Filename)
}

func buildKey(file *multipart.FileHeader) string {
	return "stone/" + file.Filename
}

func BuildImageFileLink(file *multipart.FileHeader) string {
	return fmt.Sprintf("http://qiniu.yunxue521.top/stone/%s%s", file.Filename, BuildWaterMarkUrl())
}

func BuildFileLink(file *multipart.FileHeader) string {
	return fmt.Sprintf("http://qiniu.yunxue521.top/stone/%s", file.Filename)
}

func clean(file *multipart.FileHeader) error {
	// 上传成功后删除本地文件
	err := os.Remove(file.Filename)
	if err != nil {
		return fmt.Errorf("删除本地文件失败:%w", err)
	}
	return nil
}

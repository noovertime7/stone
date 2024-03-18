package controller

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/stone/dto"
	"github.com/noovertime7/stone/middleware"
	"github.com/noovertime7/stone/services"
)

type ApiController struct {
}

func ApiRegister(router *gin.RouterGroup) {
	curd := ApiController{}
	router.POST("/users/login", curd.Login)
	router.GET("/users/info", curd.Info)
	router.GET("/loginout", curd.LoginOut)
}

func ApiLoginRegister(router *gin.RouterGroup) {
	//curd := ApiController{}
	//router.GET("/user/listpage", curd.ListPage)
	//router.POST("/user/add", curd.AddUser)
	//router.POST("/user/edit", curd.EditUser)
	//router.POST("/user/remove", curd.RemoveUser)
	//router.POST("/user/batchremove", curd.RemoveUser)
}

func (demo *ApiController) Login(c *gin.Context) {
	api := &dto.LoginInput{}
	if err := api.BindingValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	s := services.ApiService{}
	data, err := s.Login(c, api)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	middleware.ResponseSuccess(c, data)

	return
}

type info struct {
	UserName string   `json:"username"`
	Roles    []string `json:"roles"`
}

func (demo *ApiController) Info(c *gin.Context) {
	adminUser := &info{UserName: "admin", Roles: []string{"admin"}}

	middleware.ResponseSuccess(c, adminUser)

}

func (demo *ApiController) LoginOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("user")
	session.Save()
	return
}

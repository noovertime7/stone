package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/stone/dao"
	"github.com/noovertime7/stone/dto"
)

type ApiService struct {
}

func (as *ApiService) Login(c *gin.Context, input *dto.LoginInput) (*dto.LoginResult, error) {
	model := dao.User{}
	u, err := model.FindByUserName(c, input.Username)
	if err != nil {
		return nil, err
	}
	if u.Password != input.Password {
		return nil, fmt.Errorf("用户名或密码错误！")
	}
	return &dto.LoginResult{
		ID:       u.Id,
		Avatar:   u.Avatar,
		Account:  u.Account,
		Nickname: u.Nickname,
		Mobile:   u.Mobile,
		Token:    "2bbc8896-a3f5-47cf-a4b7-c997426ca8a4",
	}, nil
}

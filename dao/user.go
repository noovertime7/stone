package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/noovertime7/stone/dao/common"
)

type ListPageOutput struct {
	List  []User `form:"list" json:"list" comment:"用户列表" validate:""`
	Total int64  `form:"page" json:"page" comment:"用户总数" validate:"required"`
}

type User struct {
	Id       int    `json:"id" gorm:"column:id"`
	Account  string `json:"account"`
	Password string `json:"password"`
	Nickname string `json:"nickname,omitempty"`
	Avatar   string `json:"avatar"`
	Mobile   string `json:"mobile"`
	common.CommonModel
}

func (f *User) TableName() string {
	return "t_user"
}

func (f *User) Del(c *gin.Context, idSlice []string) error {
	err := GetDB().WithContext(c).Where("id in (?)", idSlice).Delete(&User{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (f *User) Find(c *gin.Context, id int64) (*User, error) {
	var user *User
	err := GetDB().WithContext(c).Where("id = ?", id).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (f *User) FindByUserName(c *gin.Context, name string) (*User, error) {
	var user *User
	err := GetDB().WithContext(c).Where("account = ?", name).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (f *User) Save(c *gin.Context) error {
	if err := GetDB().WithContext(c).Save(f).Error; err != nil {
		return err
	}
	return nil
}

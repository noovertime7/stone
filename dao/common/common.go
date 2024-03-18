package common

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type CommonModel struct {
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type SqlAction string

const (
	Equal        SqlAction = "="
	NotEqual     SqlAction = "!="
	LessThan     SqlAction = "<"
	LessEqual    SqlAction = "<="
	GreaterThan  SqlAction = ">"
	GreaterEqual SqlAction = ">="
)

// Option 用于将更新参数的能力给到service层
type Option func(db *gorm.DB) *gorm.DB

func WithField(field string, action SqlAction, value string) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%s %s '%s'", field, action, value))
	}
}

func WithInstanceOption(instanceID string, action SqlAction) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("instanceID %s '%s'", action, instanceID))
	}
}

func WithTenantOption(tenantID string, action SqlAction) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("tenant_id %s '%s'", action, tenantID))
	}
}

func WithRegionOption(region string, action SqlAction) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("region %s '%s'", action, region))
	}
}

func WithBetweenTime(start, end time.Time, filed string) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%s BETWEEN '%s' AND '%s'", filed, start.String(), end.String()))
	}
}

func WithJsonMapFiled(filed, key, value string, action SqlAction) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%s ->'$.%s' %s '%s' ", filed, key, action, value))
	}
}

func WithJsonSliceCONTAINS(filed, value string) Option {
	return func(db *gorm.DB) *gorm.DB {
		value = fmt.Sprintf("\"%s\"", value)
		return db.Where(fmt.Sprintf("JSON_CONTAINS(%s, '%s')", filed, value))
	}
}

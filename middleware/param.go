package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/martian/log"
	"github.com/noovertime7/stone/utils"
)

type ParamKey string

const (
	ClusterParam    ParamKey = "cluster"
	NamespaceParam  ParamKey = "namespace"
	NameParam       ParamKey = "name"
	InstanceIDParam ParamKey = "instanceID"
	ContainerParam  ParamKey = "container"
	IDParam         ParamKey = "id"
	VersionParam    ParamKey = "version"
)

func (p ParamKey) String() string {
	return string(p)
}

func GetStringFromCtx(ctx *gin.Context, key ParamKey) string {
	val, exist := ctx.Get(key.String())
	if exist {
		return val.(string)
	} else {
		return ""
	}
}

func ParamGet(keys ...ParamKey) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		for _, key := range keys {
			val := ctx.Param(key.String())
			if utils.IsStrEmpty(val) {
				log.Debugf("get [%s] from param empty", key.String())
				ResponseError(ctx, 400, fmt.Errorf("bad request,[%s] empty", key.String()))
				ctx.Abort()
			}
			ctx.Set(key.String(), val)
		}
		ctx.Next()
	}
}

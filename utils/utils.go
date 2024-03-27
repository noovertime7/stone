package utils

import (
	"fmt"
	"github.com/noovertime7/stone/pkg"
)

func IsStrEmpty(str string) bool {
	return str == ""
}

func BuildWaterMarkUrl() string {
	dissolve := 90 // 透明度
	resize := 5    // 缩放
	return fmt.Sprintf("?watermark/4/text/%s/fontsize/400/fill/Z3JheQ==/dissolve/%d/rotate/30/uw/180/uh/180/resize/%d&=", pkg.URLEncodeBase64(pkg.Watermark), dissolve, resize)
}

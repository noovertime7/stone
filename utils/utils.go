package utils

import (
	"fmt"
	"github.com/noovertime7/stone/pkg"
)

func IsStrEmpty(str string) bool {
	return str == ""
}

func BuildWaterMarkUrl() string {
	fontsize := 1000
	fill := "I0ZGRkZGRg"
	return fmt.Sprintf("?watermark/2/text/%s/fontsize/%d/fill/%s", pkg.URLEncodeBase64(pkg.Watermark), fontsize, fill)
}

package pkg

import (
	"encoding/base64"
	"strings"
)

// URLEncodeBase64 对输入字符串进行URL安全的Base64编码
func URLEncodeBase64(input string) string {
	// 使用标准库中的base64编码函数对输入进行编码
	encoded := base64.URLEncoding.EncodeToString([]byte(input))
	// 替换字符+为-，/为_
	encoded = strings.ReplaceAll(encoded, "+", "-")
	encoded = strings.ReplaceAll(encoded, "/", "_")
	// 移除末尾可能存在的=
	encoded = strings.TrimRight(encoded, "=")
	return encoded
}

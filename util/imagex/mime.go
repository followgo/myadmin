package imagex

import (
	"mime"
)

// ExtensionsByMIMEType 根据 MIME 类型获取文件扩展名
func ExtensionsByMIMEType(typ string) ([]string, error) {
	return mime.ExtensionsByType(typ)
}

// MIMETypeByExtension 通过扩展名获取 MIME 类型
func MIMETypeByExtension(ext string) string {
	return mime.TypeByExtension(ext)
}

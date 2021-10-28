package wimg

import (
	"bytes"
	"image"
	"net/http"
)

var (
	Util = util{}
)

type util struct {
	img     image.Image
	imgType string
	saveBuf bytes.Buffer
}

// FileContentType 获取文件类型
func (u *util) FileContentType(buffer []byte) string {
	contentType := http.DetectContentType(buffer)
	return contentType
}

package imagex

import (
	"bytes"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"strings"

	"github.com/chai2010/webp"

	"github.com/followgo/myadmin/util/errors"
)

// ConvertFormat 转换图片格式
// 支持的源图片格式：webp, jpg, png, gif
// 支持的目标图片格式: jpg, webp
func ConvertFormat(src io.Reader, srcMIMEType, dstMIMEType string, quality float32) (*bytes.Buffer, error) {
	var m image.Image
	var err error

	// 读取源文件
	switch strings.ToLower(srcMIMEType) {
	case "image/webp":
		m, err = webp.Decode(src)
		if err != nil {
			return nil, errors.Wrap(err, "不能解码的 image/webp 文件")
		}

	case "image/jpeg":
		m, err = jpeg.Decode(src)
		if err != nil {
			return nil, errors.Wrap(err, "不能解码的 image/jpeg 文件")
		}

	case "image/png":
		m, err = png.Decode(src)
		if err != nil {
			return nil, errors.Wrap(err, "不能解码的 image/png 文件")
		}

	case "image/gif":
		g, err := gif.DecodeAll(src)
		if err != nil {
			return nil, errors.Wrap(err, "不能解码的 image/gif 文件")
		}

		if g.Image != nil && len(g.Image) > 0 {
			m = g.Image[0]
		} else {
			return nil, errors.Wrap(err, "空 image/gif 文件")
		}

	default:
		return nil, errors.New("不支持的源文件格式")
	}

	// 目标文件
	var buf = bytes.NewBuffer(nil)
	switch strings.ToLower(dstMIMEType) {
	case "image/webp":
		err = webp.Encode(buf, m, &webp.Options{Lossless: false, Quality: quality})

	case "image/jpeg":
		err = jpeg.Encode(buf, m, &jpeg.Options{Quality: int(quality)})

	default:
		return nil, errors.New("不支持的目标文件格式")
	}

	return buf, err
}

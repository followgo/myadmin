package imagex

import (
	"bytes"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"
	"strings"

	"github.com/chai2010/webp"

	"github.com/followgo/myadmin/util/errors"
)

// WaterMark 图片打水印
func WaterMark(src io.Reader, mark image.Image, mimeType string, quality float32) (io.Reader, error) {
	var img image.Image
	var err error

	// 读取源文件
	switch strings.ToLower(mimeType) {
	case "image/webp":
		img, err = webp.Decode(src)
		if err != nil {
			return nil, errors.Wrap(err, "不能解码的 image/webp 文件")
		}

	case "image/jpeg":
		img, err = jpeg.Decode(src)
		if err != nil {
			return nil, errors.Wrap(err, "不能解码的 image/jpeg 文件")
		}

	case "image/png":
		img, err = png.Decode(src)
		if err != nil {
			return nil, errors.Wrap(err, "不能解码的 image/png 文件")
		}

	default:
		return nil, errors.New("不支持的文件格式")
	}

	// 水印位置：右下角，并偏移 10 像素
	imgRec := img.Bounds()
	markRec := mark.Bounds()
	offset := image.Pt(imgRec.Dx()-markRec.Dx()-10, imgRec.Dy()-markRec.Dy()-10)

	// 画图
	newImg := image.NewNRGBA(imgRec)
	draw.Draw(newImg, imgRec, img, image.Point{}, draw.Src)
	draw.Draw(newImg, markRec.Add(offset), mark, image.Point{}, draw.Over)

	// 写入 buffer
	var buf = bytes.NewBuffer(nil)
	switch strings.ToLower(mimeType) {
	case "image/webp":
		err = webp.Encode(buf, newImg, &webp.Options{Lossless: true, Quality: quality})

	case "image/jpeg":
		err = jpeg.Encode(buf, newImg, &jpeg.Options{Quality: int(quality)})

	case "image/png":
		err = png.Encode(buf, newImg)

	default:
		return nil, errors.New("不支持的文件格式")
	}

	return buf, err
}

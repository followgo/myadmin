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
	"github.com/nfnt/resize"

	"github.com/followgo/myadmin/util/errors"
)

// Resize 调整图片尺寸
// 支持的图片格式：webp, jpg, png, gif
func Resize(src io.Reader, mimeType string, maxWidth, maxHeight uint, quality float32) (*bytes.Buffer, error) {
	switch strings.ToLower(mimeType) {
	case "image/webp":
		m, err := webp.Decode(src)
		if err != nil {
			return nil, errors.Wrap(err, "不能解码的 image/webp 文件")
		}

		m = resize.Thumbnail(maxWidth, maxHeight, m, resize.Lanczos2)

		buf := bytes.NewBuffer(nil)
		err = webp.Encode(buf, m, &webp.Options{Lossless: false, Quality: quality})
		return buf, err

	case "image/jpeg":
		m, err := jpeg.Decode(src)
		if err != nil {
			return nil, errors.Wrap(err, "不能解码的 image/jpeg 文件")
		}

		m = resize.Thumbnail(maxWidth, maxHeight, m, resize.Lanczos2)

		buf := bytes.NewBuffer(nil)
		err = jpeg.Encode(buf, m, &jpeg.Options{Quality: int(quality)})
		return buf, err

	case "image/png":
		m, err := png.Decode(src)
		if err != nil {
			return nil, errors.Wrap(err, "不能解码的 image/png 文件")
		}

		m = resize.Thumbnail(maxWidth, maxHeight, m, resize.Lanczos2)

		buf := bytes.NewBuffer(nil)
		err = png.Encode(buf, m)
		return buf, err

	case "image/gif":
		g, err := gif.DecodeAll(src)
		if err != nil {
			return nil, errors.Wrap(err, "不能解码的 image/gif 文件")
		}

		for _, p := range g.Image {
			m := image.Image(p)
			m = resize.Thumbnail(maxWidth, maxHeight, m, resize.Lanczos2)
		}

		buf := bytes.NewBuffer(nil)
		err = gif.EncodeAll(buf, g)
		return buf, err

	default:
		return nil, errors.New("不支持的文件格式")
	}
}

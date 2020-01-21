package apiv1

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"

	. "github.com/followgo/myadmin/config"
	"github.com/followgo/myadmin/model"
	"github.com/followgo/myadmin/util"
	"github.com/followgo/myadmin/util/imagex"
)

type FileAPI struct{}

// Upload 上传文件
func (api *FileAPI) Upload(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Internal: err}
	}

	src, err := file.Open()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "不能打开文件", Internal: err}
	}
	defer src.Close()
	mimeType := file.Header.Get("Content-Type")

	// 判断文件类型是否允许
	if !util.HasStringSlice(mimeType, Cfg.Upload.AllowMIMETypes, false) {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "不允许上传此类型的文件"}
	}
	if file.Size > (Cfg.Upload.AllowMaxSizeMB << 20) {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "文件太大了"}
	}

	// 非图片文件
	if !util.HasStringSlice(mimeType, []string{"image/webp", "image/png", "image/jpeg", "image/gif"}, false) {
		// 文件hash
		fileHash := util.Hash(src, []byte(Cfg.SecuritySalt))
		_, _ = src.Seek(0, io.SeekStart)

		// 查找数据库
		modelFile := &model.File{Hash: fileHash}
		if has, err := modelFile.Get(); err != nil {
			return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "读数据库出错", Internal: err}
		} else if has {
			// 文件已经存在
			return c.JSON(http.StatusCreated, modelFile)
		}

		// 根据 MIMEType 使用文件扩展名
		extName, err := imagex.ExtensionsByMIMEType(mimeType)
		if err != nil || len(extName) == 0 {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: "未知的 MIME 类型"}
		}

		// 保存文件
		dstFilename := filepath.Join(Cfg.Upload.Directory, fileHash+extName[0])
		dst, err := os.OpenFile(dstFilename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "创建文件", Internal: err}
		}
		defer dst.Close()

		if _, err := io.Copy(dst, src); err != nil {
			return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "保存文件", Internal: err}
		}

		modelFile.Filename = fileHash + extName[0]
		modelFile.Size = file.Size
		modelFile.Hash = fileHash
		modelFile.MIMEType = mimeType
		if ok, err := modelFile.Insert(); err != nil {
			return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "插入数据出错", Internal: err}
		} else if !ok {
			return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "插入数据失败"}
		}

		return c.JSON(http.StatusCreated, modelFile)
	}

	// 图片文件
	var buf *bytes.Buffer

	// 优化图片
	buf, err = imagex.Resize(src, mimeType, 1200, 9999, 90) // 不限制长图
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "优化图片大小出错", Internal: err}
	}

	// 尝试转换图片格式为 webp
	if Cfg.Upload.ConvertPictureToWebp {
		buf1, err := imagex.ConvertFormat(buf, mimeType, "image/webp", 90)
		if err == nil {
			buf = buf1
			mimeType = "image/webp"
		}
	}

	// 文件hash
	fileHash := util.Hash(bytes.NewReader(buf.Bytes()), []byte(Cfg.SecuritySalt))

	// 查找数据库
	modelFile := &model.File{Hash: fileHash}
	if has, err := modelFile.Get(); err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "读数据库出错", Internal: err}
	} else if has {
		// 文件已经存在
		return c.JSON(http.StatusCreated, modelFile)
	}

	// 根据 MIMEType 使用文件扩展名
	extName, err := imagex.ExtensionsByMIMEType(mimeType)
	if err != nil || len(extName) == 0 {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "未知的 MIME 类型"}
	}

	// 保存文件
	dstFilename := filepath.Join(Cfg.Upload.Directory, fileHash+extName[0])
	dst, err := os.OpenFile(dstFilename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "创建文件", Internal: err}
	}
	defer dst.Close()

	size, err := io.Copy(dst, buf)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "保存文件", Internal: err}
	}

	modelFile.Filename = fileHash + extName[0]
	modelFile.Size = size
	modelFile.Hash = fileHash
	modelFile.MIMEType = mimeType
	if ok, err := modelFile.Insert(); err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "插入数据出错", Internal: err}
	} else if !ok {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "插入数据失败"}
	}

	return c.JSON(http.StatusCreated, modelFile)
}

func (api *FileAPI) UploadHTML(c echo.Context) error {
	return c.HTML(200, `<!doctype html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <title>Single file upload</title>
</head>
<body>
<h1>Upload single file with fields</h1>

<form action="/api/v1/upload" method="post" enctype="multipart/form-data">
    Name: <input type="text" name="name"><br>
    Email: <input type="email" name="email"><br>
    Files: <input type="file" name="file"><br><br>
    <input type="submit" value="Submit">
</form>
</body>
</html>`)
}

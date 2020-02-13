package apiv1

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/labstack/echo/v4"

	. "github.com/followgo/myadmin/config"
	"github.com/followgo/myadmin/model"
	"github.com/followgo/myadmin/module/orm"
	"github.com/followgo/myadmin/util"
	"github.com/followgo/myadmin/util/imagex"
)

type FileAPI struct{}

// Create 上传文件
func (api *FileAPI) Create(c echo.Context) error {
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
		extNames, err := imagex.ExtensionsByMIMEType(mimeType)
		if err != nil || len(extNames) == 0 {
			return &echo.HTTPError{Code: http.StatusBadRequest, Message: "未知的 MIME 类型"}
		}

		// 保存文件
		dstFilename := filepath.Join(Cfg.Upload.Directory, fileHash+extNames[0])
		dst, err := os.OpenFile(dstFilename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "创建文件", Internal: err}
		}
		defer dst.Close()

		if _, err := io.Copy(dst, src); err != nil {
			return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "保存文件", Internal: err}
		}

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
	buf, err = imagex.Resize(src, mimeType, 1200, 0, 90) // 不限制长图
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
	extNames, err := imagex.ExtensionsByMIMEType(mimeType)
	if err != nil || len(extNames) == 0 {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "未知的 MIME 类型"}
	}

	// 保存文件
	dstFilename := filepath.Join(Cfg.Upload.Directory, fileHash+extNames[0])
	dst, err := os.OpenFile(dstFilename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "创建文件", Internal: err}
	}
	defer dst.Close()

	size, err := io.Copy(dst, buf)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "保存文件", Internal: err}
	}

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

// Get 下载文件
func (api *FileAPI) Get(c echo.Context) error {
	file := model.File{UUID: c.Param("uuid")}
	has, err := file.Get()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Internal: err}
	} else if !has {
		return &echo.HTTPError{Code: http.StatusNotFound, Message: "无此文件信息"}
	}

	// 根据 MIMEType 获取文件扩展名
	extNames, err := imagex.ExtensionsByMIMEType(file.MIMEType)
	if err != nil || len(extNames) == 0 {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "未知的 MIME 类型"}
	}

	filePth := filepath.Join(Cfg.Upload.Directory, file.Hash+extNames[0])
	if has, err := util.HasFile(filePth); err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "读取文件发生错误", Internal: err}
	} else if !has {
		return &echo.HTTPError{Code: http.StatusNotFound, Message: "无此资源"}
	}

	return c.File(filePth)
}

// GetImage 下载图片，支持调整图片尺寸
func (api *FileAPI) GetImage(c echo.Context) error {
	file := model.File{UUID: c.Param("uuid")}
	has, err := file.Get()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Internal: err}
	} else if !has {
		return &echo.HTTPError{Code: http.StatusNotFound, Message: "无此文件信息"}
	}

	// 根据 MIMEType 使用文件扩展名
	extNames, err := imagex.ExtensionsByMIMEType(file.MIMEType)
	if err != nil || len(extNames) == 0 {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "未知的 MIME 类型"}
	}

	baseFilePth := filepath.Join(Cfg.Upload.Directory, file.Hash+extNames[0])
	if has, err := util.HasFile(baseFilePth); err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "读取文件发生错误", Internal: err}
	} else if !has {
		return &echo.HTTPError{Code: http.StatusNotFound, Message: "无此资源"}
	}

	maxWidth, _ := strconv.Atoi(c.QueryParam("max_width"))
	maxHeight, _ := strconv.Atoi(c.QueryParam("max_height"))

	// 缩量图
	if maxWidth > 0 || maxHeight > 0 {
		thumbnailFilename := filepath.Join(Cfg.Upload.Directory, fmt.Sprintf("%s_%d_%d%s", file.Hash, maxWidth, maxHeight, extNames[0]))
		if has, err := util.HasFile(thumbnailFilename); err != nil {
			return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "读取文件发生错误", Internal: err}
		} else if has { // 文件存在
			return c.File(thumbnailFilename)
		}

		// 创建缩略图
		f, err := os.OpenFile(baseFilePth, os.O_RDONLY, 0)
		if err != nil {
			return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "打开图片文件发生错误", Internal: err}
		}
		defer f.Close()

		buf, err := imagex.Resize(f, file.MIMEType, uint(maxWidth), uint(maxHeight), 80)
		if err != nil {
			return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "调整图片尺寸", Internal: err}
		}

		if err := ioutil.WriteFile(thumbnailFilename, buf.Bytes(), 0644); err != nil {
			return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "创建缩略图文件", Internal: err}
		}

		return c.File(thumbnailFilename)
	}

	return c.File(baseFilePth)
}

// Select 列出所有选择的对象
func (api *FileAPI) Select(c echo.Context) error {
	filter := new(orm.Filter)
	if err := c.Bind(filter); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "参数错误", Internal: err}
	}

	files, err := new(model.File).Find(filter)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "读取数据出错", Internal: err}
	}

	total, err := new(model.File).Count(filter)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "读取数据出错", Internal: err}
	}

	return c.JSON(http.StatusOK, echo.Map{"total": total, "data": files})
}

// Delete 删除一个对象
func (api *FileAPI) Delete(c echo.Context) error {
	file := model.File{UUID: c.Param("uuid")}
	has, err := file.Get()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Internal: err}
	} else if !has {
		return &echo.HTTPError{Code: http.StatusNotFound, Message: "无此文件信息"}
	}

	// 根据 MIMEType 获取文件扩展名
	extNames, err := imagex.ExtensionsByMIMEType(file.MIMEType)
	if err != nil || len(extNames) == 0 {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "未知的 MIME 类型"}
	}
	rmErr := os.Remove(filepath.Join(Cfg.Upload.Directory, file.Hash+extNames[0]))

	ok, err := file.Del()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "删除数据出错", Internal: err}
	}
	if !ok {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "删除数据失败"}
	}

	// 删除文件的过程中出现错误
	if rmErr != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "删除文件出错", Internal: rmErr}
	}

	return c.NoContent(http.StatusNoContent)

}

// Update 完全更新一个对象
func (api *FileAPI) Update(c echo.Context) error { return echo.ErrNotFound }

// Patch 修改一个对象的属性
func (api *FileAPI) Patch(c echo.Context) error { return echo.ErrNotFound }

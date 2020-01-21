package util

import (
	"encoding/hex"
	"io"
	"os"
	"reflect"
	"strings"

	"golang.org/x/crypto/blake2b"

	"github.com/followgo/myadmin/util/errors"
)

// StructToStruct 结构体对结构体进行赋值
func StructToStruct(dstPtr, src interface{}) error {
	dstValue := reflect.ValueOf(dstPtr)
	if dstValue.Kind() != reflect.Ptr {
		return errors.New("dst 不是一个指针")
	}
	dstValue = dstValue.Elem()
	if dstValue.Kind() != reflect.Struct {
		return errors.New("dst 不是一个指向结构体的指针")
	}

	srcValue := reflect.ValueOf(src)
	if srcValue.Kind() == reflect.Ptr {
		srcValue = srcValue.Elem()
	}
	if srcValue.Kind() != reflect.Struct {
		return errors.New("src 不是一个结构体")
	}
	srcType := srcValue.Type()

	// 遍历赋值
	for i := 0; i < srcType.NumField(); i++ {
		field := srcType.Field(i)

		// 确认 dst 的 Field 有效，类型一致，并且是可以设置的
		if v := dstValue.FieldByName(field.Name); v.IsValid() && v.CanSet() {
			if v.Type() == field.Type {
				v.Set(srcValue.FieldByName(field.Name))
			}
		}
	}

	return nil
}

// HasFile 判断文件存在
func HasFile(file string) (bool, error) {
	fi, err := os.Stat(file)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, err
	}

	return !fi.IsDir(), nil
}

// HasDir 判断目录存在
func HasDir(file string) (bool, error) {
	fi, err := os.Stat(file)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, err
	}

	return fi.IsDir(), nil
}

// Hash 哈希
func Hash(src io.Reader, salt []byte) string {
	h, _ := blake2b.New256(salt)
	h.Write([]byte("fO1HX6qlkNA7bXk3DM1SDp4L"))
	_, _ = io.Copy(h, src)
	return hex.EncodeToString(h.Sum(nil))
}

// HasStringSlice 判断字符串切片中是否包含某个字符串，带匹配大小写开关
func HasStringSlice(s string, ss []string, matchCase bool) bool {
	if ss == nil || len(ss) == 0 {
		return false
	}

	if matchCase {
		s = strings.ToLower(s)
	}

	for _, v := range ss {
		if matchCase {
			v = strings.ToLower(v)
		}

		if v == s {
			return true
		}
	}
	return false
}

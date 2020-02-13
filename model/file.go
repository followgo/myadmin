package model

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/followgo/myadmin/module/orm"
)

type File struct {
	UUID     string    `xorm:"varchar(36) pk 'uuid'" json:"uuid"`
	MIMEType string    `xorm:"varchar(64) 'mime_type'" json:"mime_type"`
	Hash     string    `xorm:"varchar(255) notnull unique" json:"hash"`
	Size     int64     `json:"size"`
	Created  time.Time `xorm:"created" json:"created"`
}

// TableName 定义数据库表名
func (f *File) TableName() string { return "files" }

// Get 根据非 nil 字段获取一条记录
func (f *File) Get() (has bool, err error) {
	has, err = orm.NewSession(nil).Get(f)
	return
}

// Find 查询多条数据
func (f *File) Find(filter *orm.Filter) (files []File, err error) {
	s := orm.NewSession(filter)

	files = make([]File, 0, 100)
	if err := s.Find(&files); err != nil {
		return nil, err
	}
	return files, err
}

// Count 统计数量
func (f *File) Count(filter *orm.Filter) (n int64, err error) {
	s := orm.NewSession(filter)
	return s.Count(new(File))
}

// Insert 插入一条记录
func (f *File) Insert() (ok bool, err error) {
	f.UUID = uuid.NewV1().String()
	n, err := orm.NewSession(nil).InsertOne(f)
	return n != 0, err
}

// Del 根据uuid删除一条记录
func (f *File) Del() (ok bool, err error) {
	n, err := orm.NewSession(&orm.Filter{Query: "uuid=?", QueryArgs: []interface{}{f.UUID}}).Delete(new(File))
	return n != 0, err
}

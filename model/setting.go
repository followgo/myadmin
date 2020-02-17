package model

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/followgo/myadmin/module/orm"
)

// Setting 网站设置
type Setting struct {
	UUID    string    `xorm:"varchar(36) pk 'uuid'" json:"uuid"`
	Name    string    `xorm:"varchar(32) unique" json:"name"`
	Value   string    `xorm:"varchar(255)" json:"value"`
	Version uint      `xorm:"version" json:"version"`
	Created time.Time `xorm:"created" json:"created"`
	Updated time.Time `xorm:"updated" json:"updated"`
}

// TableName 定义数据库表名
func (s *Setting) TableName() string { return "settings" }

// Get 根据非 nil 字段获取一条记录
func (s *Setting) Get() (has bool, err error) {
	return orm.NewSession(nil).Get(s)
}

// Count 统计数量
func (s *Setting) Count(filter *orm.Filter) (n int64, err error) {
	return orm.NewSession(filter).Count(new(Setting))
}

// Insert 插入一条记录
func (s *Setting) Insert() (ok bool, err error) {
	s.UUID = uuid.NewV1().String()
	n, err := orm.NewSession(nil).InsertOne(s)
	return n != 0, err
}

// Update 更新记录
func (s *Setting) Update() (n int64, err error) {
	n, err = orm.NewSession(&orm.Filter{
		Cols:      []string{"name", "value"},
		Query:     "uuid=?",
		QueryArgs: []interface{}{s.UUID},
	}).Update(s)

	return n, err
}

// Del 根据uuid删除一条记录
func (s *Setting) Del() (ok bool, err error) {
	n, err := orm.NewSession(&orm.Filter{Query: "uuid=?", QueryArgs: []interface{}{s.UUID}}).Delete(new(Setting))
	return n != 0, err
}

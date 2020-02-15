package model

import (
	"time"

	"github.com/followgo/myadmin/module/orm"
)

// Part 内容部件
type Part struct {
	Name      string    `xorm:"varchar(36) pk" json:"name"`
	ContentMd string    `xorm:"text" json:"content_md"`
	Version   uint      `xorm:"version" json:"version"`
	Created   time.Time `xorm:"created" json:"created"`
	Updated   time.Time `xorm:"updated" json:"updated"`
}

// TableName 定义数据库表名
func (p *Part) TableName() string { return "parts" }

// Get 根据非 nil 字段获取一条记录
func (p *Part) Get() (has bool, err error) {
	return orm.NewSession(nil).Get(p)
}

// Count 统计数量
func (p *Part) Count(filter *orm.Filter) (n int64, err error) {
	s := orm.NewSession(filter)
	return s.Count(new(Banner))
}

// Insert 插入一条记录
func (p *Part) Insert() (ok bool, err error) {
	n, err := orm.NewSession(nil).InsertOne(p)
	return n != 0, err
}

// Update 更新记录
func (p *Part) Update() (n int64, err error) {
	n, err = orm.NewSession(&orm.Filter{
		Cols:      []string{"content_raw"},
		Query:     "name=?",
		QueryArgs: []interface{}{p.Name},
	}).Update(p)

	return n, err
}

// Del 根据uuid删除一条记录
func (p *Part) Del() (ok bool, err error) {
	n, err := orm.NewSession(&orm.Filter{Query: "name=?", QueryArgs: []interface{}{p.Name}}).Delete(new(Part))
	return n != 0, err
}

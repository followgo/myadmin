package model

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/followgo/myadmin/module/orm"
)

// Banner 页面横幅
type Banner struct {
	UUID        string    `xorm:"varchar(36) pk 'uuid'" json:"uuid"`
	ImageUUID   string    `xorm:"varchar(36) notnull 'image_uuid'" json:"image_uuid"` // 图片
	Title       string    `xorm:"varchar(24)" json:"title"`                           // 标题
	Link        string    `xorm:"varchar(64)" json:"link"`                            // 点击跳转的地址
	Intro       string    `xorm:"varchar(32)" json:"intro"`                           // 简介
	OrderNumber int       `xorm:"default(1000)" json:"order_number"`                  // 排序
	Tag         string    `xorm:"varchar(16)" json:"tag"`                             // 标记，用于区分使用的地方
	Hidden      bool      `json:"hidden"`                                             // 隐藏
	Created     time.Time `xorm:"created" json:"created"`
	Updated     time.Time `xorm:"updated" json:"updated"`
}

// TableName 定义数据库表名
func (b *Banner) TableName() string { return "banners" }

// Get 根据非 nil 字段获取一条记录
func (b *Banner) Get() (has bool, err error) {
	return orm.NewSession(nil).Get(b)
}

// Find 查询多条数据
func (b *Banner) Find(filter *orm.Filter) (banners []Banner, err error) {
	s := orm.NewSession(filter)

	banners = make([]Banner, 0, filter.Limit)
	if err := s.Find(&banners); err != nil {
		return nil, err
	}
	return banners, err
}

// Count 统计数量
func (b *Banner) Count(filter *orm.Filter) (n int64, err error) {
	s := orm.NewSession(filter)
	return s.Count(new(Banner))
}

// Insert 插入一条记录
func (b *Banner) Insert() (ok bool, err error) {
	b.UUID = uuid.NewV1().String()
	n, err := orm.NewSession(nil).InsertOne(b)
	return n != 0, err
}

// Update 更新记录
func (b *Banner) Update(cols, omitCols []string) (n int64, err error) {
	n, err = orm.NewSession(&orm.Filter{
		Cols:          cols,
		OmitCols:      omitCols,
		Query:         "uuid=?",
		QueryArgs:     []interface{}{b.UUID},
		UpdateAllCols: cols == nil || len(cols) == 0,
	}).Update(b)

	return n, err
}

// Del 根据uuid删除一条记录
func (b *Banner) Del() (ok bool, err error) {
	n, err := orm.NewSession(&orm.Filter{Query: "uuid=?", QueryArgs: []interface{}{b.UUID}}).Delete(new(Banner))
	return n != 0, err
}

package model

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/followgo/myadmin/module/orm"
)

// MarketSegment 细分市场
type MarketSegment struct {
	UUID        string    `xorm:"varchar(36) pk 'uuid'" json:"uuid"`
	ParentUUID  string    `xorm:"varchar(36) 'parent_uuid'" json:"parent_uuid"`       // 父类
	Name        string    `xorm:"varchar(32)" json:"name"`                            // 分类名称
	Description string    `xorm:"varchar(64)" json:"description"`                     // SEO相关，文章简介
	Keywords    string    `xorm:"varchar(64)" json:"keywords"`                        // SEO相关
	ImageUUID   string    `xorm:"varchar(36) notnull 'image_uuid'" json:"image_uuid"` // 图片uuid
	OrderNumber int       `xorm:"default(1000)" json:"order_number"`                  // 排序
	Hidden      bool      `json:"hidden"`                                             // 隐藏
	Created     time.Time `xorm:"created" json:"created"`
	Updated     time.Time `xorm:"updated" json:"updated"`
}

// TableName 定义数据库表名
func (m *MarketSegment) TableName() string { return "market_segments" }

// Get 根据非 nil 字段获取一条记录
func (m *MarketSegment) Get() (has bool, err error) {
	return orm.NewSession(nil).Get(m)
}

// Find 查询多条数据
func (m *MarketSegment) Find(filter *orm.Filter) (segments []MarketSegment, err error) {
	s := orm.NewSession(filter)

	segments = make([]MarketSegment, 0, filter.Limit)
	if err := s.Find(&segments); err != nil {
		return nil, err
	}
	return segments, err
}

// Count 统计数量
func (m *MarketSegment) Count(filter *orm.Filter) (n int64, err error) {
	s := orm.NewSession(filter)
	return s.Count(new(MarketSegment))
}

// Insert 插入一条记录
func (m *MarketSegment) Insert() (ok bool, err error) {
	m.UUID = uuid.NewV1().String()
	n, err := orm.NewSession(nil).InsertOne(m)
	return n != 0, err
}

// Update 更新记录
func (m *MarketSegment) Update(cols, omitCols []string) (n int64, err error) {
	n, err = orm.NewSession(&orm.Filter{
		Cols:          cols,
		OmitCols:      omitCols,
		Query:         "uuid=?",
		QueryArgs:     []interface{}{m.UUID},
		UpdateAllCols: cols == nil || len(cols) == 0,
	}).Update(m)

	return n, err
}

// Del 根据uuid删除一条记录
func (m *MarketSegment) Del() (ok bool, err error) {
	n, err := orm.NewSession(&orm.Filter{Query: "uuid=?", QueryArgs: []interface{}{m.UUID}}).Delete(new(MarketSegment))
	return n != 0, err
}

package model

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/followgo/myadmin/module/orm"
)

// Partner 合作伙伴
type Partner struct {
	UUID        string    `xorm:"varchar(36) pk 'uuid'" json:"uuid"`
	Name        string    `xorm:"varchar(36)" json:"name"`
	Link        string    `xorm:"varchar(64)" json:"link"`
	ImageUUID   string    `xorm:"varchar(36) notnull 'image_uuid'" json:"image_uuid"` // 图片uuid
	OrderNumber int       `xorm:"default(1000)" json:"order_number"`                  // 排序
	Hidden      bool      `json:"hidden"`                                             // 隐藏
	Created     time.Time `xorm:"created" json:"created"`
	Updated     time.Time `xorm:"updated" json:"updated"`
}

// TableName 定义数据库表名
func (p *Partner) TableName() string { return "partners" }

// Get 根据非 nil 字段获取一条记录
func (p *Partner) Get() (has bool, err error) {
	return orm.NewSession(nil).Get(p)
}

// Find 查询多条数据
func (p *Partner) Find(filter *orm.Filter) (partners []Partner, err error) {
	s := orm.NewSession(filter)

	partners = make([]Partner, 0, filter.Limit)
	if err := s.Find(&partners); err != nil {
		return nil, err
	}
	return partners, err
}

// Count 统计数量
func (p *Partner) Count(filter *orm.Filter) (n int64, err error) {
	s := orm.NewSession(filter)
	return s.Count(new(Partner))
}

// Insert 插入一条记录
func (p *Partner) Insert() (ok bool, err error) {
	p.UUID = uuid.NewV1().String()
	n, err := orm.NewSession(nil).InsertOne(p)
	return n != 0, err
}

// Update 更新记录
func (p *Partner) Update(cols, omitCols []string) (n int64, err error) {
	n, err = orm.NewSession(&orm.Filter{
		Cols:          cols,
		OmitCols:      omitCols,
		Query:         "uuid=?",
		QueryArgs:     []interface{}{p.UUID},
		UpdateAllCols: cols == nil || len(cols) == 0,
	}).Update(p)

	return n, err
}

// Del 根据uuid删除一条记录
func (p *Partner) Del() (ok bool, err error) {
	n, err := orm.NewSession(&orm.Filter{Query: "uuid=?", QueryArgs: []interface{}{p.UUID}}).Delete(new(Partner))
	return n != 0, err
}

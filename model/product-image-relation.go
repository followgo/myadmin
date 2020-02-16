package model

import (
	"time"

	"github.com/followgo/myadmin/module/orm"
)

// ProductImageRelation 产品图片
type ProductImageRelation struct {
	UUID        string    `xorm:"varchar(36) pk 'uuid'" json:"uuid"`
	ProductUUID string    `xorm:"varchar(36) notnull" json:"product_uuid"`
	ImageUUID   string    `xorm:"varchar(36) notnull 'image_uuid'" json:"image_uuid"` // 图片uuid
	OrderNumber int       `xorm:"default(1000)" json:"order_number"`                  // 排序
	Hidden      bool      `json:"hidden"`                                             // 隐藏
	Created     time.Time `xorm:"created" json:"created"`
	Updated     time.Time `xorm:"updated" json:"updated"`
}

// TableName 定义数据库表名
func (r *ProductImageRelation) TableName() string { return "product_image_relation" }

// Get 根据非 nil 字段获取一条记录
func (r *ProductImageRelation) Get() (has bool, err error) {
	return orm.NewSession(nil).Get(r)
}

// Find 查询多条数据
func (r *ProductImageRelation) Find(filter *orm.Filter) (relations []ProductImageRelation, err error) {
	s := orm.NewSession(filter)

	relations = make([]ProductImageRelation, 0, filter.Limit)
	if err := s.Find(&relations); err != nil {
		return nil, err
	}
	return relations, err
}

// Count 统计数量
func (r *ProductImageRelation) Count(filter *orm.Filter) (n int64, err error) {
	return orm.NewSession(filter).Count(new(ProductImageRelation))
}

// Insert 插入一条记录
func (r *ProductImageRelation) Insert() (ok bool, err error) {
	n, err := orm.NewSession(nil).InsertOne(r)
	return n != 0, err
}

// Update 更新记录
func (r *ProductImageRelation) Update(cols, omitCols []string) (n int64, err error) {
	n, err = orm.NewSession(&orm.Filter{
		Cols:          cols,
		OmitCols:      omitCols,
		Query:         "uuid=?",
		QueryArgs:     []interface{}{r.UUID},
		UpdateAllCols: cols == nil || len(cols) == 0,
	}).Update(r)

	return n, err
}

// Del 根据uuid删除一条记录
func (r *ProductImageRelation) Del() (ok bool, err error) {
	n, err := orm.NewSession(&orm.Filter{Query: "uuid=?", QueryArgs: []interface{}{r.UUID}}).Delete(new(ProductImageRelation))
	return n != 0, err
}

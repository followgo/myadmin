package model

import (
	"github.com/followgo/myadmin/module/orm"
)

// ProductCategRelation 产品和类别的关系
type ProductCategRelation struct {
	UUID        string `xorm:"varchar(36) pk 'uuid'" json:"uuid"`
	CategUUID   string `xorm:"varchar(36) notnull" json:"categ_uuid"`
	ProductUUID string `xorm:"varchar(36) notnull" json:"product_uuid"`
}

// TableName 定义数据库表名
func (r *ProductCategRelation) TableName() string { return "product_categ_relation" }

// Get 根据非 nil 字段获取一条记录
func (r *ProductCategRelation) Get() (has bool, err error) {
	return orm.NewSession(nil).Get(r)
}

// Find 查询多条数据
func (r *ProductCategRelation) Find(filter *orm.Filter) (relations []ProductCategRelation, err error) {
	s := orm.NewSession(filter)

	relations = make([]ProductCategRelation, 0, filter.Limit)
	if err := s.Find(&relations); err != nil {
		return nil, err
	}
	return relations, err
}

// Count 统计数量
func (r *ProductCategRelation) Count(filter *orm.Filter) (n int64, err error) {
	return orm.NewSession(filter).Count(new(ProductCategRelation))
}

// Insert 插入一条记录
func (r *ProductCategRelation) Insert() (ok bool, err error) {
	n, err := orm.NewSession(nil).InsertOne(r)
	return n != 0, err
}

// Del 根据uuid删除一条记录
func (r *ProductCategRelation) Del() (ok bool, err error) {
	n, err := orm.NewSession(&orm.Filter{Query: "uuid=?", QueryArgs: []interface{}{r.UUID}}).Delete(new(ProductCategRelation))
	return n != 0, err
}

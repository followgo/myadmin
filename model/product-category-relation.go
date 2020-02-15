package model

import (
	"github.com/followgo/myadmin/module/orm"
)

// ProductCategoryRelation 产品和类别的关系
type ProductCategoryRelation struct {
	UUID         string `xorm:"varchar(36) pk 'uuid'" json:"uuid"`
	CategoryUUID string `xorm:"varchar(36) notnull" json:"category_uuid"`
	ProductUUID  string `xorm:"varchar(36) notnull" json:"product_uuid"`
}

// TableName 定义数据库表名
func (r *ProductCategoryRelation) TableName() string { return "product_category_relation" }

// Get 根据非 nil 字段获取一条记录
func (r *ProductCategoryRelation) Get() (has bool, err error) {
	return orm.NewSession(nil).Get(r)
}

// Find 查询多条数据
func (r *ProductCategoryRelation) Find(filter *orm.Filter) (relations []ProductCategoryRelation, err error) {
	s := orm.NewSession(filter)

	relations = make([]ProductCategoryRelation, 0, filter.Limit)
	if err := s.Find(&relations); err != nil {
		return nil, err
	}
	return relations, err
}

// Count 统计数量
func (r *ProductCategoryRelation) Count(filter *orm.Filter) (n int64, err error) {
	return orm.NewSession(filter).Count(new(ProductCategoryRelation))
}

// Insert 插入一条记录
func (r *ProductCategoryRelation) Insert() (ok bool, err error) {
	n, err := orm.NewSession(nil).InsertOne(r)
	return n != 0, err
}

// Del 根据uuid删除一条记录
func (r *ProductCategoryRelation) Del() (ok bool, err error) {
	n, err := orm.NewSession(&orm.Filter{Query: "uuid=?", QueryArgs: []interface{}{r.UUID}}).Delete(new(ProductCategoryRelation))
	return n != 0, err
}

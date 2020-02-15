package model

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/followgo/myadmin/module/orm"
)

// ProductCategory 产品类别
type ProductCategory struct {
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
func (c *ProductCategory) TableName() string { return "product_categories" }

// Get 根据非 nil 字段获取一条记录
func (c *ProductCategory) Get() (has bool, err error) {
	return orm.NewSession(nil).Get(c)
}

// Find 查询多条数据
func (c *ProductCategory) Find(filter *orm.Filter) (categories []ProductCategory, err error) {
	s := orm.NewSession(filter)

	categories = make([]ProductCategory, 0, filter.Limit)
	if err := s.Find(&categories); err != nil {
		return nil, err
	}
	return categories, err
}

// Count 统计数量
func (c *ProductCategory) Count(filter *orm.Filter) (n int64, err error) {
	s := orm.NewSession(filter)
	return s.Count(new(ProductCategory))
}

// Insert 插入一条记录
func (c *ProductCategory) Insert() (ok bool, err error) {
	c.UUID = uuid.NewV1().String()
	n, err := orm.NewSession(nil).InsertOne(c)
	return n != 0, err
}

// Update 更新记录
func (c *ProductCategory) Update(cols, omitCols []string) (n int64, err error) {
	n, err = orm.NewSession(&orm.Filter{
		Cols:          cols,
		OmitCols:      omitCols,
		Query:         "uuid=?",
		QueryArgs:     []interface{}{c.UUID},
		UpdateAllCols: cols == nil || len(cols) == 0,
	}).Update(c)

	return n, err
}

// Del 根据uuid删除一条记录
func (c *ProductCategory) Del() (ok bool, err error) {
	n, err := orm.NewSession(&orm.Filter{Query: "uuid=?", QueryArgs: []interface{}{c.UUID}}).Delete(new(ProductCategory))
	return n != 0, err
}

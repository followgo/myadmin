package model

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/followgo/myadmin/module/orm"
)

// ProductCateg 产品类别
type ProductCateg struct {
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
func (c *ProductCateg) TableName() string { return "product_categories" }

// Get 根据非 nil 字段获取一条记录
func (c *ProductCateg) Get() (has bool, err error) {
	return orm.NewSession(nil).Get(c)
}

// Find 查询多条数据
func (c *ProductCateg) Find(filter *orm.Filter) (categories []ProductCateg, err error) {
	s := orm.NewSession(filter)

	categories = make([]ProductCateg, 0, filter.Limit)
	if err := s.Find(&categories); err != nil {
		return nil, err
	}
	return categories, err
}

// Count 统计数量
func (c *ProductCateg) Count(filter *orm.Filter) (n int64, err error) {
	s := orm.NewSession(filter)
	return s.Count(new(ProductCateg))
}

// Insert 插入一条记录
func (c *ProductCateg) Insert() (ok bool, err error) {
	c.UUID = uuid.NewV1().String()
	n, err := orm.NewSession(nil).InsertOne(c)
	return n != 0, err
}

// Update 更新记录
func (c *ProductCateg) Update(cols, omitCols []string) (n int64, err error) {
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
func (c *ProductCateg) Del() (ok bool, err error) {
	n, err := orm.NewSession(&orm.Filter{Query: "uuid=?", QueryArgs: []interface{}{c.UUID}}).Delete(new(ProductCateg))
	return n != 0, err
}

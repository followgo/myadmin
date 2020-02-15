package model

import (
	"time"

	"github.com/followgo/myadmin/module/orm"
)

// ProductDetailPart 产品的详细描述
type ProductDetailPart struct {
	UUID        string    `xorm:"varchar(36) pk 'uuid'" json:"uuid"`
	ProductUUID string    `xorm:"varchar(36) notnull 'product_uuid'" json:"product_uuid"`
	Title       string    `xorm:"varchar(32)" json:"title"`
	ContentMd   string    `xorm:"text" json:"content_md"`            // 文章内容，Markdown 格式
	OrderNumber int       `xorm:"default(1000)" json:"order_number"` // 排序
	Hidden      bool      `json:"hidden"`                            // 隐藏
	Created     time.Time `xorm:"created" json:"created"`
	Updated     time.Time `xorm:"updated" json:"updated"`
}

// TableName 定义数据库表名
func (p *ProductDetailPart) TableName() string { return "product_detail_parts" }

// Get 根据非 nil 字段获取一条记录
func (p *ProductDetailPart) Get() (has bool, err error) {
	return orm.NewSession(nil).Get(p)
}

// Find 查询多条数据
func (p *ProductDetailPart) Find(filter *orm.Filter) (relations []ProductDetailPart, err error) {
	s := orm.NewSession(filter)

	relations = make([]ProductDetailPart, 0, filter.Limit)
	if err := s.Find(&relations); err != nil {
		return nil, err
	}
	return relations, err
}

// Count 统计数量
func (p *ProductDetailPart) Count(filter *orm.Filter) (n int64, err error) {
	return orm.NewSession(filter).Count(new(ProductDetailPart))
}

// Insert 插入一条记录
func (p *ProductDetailPart) Insert() (ok bool, err error) {
	n, err := orm.NewSession(nil).InsertOne(p)
	return n != 0, err
}

// Update 更新记录
func (p *ProductDetailPart) Update(cols, omitCols []string) (n int64, err error) {
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
func (p *ProductDetailPart) Del() (ok bool, err error) {
	n, err := orm.NewSession(&orm.Filter{Query: "uuid=?", QueryArgs: []interface{}{p.UUID}}).Delete(new(ProductDetailPart))
	return n != 0, err
}

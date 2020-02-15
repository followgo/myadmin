package model

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/followgo/myadmin/module/orm"
)

// TypicalCase 典型案例
type TypicalCase struct {
	UUID                       string    `xorm:"varchar(36) pk 'uuid'" json:"uuid"`
	MarketSegmentUUID          string    `xorm:"varchar(36) notnull 'market_segment_uuid'" json:"market_segment_uuid"`
	Title                      string    `json:"title"`                                                                            // 文章标题
	Description                string    `json:"description"`                                                                      // SEO相关，文章简介
	Keywords                   string    `json:"keywords"`                                                                         // SEO相关
	ImageUUID                  string    `xorm:"varchar(36) 'image_uuid'" json:"image_uuid"`                                       // 图片
	ContentMd                  string    `xorm:"text" json:"content_md"`                                                           // 文章内容，Markdown 格式
	Released                   bool      `json:"released"`                                                                         // 发布
	PostAt                     time.Time `json:"post_at"`                                                                          // 发布日期，一般按照这个进行排序
	Sticky                     bool      `json:"sticky"`                                                                           // 置顶
	RelatedProductCategoryUUID string    `xorm:"varchar(36) 'related_product_category_uuid'" json:"related_product_category_uuid"` // 关联的产品分类，用于展示的相关推荐产品
	Created                    time.Time `xorm:"created" json:"created"`
	Updated                    time.Time `xorm:"updated" json:"updated"`
}

// TableName 定义数据库表名
func (c *TypicalCase) TableName() string { return "typical_cases" }

// Get 根据非 nil 字段获取一条记录
func (c *TypicalCase) Get() (has bool, err error) {
	return orm.NewSession(nil).Get(c)
}

// Find 查询多条数据
func (c *TypicalCase) Find(filter *orm.Filter) (cases []TypicalCase, err error) {
	sess := orm.NewSession(filter)

	cases = make([]TypicalCase, 0, filter.Limit)
	if err := sess.Find(&cases); err != nil {
		return nil, err
	}
	return cases, err
}

// Count 统计数量
func (c *TypicalCase) Count(filter *orm.Filter) (n int64, err error) {
	return orm.NewSession(filter).Count(new(TypicalCase))
}

// Insert 插入一条记录
func (c *TypicalCase) Insert() (ok bool, err error) {
	c.UUID = uuid.NewV1().String()
	n, err := orm.NewSession(nil).InsertOne(c)
	return n != 0, err
}

// Update 更新记录
func (c *TypicalCase) Update(cols, omitCols []string) (n int64, err error) {
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
func (c *TypicalCase) Del() (ok bool, err error) {
	n, err := orm.NewSession(&orm.Filter{Query: "uuid=?", QueryArgs: []interface{}{c.UUID}}).Delete(new(TypicalCase))
	return n != 0, err
}

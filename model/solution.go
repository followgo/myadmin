package model

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/followgo/myadmin/module/orm"
)

// Solution 解决方案
type Solution struct {
	UUID                    string    `xorm:"varchar(36) pk 'uuid'" json:"uuid"`
	MarketSegmentUUID       string    `xorm:"varchar(36) notnull 'market_segment_uuid'" json:"market_segment_uuid"`
	Title                   string    `json:"title"`                                                                      // 文章标题
	Description             string    `json:"description"`                                                                // SEO相关，文章简介
	Keywords                string    `json:"keywords"`                                                                   // SEO相关
	ImageUUID               string    `xorm:"varchar(36) 'image_uuid'" json:"image_uuid"`                                 // 图片
	ContentMd               string    `xorm:"text" json:"content_md"`                                                     // 文章内容，Markdown 格式
	Released                bool      `json:"released"`                                                                   // 发布
	OrderNumber             int       `xorm:"default(1000)" json:"order_number"`                                          // 排序
	RelatedProductCategUUID string    `xorm:"varchar(36) 'related_product_categ_uuid'" json:"related_product_categ_uuid"` // 关联的产品分类，用于展示的相关推荐产品
	Created                 time.Time `xorm:"created" json:"created"`
	Updated                 time.Time `xorm:"updated" json:"updated"`
}

// TableName 定义数据库表名
func (s *Solution) TableName() string { return "solutions" }

// Get 根据非 nil 字段获取一条记录
func (s *Solution) Get() (has bool, err error) {
	return orm.NewSession(nil).Get(s)
}

// Find 查询多条数据
func (s *Solution) Find(filter *orm.Filter) (solutions []Solution, err error) {
	sess := orm.NewSession(filter)

	solutions = make([]Solution, 0, filter.Limit)
	if err := sess.Find(&solutions); err != nil {
		return nil, err
	}
	return solutions, err
}

// Count 统计数量
func (s *Solution) Count(filter *orm.Filter) (n int64, err error) {
	return orm.NewSession(filter).Count(new(Solution))
}

// Insert 插入一条记录
func (s *Solution) Insert() (ok bool, err error) {
	s.UUID = uuid.NewV1().String()
	n, err := orm.NewSession(nil).InsertOne(s)
	return n != 0, err
}

// Update 更新记录
func (s *Solution) Update(cols, omitCols []string) (n int64, err error) {
	n, err = orm.NewSession(&orm.Filter{
		Cols:          cols,
		OmitCols:      omitCols,
		Query:         "uuid=?",
		QueryArgs:     []interface{}{s.UUID},
		UpdateAllCols: cols == nil || len(cols) == 0,
	}).Update(s)

	return n, err
}

// Del 根据uuid删除一条记录
func (s *Solution) Del() (ok bool, err error) {
	n, err := orm.NewSession(&orm.Filter{Query: "uuid=?", QueryArgs: []interface{}{s.UUID}}).Delete(new(Solution))
	return n != 0, err
}

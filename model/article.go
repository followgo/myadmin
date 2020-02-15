package model

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/followgo/myadmin/module/orm"
)

// Article 文章
type Article struct {
	UUID                       string    `xorm:"varchar(36) pk 'uuid'" json:"uuid"`
	CategoryUUID               string    `xorm:"varchar(36) notnull 'category_uuid'" json:"category_uuid"`
	Title                      string    `json:"title"`                                                                            // 文章标题
	Description                string    `json:"description"`                                                                      // SEO相关，文章简介
	Keywords                   string    `json:"keywords"`                                                                         // SEO相关
	Source                     string    `xorm:"varchar(32)" json:"source"`                                                        // 转载来源
	SourceUrl                  string    `xorm:"varchar(64)" json:"source_url"`                                                    // 转载来源的链接
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
func (a *Article) TableName() string { return "articles" }

// Get 根据非 nil 字段获取一条记录
func (a *Article) Get() (has bool, err error) {
	return orm.NewSession(nil).Get(a)
}

// Find 查询多条数据
func (a *Article) Find(filter *orm.Filter) (articles []Article, err error) {
	s := orm.NewSession(filter)

	articles = make([]Article, 0, filter.Limit)
	if err := s.Find(&articles); err != nil {
		return nil, err
	}
	return articles, err
}

// Count 统计数量
func (a *Article) Count(filter *orm.Filter) (n int64, err error) {
	s := orm.NewSession(filter)
	return s.Count(new(Article))
}

// Insert 插入一条记录
func (a *Article) Insert() (ok bool, err error) {
	a.UUID = uuid.NewV1().String()
	n, err := orm.NewSession(nil).InsertOne(a)
	return n != 0, err
}

// Update 更新记录
func (a *Article) Update(cols, omitCols []string) (n int64, err error) {
	n, err = orm.NewSession(&orm.Filter{
		Cols:          cols,
		OmitCols:      omitCols,
		Query:         "uuid=?",
		QueryArgs:     []interface{}{a.UUID},
		UpdateAllCols: cols == nil || len(cols) == 0,
	}).Update(a)

	return n, err
}

// Del 根据uuid删除一条记录
func (a *Article) Del() (ok bool, err error) {
	n, err := orm.NewSession(&orm.Filter{Query: "uuid=?", QueryArgs: []interface{}{a.UUID}}).Delete(new(Article))
	return n != 0, err
}

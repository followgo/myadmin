package model

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/followgo/myadmin/module/orm"
)

// CompanyNews 公司资讯
type CompanyNews struct {
	UUID        string    `xorm:"varchar(36) pk 'uuid'" json:"uuid"`
	Category    string    `xorm:"varchar(16)" json:"category"`                // 类型：公司，产品，展会
	ImageUUID   string    `xorm:"varchar(36) 'image_uuid'" json:"image_uuid"` // 图片
	Description string    `xorm:"varchar(64)" json:"description"`             // SEO相关，文章简介
	Keywords    string    `xorm:"varchar(64)" json:"keywords"`                // SEO相关
	ContentMd   string    `json:"content_md"`                                 // 文章内容，Markdown 格式
	Released    bool      `json:"released"`                                   // 发布
	PostAt      time.Time `json:"post_at"`                                    // 发布时间
	Sticky      bool      `json:"sticky"`                                     // 置顶
	Created     time.Time `xorm:"created" json:"created"`
	Updated     time.Time `xorm:"updated" json:"updated"`
}

// TableName 定义数据库表名
func (c *CompanyNews) TableName() string { return "company_news" }

// Get 根据非 nil 字段获取一条记录
func (c *CompanyNews) Get() (has bool, err error) {
	return orm.NewSession(nil).Get(c)
}

// Find 查询多条数据
func (c *CompanyNews) Find(filter *orm.Filter) (news []CompanyNews, err error) {
	s := orm.NewSession(filter)

	news = make([]CompanyNews, 0, filter.Limit)
	if err := s.Find(&news); err != nil {
		return nil, err
	}
	return news, err
}

// Count 统计数量
func (c *CompanyNews) Count(filter *orm.Filter) (n int64, err error) {
	s := orm.NewSession(filter)
	return s.Count(new(CompanyNews))
}

// Insert 插入一条记录
func (c *CompanyNews) Insert() (ok bool, err error) {
	c.UUID = uuid.NewV1().String()
	n, err := orm.NewSession(nil).InsertOne(c)
	return n != 0, err
}

// Update 更新记录
func (c *CompanyNews) Update(cols, omitCols []string) (n int64, err error) {
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
func (c *CompanyNews) Del() (ok bool, err error) {
	n, err := orm.NewSession(&orm.Filter{Query: "uuid=?", QueryArgs: []interface{}{c.UUID}}).Delete(new(CompanyNews))
	return n != 0, err
}

package model

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/followgo/myadmin/module/orm"
)

// Product 产品
type Product struct {
	UUID           string    `xorm:"varchar(36) pk 'uuid'" json:"uuid"`
	CategUUID      string    `xorm:"varchar(36) 'categ_uuid'" json:"categ_uuid"`
	Title          string    `xorm:"varchar(32)" json:"title"`                                       // 主标题
	SubTitle       string    `xorm:"varchar(32)" json:"sub_title"`                                   // 副标题
	ModelCode      string    `xorm:"varchar(32)" json:"model_code"`                                  // 型号代码
	Description    string    `xorm:"varchar(64)" json:"description"`                                 // SEO相关，文章简介
	Keywords       string    `xorm:"varchar(64)" json:"keywords"`                                    // SEO相关
	CoverImageUUID string    `xorm:"varchar(36) notnull 'cover_image_uuid'" json:"cover_image_uuid"` // 图片uuid
	OrderNumber    int       `xorm:"default(1000)" json:"order_number"`                              // 排序
	Hidden         bool      `json:"hidden"`                                                         // 隐藏
	Created        time.Time `xorm:"created" json:"created"`
	Updated        time.Time `xorm:"updated" json:"updated"`
}

// TableName 定义数据库表名
func (p *Product) TableName() string { return "products" }

// Get 根据非 nil 字段获取一条记录
func (p *Product) Get() (has bool, err error) {
	return orm.NewSession(nil).Get(p)
}

// Find 查询多条数据
func (p *Product) Find(filter *orm.Filter) (products []Product, err error) {
	s := orm.NewSession(filter)

	products = make([]Product, 0, filter.Limit)
	if err := s.Find(&products); err != nil {
		return nil, err
	}
	return products, err
}

// Count 统计数量
func (p *Product) Count(filter *orm.Filter) (n int64, err error) {
	s := orm.NewSession(filter)
	return s.Count(new(Product))
}

// Insert 插入一条记录
func (p *Product) Insert() (ok bool, err error) {
	p.UUID = uuid.NewV1().String()
	n, err := orm.NewSession(nil).InsertOne(p)
	return n != 0, err
}

// Update 更新记录
func (p *Product) Update(cols, omitCols []string) (n int64, err error) {
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
func (p *Product) Del() (ok bool, err error) {
	n, err := orm.NewSession(&orm.Filter{Query: "uuid=?", QueryArgs: []interface{}{p.UUID}}).Delete(new(Product))
	return n != 0, err
}

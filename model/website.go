package model

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/followgo/myadmin/module/orm"
)

// Website 网站设置
type Website struct {
	UUID                   string    `xorm:"varchar(36) pk 'uuid'" json:"uuid"`
	OrgShortName           string    `xorm:"varchar(32)" json:"org_short_name"`     // 组织简称
	MainTitle              string    `xorm:"varchar(64)" json:"main_title"`         // 首页标题
	Description            string    `xorm:"varchar(256)" json:"description"`       // SEO描述
	Keywords               string    `xorm:"varchar(256)" json:"keywords"`          // SEO关键字
	Copyright              string    `xorm:"varchar(64)" json:"copyright"`          // 版权信息
	MIITLicence            string    `xorm:"varchar(32)" json:"miit_licence"`       // 工信部备案号
	MIITLicenceLink        string    `xorm:"varchar(128)" json:"miit_licence_link"` // 工信部备案号的链接指向
	CustomHeaderCode       string    `xorm:"text" json:"custom_header_code"`        // 在 </head> 前添加的代码
	CustomFooterCode       string    `xorm:"text" json:"custom_footer_code"`        // 在 </body> 前添加的代码
	CustomFooterContentRaw string    `xorm:"text" json:"custom_footer_content_raw"` // 网页底部的内容
	Created                time.Time `xorm:"created" json:"created"`
	Updated                time.Time `xorm:"updated" json:"updated"`
}

// TableName 定义数据库表名
func (w *Website) TableName() string { return "Website" }

// Insert 插入一条记录
func (w *Website) Insert() (ok bool, err error) {
	w.UUID = uuid.NewV1().String()
	n, err := orm.NewSession(nil).InsertOne(w)
	return n != 0, err
}

// Get 根据非 nil 字段获取一条记录
func (w *Website) Get() (has bool, err error) {
	return orm.NewSession(nil).Get(w)
}

// Update 更新记录
func (w *Website) Update(cols []string) (n int64, err error) {
	n, err = orm.NewSession(&orm.Filter{
		Cols:          cols,
		Query:         "uuid=?",
		QueryArgs:     []interface{}{w.UUID},
		UpdateAllCols: cols == nil || len(cols) == 0,
	}).Update(w)

	return n, err
}

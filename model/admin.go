package model

import (
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"

	. "github.com/followgo/myadmin/config"
	"github.com/followgo/myadmin/module/orm"
	"github.com/followgo/myadmin/util"
)

// Admin 管理员用户信息
type Admin struct {
	UUID          string   `xorm:"varchar(36) pk 'uuid'"`
	Email         string   `xorm:"varchar(64) notnull unique"`
	Username      string   `xorm:"varchar(32)"`
	Password      string   `xorm:"varchar(255) notnull"`
	Roles         []string `xorm:"varchar(128)"`
	Enabled       bool
	LastLoginFrom string `xorm:"varchar(64)"`
	LastLoginAt   time.Time
	LoginCount    uint
	Created       time.Time `xorm:"created"`
	Updated       time.Time `xorm:"updated"`
}

// TableName 定义数据库表名
func (u *Admin) TableName() string { return "users" }

// Get 根据非 nil 字段获取一条记录
func (u *Admin) Get() (has bool, err error) {
	has, err = orm.NewSession(nil).Get(u)
	u.coverPwd()
	return
}

// Find 查询多条数据
func (u *Admin) Find(filter *orm.Filter) (users []Admin, err error) {
	s := orm.NewSession(filter)

	users = make([]Admin, 0, filter.Limit)
	if err := s.Find(&users); err != nil {
		return nil, err
	}

	for i := range users {
		users[i].coverPwd()
	}
	return users, err
}

// Count 统计数量
func (u *Admin) Count(filter *orm.Filter) (n int64, err error) {
	s := orm.NewSession(filter)
	return s.Count(new(Admin))
}

// Insert 插入一条记录
func (u *Admin) Insert() (ok bool, err error) {
	u.hashPwd()
	u.UUID = uuid.NewV1().String()
	n, err := orm.NewSession(nil).InsertOne(u)
	u.coverPwd()
	return n != 0, err
}

// Update 更新记录
func (u *Admin) Update(cols, omitCols []string) (n int64, err error) {
	if u.Password != "" {
		u.hashPwd()
	}

	n, err = orm.NewSession(&orm.Filter{
		Cols:          cols,
		OmitCols:      omitCols,
		Query:         "uuid=?",
		QueryArgs:     []interface{}{u.UUID},
		UpdateAllCols: cols == nil || len(cols) == 0,
	}).Update(u)

	u.coverPwd()
	return n, err
}

// Del 根据uuid删除一条记录
func (u *Admin) Del() (ok bool, err error) {
	n, err := orm.NewSession(&orm.Filter{Query: "uuid=?", QueryArgs: []interface{}{u.UUID}}).Delete(new(Admin))
	return n != 0, err
}

// Validate 验证用户
func (u *Admin) Validate() (ok bool, err error) {
	u.hashPwd()
	_user := new(Admin)
	ok, err = orm.NewSession(&orm.Filter{
		Query:     "(username=? OR email=?) AND password=?",
		QueryArgs: []interface{}{u.Username, u.Email, u.Password},
	}).Get(_user)
	u.coverPwd()

	if err != nil {
		return false, err
	}
	if !ok {
		return false, nil
	}

	u.LastLoginAt = time.Now()
	u.LoginCount = _user.LoginCount + 1
	_, _ = orm.NewSession(&orm.Filter{Cols: []string{"last_login_from", "last_login_at", "login_count"}}).NoAutoTime().Update(u)
	return true, nil
}

// coverPwd 掩盖密码
func (u *Admin) coverPwd() { u.Password = "#########" }

// hashPwd 哈希密码
func (u *Admin) hashPwd() {
	u.Password = util.Hash(strings.NewReader(u.Password), []byte(Cfg.SecuritySalt))
}

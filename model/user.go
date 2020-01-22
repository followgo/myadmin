package model

import (
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"

	. "github.com/followgo/myadmin/config"
	"github.com/followgo/myadmin/module/orm"
	"github.com/followgo/myadmin/util"
)

// User 管理员用户信息
type User struct {
	UUID          string `xorm:"pk 'uuid'"`
	Email         string `xorm:"notnull unique"`
	Username      string
	Password      string `xorm:"notnull"`
	Roles         []string
	Enabled       bool
	LastLoginFrom string
	LastLoginAt   time.Time
	LoginCount    uint
	Created       time.Time `xorm:"created"`
	Updated       time.Time `xorm:"updated"`
}

// TableName 定义数据库表名
func (u *User) TableName() string { return "users" }

// Get 根据非 nil 字段获取一条记录
func (u *User) Get() (has bool, err error) {
	has, err = orm.NewSession(nil).Get(u)
	u.coverPwd()
	return
}

// Find 查询多条数据
func (u *User) Find(filter *orm.Filter) (users []User, err error) {
	s := orm.NewSession(filter)

	users = make([]User, 0, 100)
	if err := s.Find(&users); err != nil {
		return nil, err
	}

	for i := range users {
		users[i].coverPwd()
	}
	return users, err
}

// Count 统计数量
func (u *User) Count(filter *orm.Filter) (n int64, err error) {
	s := orm.NewSession(filter)
	return s.Count(new(User))
}

// Insert 插入一条记录
func (u *User) Insert() (ok bool, err error) {
	u.hashPwd()
	u.UUID = uuid.NewV1().String()
	n, err := orm.NewSession(nil).InsertOne(u)
	u.coverPwd()
	return n != 0, err
}

// Update 更新记录
func (u *User) Update(cols, omitCols []string) (n int64, err error) {
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
func (u *User) Del() (ok bool, err error) {
	n, err := orm.NewSession(&orm.Filter{Query: "uuid=?", QueryArgs: []interface{}{u.UUID}}).Delete(new(User))
	return n != 0, err
}

// Validate 验证用户
func (u *User) Validate() (ok bool, err error) {
	u.hashPwd()
	_user := new(User)
	ok, err = orm.NewSession(&orm.Filter{Query: "(username=? OR email=?) AND password=?", QueryArgs: []interface{}{u.Username, u.Email, u.Password}}).Get(_user)
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
func (u *User) coverPwd() { u.Password = "#########" }

// hashPwd 哈希密码
func (u *User) hashPwd() {
	u.Password = util.Hash(strings.NewReader(u.Password), []byte(Cfg.SecuritySalt))
}

package model

import (
	"encoding/hex"
	"time"

	"golang.org/x/crypto/blake2b"

	. "github.com/followgo/myadmin/config"
	. "github.com/followgo/myadmin/module/orm"

	uuid "github.com/satori/go.uuid"
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
	Created       time.Time `xorm:"created"`
	Updated       time.Time `xorm:"updated"`
	Version       int       `xorm:"version"`
}

// TableName 定义数据库表名
func (u *User) TableName() string { return "users" }

// GetByUUID 根据UUID获取一条记录
func (u *User) GetByUUID() (has bool, err error) {
	has, err = Orm.Where("uuid=?", u.UUID).Get(u)
	u.coverPwd()
	return
}

// GetByUUID 根据 用户名或邮箱 和 密码 获取一条记录
func (u *User) GetByUsernamePassword() (has bool, err error) {
	u.hashPwd()
	has, err = Orm.Where("(username=? OR email) AND password", u.UUID, u.Email, u.Password).Get(u)
	u.coverPwd()
	return
}

// Find 查询多条数据
func (u *User) Find(filter *FindFilter) (users []User, err error) {
	s := AttachFindFilter(Orm, filter)

	users = make([]User, 100)
	if err := s.Find(&users); err != nil {
		return nil, err
	}

	for i := range users {
		users[i].coverPwd()
	}
	return users, err
}

// Count 统计数量
func (u *User) Count(filter *FindFilter) (n int64, err error) {
	s := AttachFindFilter(Orm, filter)
	return s.Count(new(User))
}

// Insert 插入一条记录
func (u *User) Insert() (ok bool, err error) {
	u.hashPwd()
	u.UUID = uuid.NewV1().String()
	n, err := Orm.InsertOne(u)
	u.coverPwd()
	return  n != 0, err
}

// Update 更新记录
func (u *User) Update(filter *UpdateFilter) (ok bool, err error) {
	s := AttachUpdateFilter(Orm, filter)
	n, err := s.Update(u)
	return n != 0, err
}

// coverPwd 掩盖密码
func (u *User) coverPwd() {
	u.Password = "#########"
}

// hashPwd 哈希密码
func (u *User) hashPwd() {
	h, _ := blake2b.New384([]byte(C.SecuritySalt))
	h.Write([]byte("fO1HX6qlkNA7bXk3DM1SDp4L"))
	h.Write([]byte(u.Password))
	u.Password = hex.EncodeToString(h.Sum(nil))
}

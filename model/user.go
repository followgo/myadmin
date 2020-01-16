package model

import (
	"time"

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

// Insert 插入一条记录
func (u *User) InsertOne() error {
	u.UUID = uuid.NewV1().String()
	_, err := Orm.InsertOne(u)
	return err
}

// Insert 插入一条记录
func InsertUsers(users []User) (int64, error) {
	for i := range users {
		users[i].UUID = uuid.NewV1().String()
	}
	return Orm.Insert(users)
}

// Verify 校验用户名和密码
func (u *User) Verify() bool {

}

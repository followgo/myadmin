// onlineuser 维护在线用户列表
package onlineuser

import (
	"sync"
	"time"
)

var (
	// onlineLifetime 在线用户在 n 分钟没有刷新，则会被清理
	onlineLifetime = 30 * time.Minute

	// onlineUsers 在线用户列表
	onlineUsers = new(sync.Map)
)

// init 清理过去的用户
func init() {
	go func() {
		for range time.Tick(time.Minute) {
			now := time.Now()
			expKeys := make([]string, 0, 5)

			// 查找过期用户
			onlineUsers.Range(func(key, value interface{}) bool {
				if now.Sub(value.(User).UpdateAt) > onlineLifetime {
					expKeys = append(expKeys, key.(string))
				}
				return true
			})

			// 清理过期用户
			for _, k := range expKeys {
				onlineUsers.Delete(k)
			}
		}
	}()
}

type User struct {
	Username  string
	Roles     []string
	RemoteIP  string
	UserAgent string
	LoginAt   time.Time
	UpdateAt  time.Time
}

// GetUser 获取一个在线用户
func GetUser(uuid string) User {
	var u = User{}
	onlineUsers.Range(func(key, value interface{}) bool {
		if uuid == key.(string) {
			u = value.(User)
			return false // returns false, range stops the iteration.
		}
		return true
	})

	return u
}

// GetUsers 获取所有在线用户
func GetUsers() map[string]User {
	users := make(map[string]User)
	onlineUsers.Range(func(key, value interface{}) bool {
		users[key.(string)] = value.(User)
		return true
	})
	return users
}

// AddUser 添加用户
func AddUser(uuid string, u User) {
	u.LoginAt = time.Now()
	u.UpdateAt = u.LoginAt
	onlineUsers.Store(uuid, u)
}

// RefreshUser 更新刷新时间
func RefreshUser(uuid string) {
	v, ok := onlineUsers.Load(uuid)
	if ok {
		user := v.(User)
		user.UpdateAt = time.Now()
		onlineUsers.Store(uuid, user)
	}
}

// RemoteUser 移除用户
func RemoteUser(uuid string) {
	onlineUsers.Delete(uuid)
}

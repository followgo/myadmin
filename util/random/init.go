package random

import (
	"math/rand"
	"time"
)

// init 初始化rand，使用时间戳作为随机种子
func init() {
	rand.Seed(time.Now().UnixNano())
}

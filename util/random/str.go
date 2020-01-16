package random

import (
	"math/rand"
	"strings"
	"time"
)

// Charsets
const (
	Uppercase    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Lowercase    = "abcdefghijklmnopqrstuvwxyz"
	Alphabetic   = Uppercase + Lowercase
	Numeric      = "0123456789"
	Alphanumeric = Alphabetic + Numeric
	Symbols      = "`" + `~!@#$%^&*()-_+={}[]|\;:"<>,./?`
	Hex          = Numeric + "abcdef"
)

// String 返回随机字符串
func String(length uint8, charsets ...string) string {
	rand.Seed(time.Now().UnixNano())

	charset := strings.Join(charsets, "")
	if charset == "" {
		charset = Alphanumeric
	}

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Int63()%int64(len(charset))]
	}

	return string(b)
}

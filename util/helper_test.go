package util

import (
	"testing"
)

func TestStructToStruct(t *testing.T) {
	type (
		s1 struct {
			Name  string
			Value int
		}
		s2 struct {
			Name  string
			Value uint16
		}
	)
	v1 := s1{"abc", 123}
	v2 := s2{"ccc", 222}

	err := StructToStruct(&v1, v2)
	if err != nil {
		t.Log(err)
	}
	t.Log(v1)
	t.Log(v2)
}

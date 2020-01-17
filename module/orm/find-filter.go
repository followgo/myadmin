package orm

import (
	"github.com/xormplus/xorm"
)

// FindFilter 查询条件
type FindFilter struct {
	// AscCols 指定字段名正序排序
	AscCols []string

	// DescCols 指定字段名逆序排序
	DescCols []string

	// Cols 只查询某些指定的字段
	Cols []string

	// OmitCols 排除某些指定的字段
	OmitCols []string

	// Query 查询条件
	Query string

	// QueryArgs 查询条件的参数
	QueryArgs []interface{}

	// GroupByKeys 分组
	GroupByKeys string

	// Limit 限制获取的数目，第一个参数为条数，第二个参数表示开始位置
	Limit [2]int
}

// AttachFindFilter 附加查询条件
func AttachFindFilter(o xorm.EngineInterface, filter *FindFilter) *xorm.Session {
	s := o.NewSession()
	if filter == nil {
		return s
	}

	if filter.Cols != nil {
		s = s.Cols(filter.Cols...)
	}

	if filter.OmitCols != nil {
		s = s.Omit(filter.OmitCols...)
	}

	if filter.AscCols != nil {
		s = s.Asc(filter.AscCols...)
	}

	if filter.DescCols != nil {
		s = s.Desc(filter.DescCols...)
	}

	if filter.Query != "" {
		if filter.QueryArgs != nil {
			s = s.Where(filter.Query, filter.QueryArgs...)
		} else {
			s = s.Where(filter.Query)
		}
	}

	if filter.GroupByKeys != "" {
		s = s.GroupBy(filter.GroupByKeys)
	}

	if filter.Limit[0] != 0 {
		s = s.Limit(filter.Limit[0], filter.Limit[1])
	}

	return s
}

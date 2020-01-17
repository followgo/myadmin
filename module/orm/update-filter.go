package orm

import (
	"github.com/xormplus/xorm"
)

// UpdateFilter 更新条件
type UpdateFilter struct {
	// MustCols 某些字段必须更新
	MustCols []string

	// OmitCols 排除某些指定的字段
	OmitCols []string

	// Query 查询条件
	Query interface{}

	// QueryArgs 查询条件的参数
	QueryArgs []interface{}
}

// AttachUpdateFilter 附加更新条件
func AttachUpdateFilter(o xorm.EngineInterface, filter *UpdateFilter) *xorm.Session {
	s := o.NewSession()
	if filter == nil {
		return s
	}

	if filter.MustCols != nil {
		s = s.Cols(filter.MustCols...)
	}

	if filter.OmitCols != nil {
		s = s.Omit(filter.OmitCols...)
	}

	if filter.Query != nil {
		if filter.QueryArgs != nil {
			s = s.Where(filter.Query, filter.QueryArgs...)
		} else {
			s = s.Where(filter.Query)
		}
	}

	return s
}

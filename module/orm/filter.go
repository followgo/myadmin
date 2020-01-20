package orm

// Filter 查询条件
type Filter struct {
	// AscCols 指定字段名正序排序
	AscCols []string `json:"asc_cols" form:"asc_cols" query:"asc_cols"`

	// DescCols 指定字段名逆序排序
	DescCols []string `json:"desc_cols" form:"desc_cols" query:"desc_cols"`

	// Cols 只查询某些指定的字段
	Cols []string `json:"cols" form:"cols" query:"cols"`

	// OmitCols 排除某些指定的字段
	OmitCols []string `json:"omit_cols" form:"omit_cols" query:"omit_cols"`

	// Query 查询条件
	Query string `json:"query" form:"query" query:"query"`

	// QueryArgs 查询条件的参数
	QueryArgs []interface{} `json:"query_args" form:"query_args" query:"query_args"`

	// GroupByKeys 分组
	GroupByKeys string `json:"group_by_keys" form:"group_by_keys" query:"group_by_keys"`

	// Limit 限制获取的数目，第一个参数为条数，第二个参数表示开始位置
	Limit [2]int `json:"limit" form:"limit" query:"limit"`
}
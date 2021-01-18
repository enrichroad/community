package database

import "database/sql"

type ParamPair struct {
	Query string        // 查询
	Args  []interface{} // 参数
}

// 排序信息
type OrderByCol struct {
	Column string // 排序字段
	Asc    bool   // 是否正序
}

func SqlNullString(value string) sql.NullString {
	return sql.NullString{
		String: value,
		Valid:  len(value) > 0,
	}
}

package serializing

import (
	"github.com/enrichroad/community/pagination"
)

// 分页返回数据
type PageResult struct {
	Page    *pagination.Paging `json:"page"`    // 分页信息
	Results interface{}        `json:"results"` // 数据
}

// Cursor分页返回数据
type CursorResult struct {
	Results interface{} `json:"results"` // 数据
	Cursor  string      `json:"cursor"`  // 下一页
}

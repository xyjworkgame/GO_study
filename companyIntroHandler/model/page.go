package model

const (
	DefaultPageSize = 20
)

type Pagination struct {
	PageNum  int32 `json:"pageNum"` //当前页
	PageSize int32 `json:"pageSize"` // 页面数据数量
	Total    int32 `json:"total"` // 数据量
	PageTotal int32 `json:"pageTotal"` // 页面总数
}

func (p *Pagination) Check() {
	if p.PageNum <= 0 {
		p.PageNum = 1
	}
	if p.PageSize <= 0 || p.PageSize > 100 {
		p.PageSize = DefaultPageSize
	}
}


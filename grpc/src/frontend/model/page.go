package model

const (
	DefaultPageSize = 20
)

type Pagination struct {
	PageNum          int   `json:"page_num"`  //当前页
	PageSize         int   `json:"page_size"` // 页面数据数量
	Size             int   `json:"size"`      // 当前页面数据量
	Pages            int   `json:"pages"`     // 页面总数
	Total            int   `json:"total"`
	LastPage         int   `json:"last_page"`
	IsFirstPage      bool  `json:"is_first_page"`
	IsLastPage       bool  `json:"is_last_page"`
}

// 补充其余数据，例如HasNext 参数
func (p *Pagination) Check() {
	//	1. 首先页面数据大小为 20 ，防止请求过多
	if p.PageNum <= 0 {
		p.PageNum = 1
	}
	if p.PageSize <= 0 || p.PageSize > 100 {
		p.PageSize = DefaultPageSize
	}
}

// 数据库查找完毕后，执行方法
func (p *Pagination) CheckOut(mount int){
	p.Total  = mount

	//	2. 如果Total 数据量小于 pageSIze，那么
	if p.Total < p.PageSize {
		p.Size  = p.Total
	}else {
		p.Size = p.PageSize
	}

	p.Pages = int(p.Total/p.PageSize) + 1
	p.LastPage = p.Pages

	if p.PageNum ==1 {
		p.IsFirstPage  = true
	}else {
		p.IsFirstPage = false
	}

	if p.Pages == p.PageNum{
		p.IsLastPage = true
	}else{
		p.IsLastPage = false
	}
}

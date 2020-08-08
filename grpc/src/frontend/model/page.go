package model

const (
	DefaultPageSize = 20
)

type Pagination struct {
	PageNum          int   `json:"page_num"`  //当前页
	PageSize         int   `json:"page_size"` // 页面数据数量
	Size             int   `json:"size"`      // 数据量
	Pages            int   `json:"pages"`     // 页面总数
	Total            int   `json:"total"`
	LastPage         int   `json:"last_page"`
	IsFirstPage      bool  `json:"is_first_page"`
	IsLastPage       bool  `json:"is_last_page"`
	HasPreviousPage  bool  `json:"has_previous_page"`
	HasNextPage      bool  `json:"has_next_page"`
	NavigatePages    int   `json:"navigate_pages"`     //导航页
	NavigatePageNums [10]int `json:"navigate_page_nums"` // 导航页列表,最多传10页
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

	//	2.

}

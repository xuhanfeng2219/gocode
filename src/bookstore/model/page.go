package model

type Page struct {
	Books       []*Book
	PageNo      int64
	PageSize    int64
	TotalPageNo int64
	TotalRecord int64
	MinPrice    string
	MaxPrice    string
	IsLogin     bool
	Username    string
}

func (p *Page) IsHasPrev() bool {
	return p.PageNo > 1
}

func (p *Page) IsHasNext() bool {
	return p.PageNo < p.TotalPageNo
}

func (p *Page) GetPrevPage() int64 {
	if p.IsHasPrev() {
		return p.PageNo - 1
	}
	return 1
}

func (p *Page) GetNextPageNo() int64 {
	if p.IsHasNext() {
		return p.PageNo + 1
	}
	return p.TotalPageNo
}

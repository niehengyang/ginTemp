package response

import (
	"gorm.io/gorm"
)

// PageOption 分页参数
type OpenPageOption struct {
	DB      *gorm.DB
	Start   int
	Size    int
	Odkey   string
	Reverse int
	ShowSQL bool
}

type openPageMeta struct {
	Total int64 `json:"total"`
	Size  int   `json:"size"`
	Start int   `json:"start"`
}

// Paginator 分页返回
type OpenPaginator struct {
	PageInfo openPageMeta `json:"pageInfo"`
	Items    interface{}  `json:"items"`
}

// Paging 分页
func openPaging(p *OpenPageOption, result interface{}) *OpenPaginator {

	db := p.DB

	if p.ShowSQL {
		db = db.Debug()
	}
	if p.Start < 1 {
		p.Start = 1
	}
	if p.Size == 0 {
		p.Size = 10
	}
	if p.Odkey != "" {
		orderStr := p.Odkey
		if p.Reverse != 0 {
			orderStr += " ASC"
		} else {
			orderStr += " DESC"
		}
		db = db.Order(orderStr)
	}

	var count int64
	var offset int

	db.Model(result).Count(&count)
	if p.Start == 1 {
		offset = 0
	} else {
		offset = (p.Start - 1) * p.Size
	}

	db.Limit(p.Size).Offset(offset).Find(result)
	//<-done

	paginator := OpenPaginator{
		PageInfo: openPageMeta{
			Total: count,
			Size:  p.Size,
			Start: p.Start,
		},
		Items: result,
	}

	return &paginator
}

// Pagination
//
//	@Description: 分页
//	@param option 分页参数
//	@param result  分页数据
//	@param args 数据过滤所需的额外参数
//	@param callback 数据过滤
//	@return *types.Response
func OpenPagination(option *OpenPageOption, result interface{}, callback func(result interface{}, arg []interface{}) interface{}, args ...interface{}) *OpenPaginator {
	paginator := openPaging(option, result)
	paginator.Items = callback(paginator.Items, args)
	return paginator
}

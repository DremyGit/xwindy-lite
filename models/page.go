package models

import (
	"math"
	"strconv"
)

type Pagination struct {
	Page       int
	PerPage    int
	TotalCount int
	Offset     int
}

func NewPagination(page, perPage int) *Pagination {
	var pagination Pagination
	pagination.Page = page
	pagination.PerPage = perPage
	pagination.Offset = (pagination.Page - 1) * pagination.PerPage
	return &pagination
}

func (p *Pagination) IsFirst() bool {
	return p.Page == 1
}

func (p *Pagination) IsLast() bool {
	pageCount := p.GetPageCount()
	return p.Page == pageCount
}

func (p *Pagination) HasNext() bool {
	return !p.IsLast()
}

func (p *Pagination) HasPrev() bool {
	return !p.IsFirst()
}

func (p *Pagination) GetPageCount() int {
	return int(math.Ceil(float64(p.TotalCount) / float64(p.PerPage)))
}

func (p Pagination) ToQueryString(page int) string {
	return "page=" + strconv.Itoa(page) + "&per_page=" + strconv.Itoa(p.PerPage)
}

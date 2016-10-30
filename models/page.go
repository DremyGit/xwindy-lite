package models

import (
	"math"
	"strconv"
)

// Pagination is the pagination object
type Pagination struct {
	Page       int
	PerPage    int
	TotalCount int
	Offset     int
}

// NewPagination create new pagination
func NewPagination(page, perPage int) *Pagination {
	var pagination Pagination
	pagination.Page = page
	pagination.PerPage = perPage
	pagination.Offset = (pagination.Page - 1) * pagination.PerPage
	return &pagination
}

// IsFirst is the page the first one
func (p *Pagination) IsFirst() bool {
	return p.Page == 1
}

// IsLast is the page the last one
func (p *Pagination) IsLast() bool {
	pageCount := p.GetPageCount()
	return p.Page == pageCount
}

// HasNext does it has next page
func (p *Pagination) HasNext() bool {
	return !p.IsLast()
}

// HasPrev does it has previous page
func (p *Pagination) HasPrev() bool {
	return !p.IsFirst()
}

// GetPageCount get the total page count of the results
func (p *Pagination) GetPageCount() int {
	return int(math.Ceil(float64(p.TotalCount) / float64(p.PerPage)))
}

// ToQueryString change the pagination to query string
func (p Pagination) ToQueryString(page int) string {
	return "page=" + strconv.Itoa(page) + "&per_page=" + strconv.Itoa(p.PerPage)
}

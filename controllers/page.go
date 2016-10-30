package controllers

import (
	"strconv"

	"strings"

	"github.com/dremygit/xwindy-lite/models"
)

// ParsePagination to parse pagination params in query string
func (c *BaseController) ParsePagination() *models.Pagination {
	var err error
	var page, perPage int

	if page, err = strconv.Atoi(c.Ctx.Input.Query("page")); err != nil {
		page = 1
	}

	perPage, _ = strconv.Atoi(c.Ctx.Input.Query("per_page"))
	if perPage == 0 {
		perPage = 1000
	}

	return models.NewPagination(page, perPage)
}

type link struct {
	Href string
	Rel  string
}

func (l link) ToString() string {
	return "<" + l.Href + ">; rel=\"" + l.Rel + "\""
}

// SuccessWithPagination set the pagination info to the header and response data
func (c *BaseController) SuccessWithPagination(code int, data interface{}, p *models.Pagination) {

	var linkFirst, linkPrev, linkNext, linkLast link
	var links []string

	baseURL := API_BASE + c.Ctx.Input.URL()

	if !p.IsFirst() {
		linkFirst.Href = baseURL + "?" + p.ToQueryString(1)
		linkFirst.Rel = "first"
		links = append(links, linkFirst.ToString())
	}

	if p.HasPrev() {
		linkPrev.Href = baseURL + "?" + p.ToQueryString(p.Page-1)
		linkPrev.Rel = "prev"
		links = append(links, linkPrev.ToString())
	}

	if p.HasNext() {
		linkNext.Href = baseURL + "?" + p.ToQueryString(p.Page+1)
		linkNext.Rel = "next"
		links = append(links, linkNext.ToString())
	}

	if !p.IsLast() {
		linkLast.Href = baseURL + "?" + p.ToQueryString(p.GetPageCount())
		linkLast.Rel = "last"
		links = append(links, linkLast.ToString())
	}

	if p.Page > p.GetPageCount() {
		links = []string{""}
	}

	c.Ctx.Output.Header("X-Total-Count", strconv.Itoa(p.TotalCount))
	c.Ctx.Output.Header("Link", strings.Join(links, ", "))

	c.Success(code, data)
}

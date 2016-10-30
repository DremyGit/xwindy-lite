package controllers

import (
	"strconv"

	"github.com/dremygit/xwindy-lite/models"
	"github.com/jinzhu/copier"
)

type NewsController struct {
	BaseController
}

// GetNewsList get news list
// @Title GetNewsList
// @Description GetNewsList
// @Success 200 {object} []models.NewsListResource
// @Param page query int false "Page Number"
// @Param per_page query int false "Per Page"
// @router / [get]
func (c *NewsController) GetNewsList() {

	p := c.ParsePagination()
	newsDBList, _, err := models.GetNewsList(p)
	if err != nil {
		c.Failure(500, err.Error())
		return
	}

	newsList := []*models.NewsListResource{}
	copier.Copy(&newsList, &newsDBList)

	c.SuccessWithPagination(200, newsList, p)
}

// GetNewsByID Get news by id
// @Title GetNewsByID
// @Description GetNewsByID
// @Param newsid path int true "News ID"
// @Success 200 {object} models.NewsDetailResource
// @router /:newsid [get]
func (c *NewsController) GetNewsByID() {

	id, err := strconv.Atoi(c.Ctx.Input.Param(":newsid"))
	if err != nil {
		c.Failure(400, "ID 错误")
		return
	}

	var newsDB models.News
	err = newsDB.GetByID(id)
	if err != nil {
		c.Failure(404, "新闻不存在")
		return
	}

	var news models.NewsDetailResource
	copier.Copy(&news, &newsDB)
	c.Success(200, news)
}

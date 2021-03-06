package controllers

import (
	"strconv"

	. "github.com/bitly/go-simplejson"

	"github.com/dremygit/xwindy-lite/models"
)

// CommentController handle the request of comment
type CommentController struct {
	BaseController
}

// GetCommentListByNewsID to get comment list by news id
// @Title GetCommentListByNewsID
// @Description GetCommentListByNewsID
// @Param newsid path int true "News ID"
// @Param page query int false "Page Number"
// @Param per_page query int false "Per Page"
// @Success 200 {object} []models.CommentResource
// @router /:newsid/comments [get]
func (c *CommentController) GetCommentListByNewsID() {

	newsID, err := strconv.Atoi(c.Ctx.Input.Param(":newsid"))
	if err != nil {
		c.Failure(400, "ID 错误")
		return
	}

	p := c.ParsePagination()

	commentDBList, err := models.GetCommentListByNewsID(newsID, p)
	if err != nil {
		c.Failure(500, err.Error())
		return
	}

	comments := []*models.CommentResource{}
	for _, commentDB := range commentDBList {
		comments = append(comments, commentDB.ToResource())
	}
	c.SuccessWithPagination(200, comments, p)
}

// CreateComment to create comment to the news
// @Title CreateComment
// @Description Create comment to the news
// @Param newsid path int true "News ID"
// @Param body body models.CommentPayload true body
// @Success 201 {object} models.CommentResource
// @router /:newsid/comments [post]
func (c *CommentController) CreateComment() {
	newsID, error := strconv.Atoi(c.Ctx.Input.Param(":newsid"))
	if error != nil {
		c.Failure(400, "ID 错误")
		return
	}

	token, err := c.ParseToken()
	if err != nil {
		c.Failure(401, err.Error())
		return
	}
	sno := token["sno"].(string)

	js, err := NewJson(c.Ctx.Input.RequestBody)
	if err != nil {
		c.Failure(400, "JSON 格式错误")
		return
	}

	content, err := js.Get("content").String()
	if err != nil || len(content) == 0 {
		c.Failure(400, "请填写内容")
		return
	}

	var commentDB models.Comment
	if err := commentDB.Create(newsID, sno, content); err != nil {
		c.Failure(500, "添加评论失败")
		return
	}

	comment := commentDB.ToResource()
	c.Success(201, &comment)
}

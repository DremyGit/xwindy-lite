package models

import (
	"time"

	"github.com/jinzhu/copier"
)

// Comment is Database model
type Comment struct {
	ID      int       `json:"id"      orm:"size(6);column(id);pk;auto"`
	News    *News     `json:"news_id" orm:"rel(fk)"`
	User    *User     `json:"user"    orm:"rel(fk)"`
	Time    time.Time `json:"time"    orm:"auto_now_add"`
	Content string    `json:"content" orm:"type(text)"`
}

// CommentResource is the resource of comment
type CommentResource struct {
	ID      int       `json:"id"`
	NewsID  int       `json:"news_id"`
	User    string    `json:"user"`
	Time    time.Time `json:"time"`
	Content string    `json:"content"`
}

// CommentPayload is the request payload
type CommentPayload struct {
	Content string `json:"content"`
}

// ToResource change Comment to CommentResource
func (comment Comment) ToResource() *CommentResource {
	var resource CommentResource
	copier.Copy(&resource, &comment)
	resource.NewsID = comment.News.ID
	resource.User = comment.User.Nickname
	return &resource
}

// GetCommentListByNewsID get comment list by news ID
func GetCommentListByNewsID(newsID int, p *Pagination) ([]*Comment, error) {
	var comments []*Comment
	_, err := o.QueryTable("comment").
		Filter("news_id", newsID).
		RelatedSel("user").
		Limit(p.PerPage, p.Offset).
		All(&comments)
	if err != nil {
		return nil, err
	}

	totalCount, _ := o.QueryTable("comment").Filter("news_id", newsID).Count()
	p.TotalCount = int(totalCount)

	return comments, nil
}

// Create create comment in database
func (comment *Comment) Create(newsID int, sno, content string) error {
	var news News
	var user User

	news.ID = newsID
	user.Sno = sno

	comment.News = &news
	comment.User = &user
	comment.Content = content

	o.Begin()
	if _, err := o.Insert(comment); err != nil {
		o.Rollback()
		return err
	}

	if err := news.UpdateCommentCount(); err != nil {
		o.Rollback()
		return err
	}

	o.Commit()
	o.QueryTable("comment").Filter("id", comment.ID).RelatedSel("user").One(comment)
	return nil
}

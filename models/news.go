package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// News is database model
type News struct {
	ID           int       `json:"id"               orm:"size(6);column(id);pk;auto"`
	Title        string    `json:"title"            orm:"size(255)"`
	Time         time.Time `json:"time"             orm:"auto_now_add"`
	Summary      string    `json:"summary"          orm:"size(255)"`
	Content      string    `json:"content"          orm:"size(255);type(text)"`
	SourceURL    string    `json:"source_url"       orm:"size(255);column(source_url)"`
	ClickCount   int       `json:"click_count"      orm:"size(6)"`
	CommentCount int       `json:"comment_count"    orm:"size(6)"`
	Date         time.Time `json:"date"             orm:"type(date)"`
}

// NewsListResource is the resource of news list
type NewsListResource struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Time         time.Time `json:"time"`
	Summary      string    `json:"summary"`
	SourceURL    string    `json:"source_url"`
	ClickCount   int       `json:"click_count"`
	CommentCount int       `json:"comment_count"`
}

// NewsDetailResource is the resource of news detail
type NewsDetailResource struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Time         time.Time `json:"time"`
	Content      string    `json:"content"`
	SourceURL    string    `json:"source_url"`
	ClickCount   int       `json:"click_count"`
	CommentCount int       `json:"comment_count"`
}

// GetNewsList to get news list
func GetNewsList(p *Pagination) ([]*News, int64, error) {
	var news []*News
	total, err := o.QueryTable("news").Limit(p.PerPage, p.Offset).All(&news)
	if err != nil {
		return nil, 0, err
	}

	totalCount, _ := o.QueryTable("news").Count()
	p.TotalCount = int(totalCount)
	return news, total, nil
}

// GetByID to get news by id
func (news *News) GetByID(id int) error {
	news.ID = id
	if err := o.Read(news); err != nil {
		return err
	}
	return nil
}

// Create news
func (news *News) Create() error {
	_, err := o.Insert(news)
	return err
}

// UpdateCommentCount to update comment_count in news table to match the comment count
func (news *News) UpdateCommentCount() error {
	sql := "UPDATE news SET comment_count = (" +
		"SELECT COUNT(0) FROM comment WHERE news_id = ?" +
		") WHERE id = ?"
	_, err := o.Raw(sql, news.ID, news.ID).Exec()
	return err
}

// IncreaseClickCount to increase click_count of the news
func (news *News) IncreaseClickCount() error {
	_, err := o.QueryTable("news").Filter("id", news.ID).Update(orm.Params{
		"click_count": orm.ColValue(orm.ColAdd, 1),
	})
	return err
}

// IsNewsExisted lookup the repeat news
func (news *News) IsNewsExisted() (bool, error) {
	count, err := o.QueryTable("news").Filter("title", news.Title).Filter("date", news.Date).Count()
	if err != nil {
		return false, err
	}
	return count != 0, nil
}

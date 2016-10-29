package models

import (
	"time"
)

// News 新闻表模型
type News struct {
	ID           int       `json:"id"               orm:"size(6);column(id);pk;auto"`
	Title        string    `json:"title"            orm:"size(255)"`
	Time         time.Time `json:"time"`
	Summary      string    `json:"summary"          orm:"size(255)"`
	Content      string    `json:"content"          orm:"size(255);type(text)"`
	SourceURL    string    `json:"source_url"       orm:"size(255);column(source_url)"`
	ClickCount   int       `json:"click_count"      orm:"size(6)"`
	CommentCount int       `json:"comment_count"    orm:"size(6)"`
}

type NewsListResource struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Time         time.Time `json:"time"`
	Summary      string    `json:"summary"`
	SourceURL    string    `json:"source_url"`
	ClickCount   int       `json:"click_count"`
	CommentCount int       `json:"comment_count"`
}

type NewsDetailResource struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Time         time.Time `json:"time"`
	Content      string    `json:"content"`
	SourceURL    string    `json:"source_url"`
	ClickCount   int       `json:"click_count"`
	CommentCount int       `json:"comment_count"`
}

func GetAllNews() ([]*News, int64, error) {
	var news []*News
	total, err := o.QueryTable("news").All(&news)
	if err != nil {
		return nil, 0, err
	}
	return news, total, nil
}

func (news *News) GetByID(id int) error {
	news.ID = id
	if err := o.Read(news); err != nil {
		return err
	}
	return nil
}

package routes

import (
	"fmt"
	"time"
)

// Post contains all data related to user's article posts.
type Post struct {
	ID          int64     `json:"id"`
	Author      User      `json:"author"`
	IsPublished bool      `json:"is_published"`
	PublishedAt time.Time `json:"published_at"`
	ModifiedAt  time.Time `json:"modified_at"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Abstract    string    `json:"excerpt"`
	URL         string    `json:"url"`
}

func (p *Post) FormatPostPublishedDate() string {
	year, month, day := p.PublishedAt.Date()
	return fmt.Sprintf("%v %d %d", month, day, year)
}

func (p *Post) FormatAbstract() string {
	if p.Abstract != "" {
		return p.Abstract
	}
	return p.Content
}

package routes

import (
	"fmt"
	"github.com/9uuso/excerpt"
	"time"
)

const (
	ExcerptWordLength = 20
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

func (post *Post) FormatPublishedAt() string {
	year, month, day := post.PublishedAt.Date()
	return fmt.Sprintf("%v %d %d", month, day, year)
}

func (post *Post) FormatModifiedAt() string {
	year, month, day := post.ModifiedAt.Date()
	return fmt.Sprintf("%v %d %d", month, day, year)
}

func (post *Post) FormatAbstract() string {
	if post.Abstract != "" {
		return post.Abstract
	}
	post.Abstract = excerpt.Make(post.Content, ExcerptWordLength)
	return post.Abstract
}

package routes

import (
	"fmt"
	"time"
)

const (
	AbstractWordLength = 80
	IndexPageTitle     = "Rong's Blog"
	PostsPath          = "./content/"
)

// Post contains all data related to user's article posts.
type Post struct {
	ID          int64     `json:"id"`
	Author      User      `json:"author"`
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
	if len(p.Content) > AbstractWordLength {
		return p.Content[:AbstractWordLength]
	}
	return p.Content
}

type IndexPageData struct {
	Title      string `json:"title"`
	TotalPosts int    `json:"total_posts"`
	Posts      []Post `json:"posts"`
	Status     string `json:"status"`
}

func GetIndexPageData() IndexPageData {
	posts := getBlogPosts()

	return IndexPageData{
		Title:      IndexPageTitle,
		TotalPosts: len(posts),
		Posts:      posts,
		Status:     "",
	}
}

func getBlogPosts() []Post {
	posts := make([]Post, 0)

	// TODO: Extract all the posts under directory "../content/"

	return posts
}

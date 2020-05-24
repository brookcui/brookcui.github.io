package routes

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

const (
	AbstractWordLength = 80
	IndexPageTitle = "Rong's Blog"
	PostsPath = "./content/"
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

type Nav struct {
	Topic string  `json:"topic"`
	URL   string  `json:"url"`
}

type IndexPageData struct {
	Title      string `json:"title"`
	TotalPosts int    `json:"total_posts"`
	Navs       []Nav  `json:"navs"`
	Posts      []Post `json:"posts"`
	Status     string `json:"status"`
}

func GetIndexPageData() IndexPageData {
	navs := getBlogNavs()
	posts := getBlogPosts()

	return IndexPageData{
		Title:      IndexPageTitle,
		TotalPosts: len(posts),
		Navs:       navs,
		Posts:      posts,
		Status:     "",
	}
}

func getBlogNavs() []Nav {
	return []Nav{
		{
			Topic: "Home",
			URL:   "/",
		},
		{
			Topic: "About",
			URL:   "/about",
		},
		{
			Topic: "Github",
			URL:   "https://github.com/brookcui",
		},
	}
}

func getBlogPosts() []Post {
	posts := make([]Post, 0)

	err := filepath.Walk(PostsPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			post := Post{
				ID:          0,
				Author:      User{
					ID:        0,
					Username:  "Rong Cui",
					Email:     "brookcui97@gmail.com",
					CreatedAt: time.Now(),
				},
				PublishedAt: time.Now(),
				ModifiedAt:  time.Now(),
				Title:       info.Name(),
				Content:     "",
				Abstract:    "",
				URL:         "",
			}
			posts = append(posts, post)
			return nil
		})
	if err != nil {
		log.Println(err)
	}

	return posts
}

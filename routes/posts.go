package routes

import (
	"time"
)

const (
	title = "Rong's Blog"
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
		Title:      title,
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
	return []Post{}
}

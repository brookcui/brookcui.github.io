package routes

type Page struct {
	Title string `json:"title"`
	Body  []byte `json:"body"`
}

type IndexPage struct {
	Page Page `json:"page"`
	TotalPosts int    `json:"total_posts"`
	Posts      []Post `json:"posts"`
	Status     string `json:"status"`
}

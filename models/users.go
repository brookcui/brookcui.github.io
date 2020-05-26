package routes

import (
	"fmt"
	"time"
)

type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Emails    []string   `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	URL       string    `json:"url"`
}

func (u *User) FormatUserInformation() string {
	return fmt.Sprintf("%s", u.Username)
}

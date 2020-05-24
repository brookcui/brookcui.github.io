package routes

import (
	"fmt"
	"time"
)

type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *User) FormatUserInformation() string {
	return fmt.Sprintf("%s", u.Username)
}

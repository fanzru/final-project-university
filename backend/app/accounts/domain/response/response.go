package response

import (
	"time"

	"github.com/volatiletech/null/v9"
)

type UserLoginRes struct {
	AccessToken string `json:"access_token"`
}

type UserProfileRes struct {
	ID        int64       `json:"id"`
	Name      string      `json:"name" `
	Email     string      `json:"email"`
	PhotoUrl  null.String `json:"photo_url"`
	CreatedAt time.Time   `json:"created_at"`
	DeletedAt null.Time   `json:"deleted_at"`
}

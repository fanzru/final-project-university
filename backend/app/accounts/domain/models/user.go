package models

import (
	"time"

	"github.com/volatiletech/null/v9"
)

type User struct {
	ID        int64       `json:"id" gorm:"id"`
	Name      string      `json:"name" gorm:"name"`
	Email     string      `json:"email" gorm:"email"`
	Password  string      `json:"password" gorm:"password"`
	PhotoUrl  null.String `json:"photo_url" gorm:"photo_url"`
	CreatedAt time.Time   `json:"created_at" gorm:"created_at"`
	DeletedAt null.Time   `json:"deleted_at" gorm:"deleted_at"`
}

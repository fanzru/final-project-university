package models

import (
	"time"

	"github.com/volatiletech/null/v9"
)

type User struct {
	Id        int64       `json:"id" gorm:"id"`
	Name      string      `json:"name" gorm:"name"`
	Email     string      `json:"email" gorm:"email"`
	Password  string      `json:"password" gorm:"password"`
	PhotoUrl  null.String `json:"photo_url" gorm:"photo_url"`
	CreatedAt time.Time   `json:"created_at" gorm:"created_at"`
	DeletedAt null.Time   `json:"deleted_at" gorm:"deleted_at"`
}

type SentencesLabel struct {
	Id          int64  `json:"id" gorm:"id"`
	PaperId     int64  `json:"paper_id" gorm:"paper_id"`
	Head        string `json:"head" gorm:"head"`
	Text        string `json:"text" gorm:"text"`
	IsImportant bool   `json:"is_important" gorm:"is_important"`
}

type PapersUser struct {
	Id             int64             `json:"id" gorm:"id"`
	UserId         int64             `json:"user_id" gorm:"user_id"`
	PaperName      string            `json:"paper_name" gorm:"paper_name"`
	LinkPdf        string            `json:"link_pdf" gorm:"link_pdf"`
	IsDone         bool              `json:"is_done" gorm:"is_done"`
	CreatedAt      time.Time         `json:"created_at" gorm:"created_at"`
	DeletedAt      *time.Time        `json:"deleted_at,omitempty" gorm:"deleted_at"`
	SentencesLabel *[]SentencesLabel `json:"sentences_labels" gorm:"foreignKey:paper_id;references:id"`
}

type Profile struct {
	Id          int64         `json:"id" gorm:"id"`
	Name        string        `json:"name" gorm:"name"`
	Email       string        `json:"email" gorm:"email"`
	PhotoUrl    null.String   `json:"photo_url" gorm:"photo_url"`
	CreatedAt   time.Time     `json:"created_at" gorm:"created_at"`
	DeletedAt   null.Time     `json:"deleted_at" gorm:"deleted_at"`
	PapersUsers *[]PapersUser `json:"papers_users" gorm:"foreignKey:user_id;references:id"`
}

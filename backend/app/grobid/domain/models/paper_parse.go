package models

import "time"

type SentencesLabel struct {
	Id          int64  `json:"id" gorm:"id"`
	PaperId     int64  `json:"paper_id" gorm:"paper_id"`
	Head        string `json:"head" gorm:"head"`
	Text        string `json:"text" gorm:"text"`
	IsImportant bool   `json:"is_important" gorm:"is_important"`
}

type PapersUsers struct {
	Id          int64      `json:"id" gorm:"id"`
	UserId      int64      `json:"user_id" gorm:"user_id"`
	DomainPaper string     `json:"domain_paper" gorm:"domain_paper"`
	PaperName   string     `json:"paper_name" gorm:"paper_name"`
	LinkPdf     string     `json:"link_pdf" gorm:"link_pdf"`
	IsDone      bool       `json:"is_done" gorm:"is_done"`
	CreatedAt   time.Time  `json:"created_at" gorm:"created_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" gorm:"deleted_at"`
}

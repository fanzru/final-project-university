package models

import (
	"time"

	"github.com/volatiletech/null/v9"
)

type TuwienSummaLmjm struct {
	Id        int64     `json:"id" gorm:"id"`
	PaperId   int64     `json:"paper_id" gorm:"paper_id"`
	Sent      string    `json:"sent" gorm:"sent"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	DeletedAt null.Time `json:"deleted_at" gorm:"deleted_at"`
}

type TuwienSummaBM25 struct {
	Id        int64     `json:"id" gorm:"id"`
	PaperId   int64     `json:"paper_id" gorm:"paper_id"`
	Sent      string    `json:"sent" gorm:"sent"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	DeletedAt null.Time `json:"deleted_at" gorm:"deleted_at"`
}

type DatasetPaper struct {
	Id        int64     `json:"id"`
	UserId    int64     `json:"user_id"`
	PaperName string    `json:"paper_name"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt null.Time `json:"deleted_at"`
}

type GrobidExtraction struct {
	Id                int64     `json:"id"`
	PaperId           int64     `json:"paper_id"`
	ExtractionHeading string    `json:"extraction_heading"`
	ActualHeading     string    `json:"actual_heading"`
	Text              string    `json:"text"`
	CreatedAt         time.Time `json:"created_at"`
	DeletedAt         null.Time `json:"deleted_at"`
}

package models

type SentencesLabel struct {
	PaperId     int64  `json:"paper_id" gorm:"paper_id"`
	Head        string `json:"head" gorm:"head"`
	Text        string `json:"text" gorm:"text"`
	IsImportant bool   `json:"is_important" gorm:"is_important"`
}

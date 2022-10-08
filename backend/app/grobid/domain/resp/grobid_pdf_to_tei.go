package resp

import (
	"backend/app/grobid/domain/outbound"
	"strings"
)

type PDFToTEI struct {
	Body []Body `json:"body"`
}

type Body struct {
	Head      string     `json:"head"`
	Sentences []Sentence `json:"Sentences"`
}

type Sentence struct {
	Text        string `json:"text"`
	IsImportant bool   `json:"is_important"`
}

func (b *PDFToTEI) MapToTEIParse(data *outbound.TEI) {
	for _, v := range data.Text.Body.Div {
		sentences := []Sentence{}
		for _, s := range v.P {
			strSplit := strings.Split(s.Text, ". ")
			for _, s2 := range strSplit {
				sentences = append(sentences, Sentence{
					Text:        s2 + ".",
					IsImportant: false,
				})
			}

		}

		b.Body = append(b.Body, Body{
			Head:      v.Head.Text,
			Sentences: sentences,
		})
	}
}

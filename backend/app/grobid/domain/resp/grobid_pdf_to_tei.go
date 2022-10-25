package resp

import (
	"backend/app/grobid/domain/outbound"
	"strings"
)

type PDFToTEI struct {
	LenHead int    `json:"len_head"`
	Body    []Body `json:"body"`
}

type Body struct {
	HeadKey   int        `json:"head_key"`
	Head      string     `json:"head"`
	Sentences []Sentence `json:"sentences"`
}

type Sentence struct {
	Text        string `json:"text"`
	IsImportant bool   `json:"is_important"`
}

func (b *PDFToTEI) MapToTEIParse(data *outbound.TEI) {
	for i, v := range data.Text.Body.Div {
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
			HeadKey:   i + 1,
			Head:      v.Head.Text,
			Sentences: sentences,
		})
	}
	b.LenHead = len(data.Text.Body.Div)
}

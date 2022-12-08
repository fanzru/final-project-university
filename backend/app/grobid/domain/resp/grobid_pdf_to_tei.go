package resp

import (
	"backend/app/grobid/domain/models"
	"backend/app/grobid/domain/outbound"
	"strings"
)

type CSVresp struct {
	PaperDetail    models.PapersUsers      `json:"paper_detail"`
	SentencesLabel []models.SentencesLabel `json:"sentences_label"`
}
type PDFToTEI struct {
	LinkPdf string                   `json:"link_pdf"`
	PaperId int64                    `json:"paper_id"`
	LenHead int                      `json:"len_head"`
	Body    []Body                   `json:"body"`
	Lmjm    []models.TuwienSummaLmjm `json:"lmjm"`
	BM25    []models.TuwienSummaBM25 `json:"bm25"`
}

type Body struct {
	HeadKey   int         `json:"head_key"`
	Head      string      `json:"head"`
	Sentences []Sentence  `json:"sentences"`
	Paragraph []Paragraph `json:"paragraph"`
}

type Sentence struct {
	SentID      int64  `json:"sent_id"`
	Text        string `json:"text"`
	IsImportant bool   `json:"is_important"`
}

type Paragraph struct {
	ID   int64  `json:"sent_id"`
	Text string `json:"text"`
}

func (b *PDFToTEI) MapToTEIParse(data *outbound.TEI) {
	for i, v := range data.Text.Body.Div {
		sentences := []Sentence{}
		for _, s := range v.P {
			strSplit := strings.Split(s.Text, ". ")
			for _, s2 := range strSplit {
				sentences = append(sentences, Sentence{
					// can customize to add dot or not in Text field.
					Text:        s2,
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

func (b *PDFToTEI) MapToTEIParseParagraft(data *outbound.TEI) {
	for i, v := range data.Text.Body.Div {
		paragraft := []Paragraph{}
		for _, s := range v.P {
			paragraft = append(paragraft, Paragraph{
				Text: s.Text,
			})
		}

		b.Body = append(b.Body, Body{
			HeadKey:   i + 1,
			Head:      v.Head.Text,
			Paragraph: paragraft,
		})
	}
	b.LenHead = len(data.Text.Body.Div)
}

func (b *PDFToTEI) MapToResponse(papersUsers *models.PapersUsers, sentencesLabel *[]models.SentencesLabel) {
	if papersUsers == nil || sentencesLabel == nil {
		return
	}
	b.LinkPdf = papersUsers.LinkPdf
	b.PaperId = papersUsers.Id
	headKey := []string{}
	for _, sent := range *sentencesLabel {
		headKey = b.appendIfMissing(headKey, sent.Head)
	}
	b.LenHead = len(headKey)
	for i, head := range headKey {
		body := Body{}
		for _, sentence := range *sentencesLabel {
			if sentence.Head == head {
				body.Head = head
				body.HeadKey = i + 1
				body.Sentences = append(body.Sentences, Sentence{
					SentID:      sentence.Id,
					Text:        sentence.Text,
					IsImportant: sentence.IsImportant,
				})
			}
		}
		b.Body = append(b.Body, body)
	}
}

func (b *PDFToTEI) appendIfMissing(slice []string, i string) []string {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}

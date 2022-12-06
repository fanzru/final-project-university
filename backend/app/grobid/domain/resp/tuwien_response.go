package resp

type DataPaperArtuSummary struct {
	DocId     string      `json:"doc_id"`
	Summaries []Summaries `json:"summaries"`
}

type Summaries struct {
	Method       string         `json:"method"`
	Summary      []string       `json:"summary"`
	ZonesSummary []ZonesSummary `json:"zones_summary"`
}

type ZonesSummary struct {
	Category        string   `json:"category"`
	CategorySummary []string `json:"category_summary"`
}

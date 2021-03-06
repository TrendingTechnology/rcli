package v4

type IndexResult struct {
	Themes    []string `json:"themes"`
	Sentiment string   `json:"sentiment"`
	Entities  []Entity `json:"entities"`
}

type SearchResult struct {
	Total     int        `json:"total"`
	Documents []Document `json:"matches"`
}

type Document struct {
	Text     string   `json:"text"`
	Entities []Entity `json:"entities"`
}

type Entity struct {
	Title           string   `json:"title"`
	Classifications []string `json:"classifications"`
}

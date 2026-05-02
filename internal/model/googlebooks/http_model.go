package googlebooks

type VolumenInfo struct {
	Title         string   `json:"title"`
	Subtitle      string   `json:"subtitle,omitempty"`
	Authors       []string `json:"authors,omitempty"`
	Categories    []string `json:"categories,omitempty"`
	PublishedDate string   `json:"publishedDate,omitempty"`
	PageCount     int      `json:"pageCount,omitempty"`
	PreviewLink   string   `json:"previewLink,omitempty"`
}

type Item struct {
	SelfLink    string      `json:"selfLink"`
	VolumenInfo VolumenInfo `json:"volumeInfo"`
	SaleInfo    SaleInfo    `json:"saleInfo"`
}

type SaleInfo struct {
	BuyLink string `json:"buyLink,omitempty"`
}

type GoogleBooksResponse struct {
	Items []Item `json:"items"`
}

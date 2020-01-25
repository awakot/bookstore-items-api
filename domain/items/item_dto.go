package items

type Item struct {
	ID                string        `json:"id"`
	Author            int64         `json:"author"`
	Title             string        `json:"title"`
	Description       Description   `json:"description"`
	Pictures          []ItemPicture `json:"pictures"`
	VideoURL          string        `json:"video_url"`
	Price             float32       `json:"price"`
	AvailableQuantity int           `json:"available_quantity"`
	SoldQuantity      int           `json:"sold_quantity"`
	Status            string        `json:"status"`
}

type Description struct {
	PlainText string `json:"plain_text"`
	HTML      string `json:"html"`
}
type ItemPicture struct {
	ID  int64  `json:"id"`
	URL string `json:"url"`
}

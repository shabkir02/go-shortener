package models

type ShortURLStruct struct {
	HashURL string `json:"hashUrl"`
	URL     string `json:"url"`
}

type AllURLsStruct struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

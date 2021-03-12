package domain

type WordRequest struct {
	EnglishWord string `json:"english-word"`
}

type WordResponse struct {
	GopherWord string `json:"gopher-word"`
}

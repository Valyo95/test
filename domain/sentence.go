package domain

type SentenceRequest struct {
	EnglishSentence string `json:"english-sentence"`
}

type SentenceResponse struct {
	GopherSentence string `json:"gopher-sentence"`
}

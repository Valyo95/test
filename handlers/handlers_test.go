package handlers

import (
	"net/http"
	"testing"
)

func Test_testHandler(t *testing.T) {
	tests := []TestHandlerStruct{
		// Testing WordHandler successes
		{WordHandler, clearHistory, map[string]string{"english-word": "alien"}, "application/json", http.StatusOK, `{"gopher-word":"galien"}`},
		{WordHandler, clearHistory, map[string]string{"english-word": "apple"}, "application/json", http.StatusOK, `{"gopher-word":"gapple"}`},
		{WordHandler, clearHistory, map[string]string{"english-word": "xray"}, "application/json", http.StatusOK, `{"gopher-word":"gexray"}`},
		{WordHandler, clearHistory, map[string]string{"english-word": "chair"}, "application/json", http.StatusOK, `{"gopher-word":"airchogo"}`},
		{WordHandler, clearHistory, map[string]string{"english-word": "square"}, "application/json", http.StatusOK, `{"gopher-word":"aresquogo"}`},
		// Testing WordHandler errors
		{WordHandler, clearHistory, map[string]string{"english-word": "square"}, "application/XML", http.StatusUnsupportedMediaType, `{"error":"Content Type is not application/json"}`},
		{WordHandler, clearHistory, map[string]string{"undefined": "square"}, "application/json", http.StatusBadRequest, `{"error":"Could not decode body: json: unknown field \"undefined\""}`},
		{WordHandler, clearHistory, map[string]string{"english-word": "Hey there"}, "application/json", http.StatusBadRequest, `{"error":"Error, '/word/' must be provided with a single word"}`},

		// Testing SentenceHandler successes
		{SentenceHandler, clearHistory, map[string]string{"english-sentence": "Hey there"}, "application/json", http.StatusOK, `{"gopher-sentence":"eyHogo erethogo"}`},
		{SentenceHandler, clearHistory, map[string]string{"english-sentence": "Hey there!"}, "application/json", http.StatusOK, `{"gopher-sentence":"eyHogo erethogo !"}`},
		{SentenceHandler, clearHistory, map[string]string{"english-sentence": "hey"}, "application/json", http.StatusOK, `{"gopher-sentence":"eyhogo"}`},
		// Testing SentenceHandler errors
		{SentenceHandler, clearHistory, map[string]string{"english-sentence": "hey"}, "text/csv", http.StatusUnsupportedMediaType, `{"error":"Content Type is not application/json"}`},
		{SentenceHandler, clearHistory, map[string]string{"sentence": "hey"}, "application/json", http.StatusBadRequest, `{"error":"Could not decode body: json: unknown field \"sentence\""}`},

		// Testing HistoryHandler
		{
			HistoryHandler,
			clearHistory,
			nil,
			"application/json",
			http.StatusOK,
			`{"history":[]}`,
		},
		{
			HistoryHandler,
			func() {
				clearHistory()
				history["square"] = "aresquogo"
				history["alien"] = "galien"
				history["aapple"] = "gapple"
			},
			nil,
			"application/json",
			http.StatusOK,
			`{"history":[{"aapple":"gapple"},{"alien":"galien"},{"square":"aresquogo"}]}`,
		},
		{
			HistoryHandler,
			func() {
				clearHistory()
				history["hye"] = "aresquogo"
				history["all around the world"] = "galien"
			},
			nil,
			"application/json",
			http.StatusOK,
			`{"history":[{"all around the world":"galien"},{"hye":"aresquogo"}]}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.expectedBody, func(t *testing.T, ) {
			testHandler(t, tt)
		})
	}
}

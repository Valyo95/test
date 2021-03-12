package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/valyo95/gopher-translator/domain"
	"github.com/valyo95/gopher-translator/stringutil"
	"github.com/valyo95/gopher-translator/translator"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"regexp"
	"sort"
	"strings"
	"testing"
)

type HistoryResponse struct {
	History []map[string]string `json:"history"`
}

type ErrorResponse struct {
	Message string `json:"error"`
}

var history = make(map[string]string)

func WordHandler(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		createErrorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	var wordRequest domain.WordRequest
	body, _ := ioutil.ReadAll(r.Body)
	reader := strings.NewReader(string(body))

	decoder := json.NewDecoder(reader)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&wordRequest)
	if err != nil {
		createErrorResponse(w, "Could not decode body: "+err.Error(), http.StatusBadRequest)
		return
	}
	// trim leading and trailing white spaces
	wordRequest.EnglishWord = strings.TrimSpace(wordRequest.EnglishWord)

	// check whether body contains multiple words
	if regexp.MustCompile(`\s+`).MatchString(wordRequest.EnglishWord) {
		createErrorResponse(w, "Error, '/word/' must be provided with a single word", http.StatusBadRequest)
		return
	}
	gopherWord := domain.WordResponse{GopherWord: translator.TranslateWord(wordRequest.EnglishWord)}
	// add sentence to history
	history[wordRequest.EnglishWord] = gopherWord.GopherWord
	createResponse(w, gopherWord)
}

func SentenceHandler(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		createErrorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	var sentenceRequest domain.SentenceRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&sentenceRequest)
	if err != nil {
		createErrorResponse(w, "Could not decode body: "+err.Error(), http.StatusBadRequest)
		return
	}
	// trim leading and trailing white spaces
	sentenceRequest.EnglishSentence = strings.TrimSpace(sentenceRequest.EnglishSentence)

	// split sentence into slice of strings
	words := stringutil.SplitSentenceIntoWords(sentenceRequest)
	gopherWords := make([]string, len(words))

	// translate each word
	for i, word := range words {
		englishWord := translator.TranslateWord(word)
		gopherWords[i] = englishWord
	}
	sentence := strings.Join(gopherWords, " ")
	gopherSentence := domain.SentenceResponse{GopherSentence: sentence}

	// add sentence to history
	history[sentenceRequest.EnglishSentence] = gopherSentence.GopherSentence
	createResponse(w, gopherSentence)
}

func HistoryHandler(w http.ResponseWriter, r *http.Request) {
	keys := make([]string, 0, len(history))
	for k := range history {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	var historyResponse = make([]map[string]string, 0, len(history))
	for _, key := range keys {
		entry := map[string]string{key: history[key]}
		historyResponse = append(historyResponse, entry)
	}
	createResponse(w, HistoryResponse{History: historyResponse})
}

func createErrorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	errorResponse := ErrorResponse{Message: message}

	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(true)
	err := encoder.Encode(errorResponse)
	if err != nil {
		log.Printf("Failed to serialize the response: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func createResponse(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(true)
	err := encoder.Encode(response)

	if err != nil {
		log.Printf("Failed to serialize the response: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func clearHistory() {
	for k := range history {
		delete(history, k)
	}
}

// Special structure that can test multiple Handlers
type TestHandlerStruct struct {
	handler            func(w http.ResponseWriter, r *http.Request) // The handler you want to test
	setup              func()                                       // init setup function
	request            map[string]string                            // the request body of the request
	contentType        string                                       // HTTP request content type
	expectedHTTPStatus int                                          // HTTP expected status
	expectedBody       string                                       // the expected response's body
}

/*
	A general method that can test any handler functionality by provided by the TestHandlerStruct
 */
func testHandler(t *testing.T, test TestHandlerStruct) {

	// Do the initialization with the provided setup() method
	test.setup()

	requestBody, err := json.Marshal(test.request)

	// Here I'm only testing the handler and not the routing mechanism
	// This is way the method and url are empty (cause each handler can be set to handle each method and pat)
	req, err := http.NewRequest("", "", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", test.contentType)

	// create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(test.handler)

	// Serve the request
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != test.expectedHTTPStatus {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
		//return
	}

	// Check the response body is what we expect.
	if stringutil.TrimSuffix(rr.Body.String(), "\n") != test.expectedBody {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), test.expectedBody)
		//return
	}
}

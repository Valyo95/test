package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/valyo95/gopher-translator/handlers"
	"net/http"
	"strconv"
)

func main() {
	port := flag.Int("port", 8080, "an integer")

	flag.Parse()
	if *port == 8080 {
		fmt.Printf("Port is not defined. Using default port: %d\n", *port)
	}

	router := mux.NewRouter()

	router.HandleFunc("/word/", handlers.WordHandler).
		Methods("POST").
		Schemes("http")

	router.HandleFunc("/sentence/", handlers.SentenceHandler).
		Methods("POST").
		Schemes("http")

	router.HandleFunc("/history", handlers.HistoryHandler).
		Methods("GET").
		Schemes("http")

	// Start the server.
	fmt.Printf("Starting server on port: %d\n", *port)
	fmt.Printf("Listening...\n")
	http.ListenAndServe(":"+strconv.Itoa(*port), router) // mux.Router now in play
}

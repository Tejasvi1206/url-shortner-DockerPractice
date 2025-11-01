package main

import (
	"fmt"
	"net/http"

	"url-shortner/handlers"
)

func main() {
	http.HandleFunc("/shorten", handlers.ShortURLHandler)
	http.HandleFunc("/redirect/", handlers.RedirectURLHandler)
	http.HandleFunc("/metrics", handlers.MetricsHandler)

	fmt.Println("Starting server on port 3000...")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

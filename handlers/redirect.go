package handlers

import (
	"errors"
	"net/http"
)

func getURL(id string) (string, error) {
	url, ok := URLDB[id]
	if !ok {
		return "", errors.New("URL not found")
	}
	return url.OriginalURL, nil
}

func RedirectURLHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/redirect/"):]
	originalURL, err := getURL(id)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, originalURL, http.StatusFound)
}

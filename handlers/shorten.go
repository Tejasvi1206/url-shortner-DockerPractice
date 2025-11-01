package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"url-shortner/models"
	"url-shortner/utils"
)

var URLDB = make(map[string]models.URL)

func CreateURL(originalURL string) string {
	for _, v := range URLDB {
		if v.OriginalURL == originalURL {
			return v.ShortURL
		}
	}

	shortURL := utils.GenerateShortURL(originalURL)
	id := shortURL
	URLDB[id] = models.URL{
		ID:           id,
		OriginalURL:  originalURL,
		ShortURL:     shortURL,
		CreationDate: time.Now(),
	}
	return shortURL
}

func ShortURLHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	shortURL := CreateURL(data.URL)
	response := struct {
		ShortURL string `json:"short_url"`
	}{
		ShortURL: "http://localhost:3000/redirect/" + shortURL,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"
	"sort"
)

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	domainCount := make(map[string]int)

	for _, v := range URLDB {
		parsed, err := url.Parse(v.OriginalURL)
		if err == nil {
			domainCount[parsed.Host]++
		}
	}

	type kv struct {
		Domain string `json:"domain"`
		Count  int    `json:"count"`
	}
	var sorted []kv
	for k, v := range domainCount {
		sorted = append(sorted, kv{k, v})
	}

	sort.Slice(sorted, func(i, j int) bool { return sorted[i].Count > sorted[j].Count })
	if len(sorted) > 3 {
		sorted = sorted[:3]
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sorted)
}

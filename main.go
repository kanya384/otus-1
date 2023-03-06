package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/health/", successResponseHandler)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}

func successResponseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "OK"})
}

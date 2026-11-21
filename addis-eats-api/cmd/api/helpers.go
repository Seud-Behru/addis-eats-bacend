package main

import (
	"encoding/json"
	"net/http"
)

// writeJSON sends a JSON response with the correct headers and status
func writeJSON(w http.ResponseWriter, status int, data any) {
	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "encoding failure", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(b)
}

// writeError sends a JSON error response
func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}
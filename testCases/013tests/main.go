package main

import (
	"net/http"
	"io"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, `{"alive": true}`)
}

func main() {
	http.HandleFunc("/health-check", HealthCheckHandler)
	http.ListenAndServe(":8080", nil)
}

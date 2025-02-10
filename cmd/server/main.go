package main

import (
	"fmt"
	"net/http"

	"wx-miniprogram-backend/internal/log"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", helloHandler)

	log.Logger.Info().Msg("Server starting on port 8080...")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Logger.Fatal().Err(err).Msg("Server failed to start")
	}
}

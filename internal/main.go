package main

import (
	"net/http"

	"github.com/juniorrodes/exchange/internal/infrastructure/server"
)

func main() {
	router := server.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello world"))
	})

	router.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("só um teste"))
	})

	router.Serve(":8080")
}

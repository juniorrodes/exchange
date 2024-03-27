package main

import (
	"net/http"

	"github.com/juniorrodes/exchange/internal/router"
)

func main() {
	r := router.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello world"))
	})

	r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("só um teste"))
	})

	r.Serve(":8080")
}

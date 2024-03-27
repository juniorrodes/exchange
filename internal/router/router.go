package router

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Router struct {
	mux *http.ServeMux
}

func NewRouter() Router {
	return Router{
		mux: http.NewServeMux(),
	}
}

func (r *Router) Get(path string, handler http.HandlerFunc) {
	methodPath := fmt.Sprintf("GET %s", path)

	r.mux.Handle(methodPath, handler)
}

func (r *Router) Serve(address string) {
	logger := log.New(os.Stdout, "[exchange-server]", log.Ldate|log.Ltime)

	logger.Println("Listening at", address)

	http.ListenAndServe(address, r.mux)
}

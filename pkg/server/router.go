package server

import (
	"fmt"
	"log"
	"net/http"
)

type Router struct {
	mux *http.ServeMux
    logger *log.Logger
}

func NewRouter(logger *log.Logger) Router {
	return Router{
		mux: http.NewServeMux(),
        logger: logger,
	}
}

func (r *Router) Get(path string, handler http.HandlerFunc) {
	methodPath := fmt.Sprintf("GET %s", path)
       
	r.mux.Handle(methodPath, handler)
}

func (r *Router) Post(path string, handler http.HandlerFunc) {
	methodPath := fmt.Sprintf("POST %s", path)
    
	r.mux.Handle(methodPath, handler)
}

func (r *Router) Serve(address string) {
	r.logger.Println("Listening at", address)

	http.ListenAndServe(address, r.mux)
}

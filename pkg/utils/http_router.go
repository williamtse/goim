package utils

import (
	"log"
	"net/http"
)

type middleware func(http.Handler) http.Handler

type Router struct {
	middlewareChain []middleware
	mux             map[string]http.Handler
}

func HttpRouter() *Router {
	return &Router{
		middlewareChain: []middleware{},
		mux:             make(map[string]http.Handler),
	}
}

func (r *Router) Use(m middleware) {
	r.middlewareChain = append(r.middlewareChain, m)
}

func (r *Router) Add(route string, h http.Handler) {
	var mergedHandler = h

	for i := len(r.middlewareChain) - 1; i >= 0; i-- {
		mergedHandler = r.middlewareChain[i](mergedHandler)
	}

	r.mux[route] = mergedHandler
}

func (r *Router) Exclude(route string, h http.Handler) {
	r.mux[route] = h
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	urlPath := req.URL.Path
	handler, ok := r.mux[urlPath]
	if ok {
		handler.ServeHTTP(w, req)
	} else {
		log.Println("ERROR: no handler find:", urlPath)
		http.Error(w, "route error", http.StatusInternalServerError)
	}
}

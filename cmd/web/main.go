package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /{$}", getHome)
	mux.HandleFunc("GET /blog/view/{id}/{$}", getBlogView)
	mux.HandleFunc("GET /blog/create", getBlogCreate)
	mux.HandleFunc("POST /blog/create", postBlogCreate)

	log.Print("starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

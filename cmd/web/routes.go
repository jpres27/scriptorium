package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /{$}", app.getHome)
	mux.HandleFunc("GET /blog/view/{id}/{$}", app.getBlogView)
	mux.HandleFunc("GET /blog/create", app.getBlogCreate)
	mux.HandleFunc("POST /blog/create", app.postBlogCreate)

	return mux
}

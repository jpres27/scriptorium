package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /{$}", app.getHome)
	mux.HandleFunc("GET /text/view/{id}/{$}", app.getBlogView)
	mux.HandleFunc("GET /text/create", app.getBlogCreate)
	mux.HandleFunc("POST /text/create", app.postBlogCreate)

	return app.logRequest(commonHeaders(mux))
}

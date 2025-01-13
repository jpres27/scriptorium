package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"jpp.blog/internal/models"
)

func (app *application) getHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	texts, err := app.texts.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	for _, text := range texts {
		fmt.Fprintf(w, "%+v\n", text)
	}

	/*
		files := []string{
			"./ui/html/base.tmpl.html",
			"./ui/html/partials/nav.tmpl.html",
			"./ui/html/pages/home.tmpl.html",
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.serverError(w, r, err)
			return
		}

		err = ts.ExecuteTemplate(w, "base", nil)
		if err != nil {
			app.serverError(w, r, err)
		}
	*/
}

func (app *application) getBlogView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	text, err := app.texts.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	fmt.Fprintf(w, "%+v", text)
}

func (app *application) getBlogCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new blog post..."))
}

func (app *application) postBlogCreate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	// TODO: Implement this instead of the example
	w.Write([]byte("Save a new blog post..."))
}

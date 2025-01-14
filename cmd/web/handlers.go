package main

import (
	"errors"
	"net/http"
	"strconv"

	"jpp.blog/internal/models"
)

func (app *application) getHome(w http.ResponseWriter, r *http.Request) {
	texts, err := app.texts.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := app.newTemplateData(r)
	data.Texts = texts

	app.render(w, r, http.StatusOK, "home.tmpl.html", data)
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

	data := app.newTemplateData(r)
	data.Text = text

	app.render(w, r, http.StatusOK, "view.tmpl.html", data)
}

func (app *application) getBlogCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new blog post..."))
}

func (app *application) postBlogCreate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	// TODO: Implement this instead of the example
	w.Write([]byte("Save a new blog post..."))
}

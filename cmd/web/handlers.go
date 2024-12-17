package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func getHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	w.Write([]byte("Hello from Snippetbox"))
}

func getBlogView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific blog post with ID %d...", id)
}

func getBlogCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new blog post..."))
}

func postBlogCreate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new blog post..."))
}

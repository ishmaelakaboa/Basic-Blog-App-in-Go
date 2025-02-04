package main

import (
	"html/template"
	"net/http"
	"sync"
)

type Post struct {
	ID      int
	Title   string
	Content string
}

var (
	posts   = []Post{}
	nextID  = 1
	postMux sync.Mutex // Mutex for safe concurrent access
)

func homeHandler(w http.ResponseWriter, r *http.Request){
	tmpl, err := template.ParseFiles("templates/index.html")
}
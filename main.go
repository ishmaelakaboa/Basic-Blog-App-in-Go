package main

import "sync"

type Post struct {
	ID      int
	Title   string
	Content string
}

var (
	posts   = []Post{}
	nextID  = 1
	postMux sync.Mutex
)
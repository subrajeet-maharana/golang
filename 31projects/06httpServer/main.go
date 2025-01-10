package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type Post struct {
  ID int `json:"id"`
  Body string `json:"body"`
}

var (
  posts = make(map[int]Post)
  nextId = 1
  postsMu sync.Mutex
)

func main() {
  http.HandleFunc("/posts", postsHandler)
  http.HandleFunc("/post", postHandler)

  fmt.Println("Server is running at 3000")
  log.Fatal(http.ListenAndServe(":3000", nil))
}

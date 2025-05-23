package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type Post struct {
	ID   int    `json:"id"`
	Body string `json:"body"`
}

var (
	posts   = make(map[int]Post)
	nextId  = 1
	postsMu sync.Mutex
)

func postsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleGetPosts(w, r)
	case "POST":
		handlePostPosts(w, r)
	default:
		http.Error(w, "Method not allowed!", http.StatusMethodNotAllowed)
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/posts/"):])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "GET":
		handleGetPost(w, r, id)
	case "DELETE":
		handleDeletePost(w, r, id)
	default:
		http.Error(w, "Method not allowed!", http.StatusMethodNotAllowed)
	}
}

func handleGetPosts(w http.ResponseWriter, r *http.Request) {
	postsMu.Lock()
	defer postsMu.Unlock()

	newPosts := make([]Post, 0, len(posts))
	for _, post := range posts {
		newPosts = append(newPosts, post)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newPosts)
}

func handlePostPosts(w http.ResponseWriter, r *http.Request) {
	var newPost Post

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body...", http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(body, &newPost); err != nil {
		http.Error(w, "Error parsing request body...", http.StatusBadRequest)
		return
	}

	postsMu.Lock()
	defer postsMu.Unlock()

	newPost.ID = nextId
	nextId++
	posts[newPost.ID] = newPost

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPost)
}

func handleGetPost(w http.ResponseWriter, r *http.Request, id int) {
	postsMu.Lock()
	defer postsMu.Unlock()

	post, ok := posts[id]
	if !ok {
		http.Error(w, "Post with this ID is not found!", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func handleDeletePost(w http.ResponseWriter, r *http.Request, id int) {
	postsMu.Lock()
	defer postsMu.Unlock()

	_, ok := posts[id]
	if !ok {
		http.Error(w, "Post with this ID is not found!", http.StatusNotFound)
		return
	}

	delete(posts, id)
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/posts", postsHandler)
	http.HandleFunc("/post", postHandler)

	fmt.Println("Server is running at 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

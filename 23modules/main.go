package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to golang modules!")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", r))
}

// go mod tidy/anything is a very expensive operation so always lookout on that
// go mod tidy / graph / why
// to package go mod vendor
// go run -mod=vendor main.go --> tries to run from vendor
// go modules have replaced workspaces
func greeter() {
	fmt.Println("Hey there mod users!")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to modules tutorial in Golang!</h1>"))
}

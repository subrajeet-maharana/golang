package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/subrajeet-maharana/mongoapi/router"
)

func main() {
	fmt.Println("MongoAPI Success!")
	r := router.Router()
	fmt.Println("Server running successfully...")
	log.Fatal(http.ListenAndServe(":7000", r))
	fmt.Println("Listening ...")

}

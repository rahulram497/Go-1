package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("starting server on :5000")

	err := http.ListenAndServe(":5000", mux)
	log.Fatal(err)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("It`s work"))
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

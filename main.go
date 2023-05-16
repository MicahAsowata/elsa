package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Elsa")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	log.Println("Listening on port :3000")
	http.ListenAndServe(":3000", mux)
}

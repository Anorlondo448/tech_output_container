package main

import (
	"fmt"
	"net/http"
)

func topHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "app-top")
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "app-pong")
}

func main() {
	http.HandleFunc("/", topHandler)
	http.HandleFunc("/ping", pingHandler)
	http.ListenAndServe(":8081", nil)
}

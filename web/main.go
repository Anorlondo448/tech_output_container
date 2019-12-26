package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func topHandler(w http.ResponseWriter, r *http.Request) {

	url := "http://app:8081/ping"
	resp, _ := http.Get(url)
	byteArray, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	fmt.Fprintf(w, "web-"+string(byteArray))
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "web-pong")
}

func main() {
	http.HandleFunc("/", topHandler)
	http.HandleFunc("/ping", pingHandler)
	http.ListenAndServe(":8080", nil)
}

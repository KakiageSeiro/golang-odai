package main

import (
	"my-golang-odai/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/form", handler.FormHandler)
	http.HandleFunc("/tweet", handler.TweetHandler)
	http.ListenAndServe(":80", nil)
}

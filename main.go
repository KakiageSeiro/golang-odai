package main

import (
	"my-golang-odai/handler"
	"net/http"
)

func main() {
  http.HandleFunc("/", handler.IndexHandler)
  http.ListenAndServe(":80", nil)
}

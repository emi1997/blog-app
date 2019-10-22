package main

import (
	"github.com/emi1997/blog-app/api"
  "net/http"
  "log"
)

// BasicHome = Create a struct wich holds the information to be shown on the home page
type BasicHome struct{
  Welcome string
  Post string
  Time string
}


func main() {
  http.HandleFunc("/", api.EmptyCall)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/emi1997/blog-app/api"
)

// BasicHome = Create a struct wich holds the information to be shown on the home page
type BasicHome struct {
	Welcome string
	Post    string
	Time    string
}

func main() {
	http.HandleFunc("/", api.EmptyCall)
	http.HandleFunc("/aboutme", api.AboutMe)
	log.Fatal(http.ListenAndServe(":8080", nil))
	info, err := os.IsNotExist(err)
}

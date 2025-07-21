package main

import (
	"log"
	"net/http"

	"github.com/rslbn/blog/internal/app"
)

func main() {
	router := app.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

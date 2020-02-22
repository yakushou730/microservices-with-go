package main

import (
	"net/http"

	"github.com/yakushou730/microservices-with-go/handlers"
)

func main() {
	http.HandleFunc("/example", handlers.MyHandler)
	http.ListenAndServe(":8080", nil)
}

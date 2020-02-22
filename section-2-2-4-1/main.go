package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "Server: ", log.Lshortfile)
	logDecorator := logDecoratorCreator(logger)

	http.Handle("/example", logDecorator(HeaderDecorator(http.HandlerFunc(myHandlerFunc))))
	http.Handle("/example2", logDecorator(HeaderDecorator(http.HandlerFunc(myHandlerFunc2))))

	log.Println("Starting Server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func myHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// Write Status Code
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/text")
	fmt.Fprintln(w, "Example-1; Request was successful")
}

func myHandlerFunc2(w http.ResponseWriter, r *http.Request) {
	// Write Status Code
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/text")
	fmt.Fprintln(w, "Example-2; Request was successful")
}

func HeaderDecorator(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Passed-Through-Decorator", "YES")
		h.ServeHTTP(w, r)
	})
}

type HttpDecorator func(http.Handler) http.Handler

func logDecoratorCreator(logger *log.Logger) HttpDecorator {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Printf(r.RequestURI+": We have a connection from %s", r.RemoteAddr)
			defer logger.Println("Connection Closed")
			h.ServeHTTP(w, r)
		})
	}
}

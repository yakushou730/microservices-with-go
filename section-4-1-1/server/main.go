package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	certPath := "server.pem"
	keyPath := "server.key"

	http.HandleFunc("/hello", myHandlerFunc)

	if err := http.ListenAndServeTLS(":8443", certPath, keyPath, nil); err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}

func myHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// Write Status Code
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello world from HTTPS.")
}

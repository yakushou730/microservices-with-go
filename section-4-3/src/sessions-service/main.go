package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yakushou730/microservices-with-go/section-4-3/src/sessions-service/handlers"
	"github.com/yakushou730/microservices-with-go/section-4-3/src/sessions-service/repositories"
)

func main() {
	handler := &handlers.Handlers{
		Repo: repositories.NewRedisSessionsRepository(),
	}

	r := mux.NewRouter()
	r.HandleFunc("/session/{key}", handler.GetSession).Methods("GET")
	r.HandleFunc("/session/{key}", handler.SetSession).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8001", r))
}

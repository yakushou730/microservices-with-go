package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yakushou730/microservices-with-go/api/handlers"
	"github.com/yakushou730/microservices-with-go/api/repository"
)

func main() {
	repo := repository.NewRepository(
		"mongodb://localhost:27017",
		"packt",
		"timeZones",
	)
	defer repo.Close()

	h := handlers.Handlers{
		Repo: repo,
	}

	r := mux.NewRouter()
	r.HandleFunc("/timeZones", h.All).Methods("GET")
	r.HandleFunc("/timeZones/{timeZone}", h.GetByTZ).Methods("GET")

	r.HandleFunc("/timeZones", h.Insert).Methods("POST")
	r.HandleFunc("/timeZones/{timeZone}", h.Delete).Methods("DELETE")
	r.HandleFunc("/timeZones/{timeZone}", h.Update).Methods("PATCH")

	log.Fatal(http.ListenAndServe(":8080", r))
}

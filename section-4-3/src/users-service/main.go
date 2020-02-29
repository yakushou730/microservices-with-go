package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yakushou730/microservices-with-go/section-4-3/src/users-service/handlers"
	"github.com/yakushou730/microservices-with-go/section-4-3/src/users-service/repositories"
)

func main() {
	handler := &handlers.Handlers{
		Repo: repositories.NewMySQLUserRepository(),
	}
	defer handler.Repo.Close()

	r := mux.NewRouter()
	r.HandleFunc("/user/{username}", handler.GetUserByUsernameHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}

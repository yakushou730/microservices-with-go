package main

import (
	"log"
	"net/http"

	"github.com/yakushou730/microservices-with-go/section-5-1/src/users-service/repositories"

	"github.com/yakushou730/microservices-with-go/section-5-1/src/users-service/handlers"

	"github.com/gorilla/mux"
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

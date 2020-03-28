package main

import (
	"log"
	"net/http"

	"github.com/yakushou730/microservices-with-go/section-6-2/src/users-service/usecases"

	"github.com/yakushou730/microservices-with-go/section-6-2/src/users-service/repositories"

	"github.com/yakushou730/microservices-with-go/section-6-2/src/users-service/handlers"

	"github.com/gorilla/mux"
)

func main() {
	cacheRepo := repositories.NewRedisUsersRepository()
	repo := repositories.NewMySQLUserRepository()
	defer repo.Close()

	handler := &handlers.Handlers{
		GetUserUseCase: &usecases.GetUserUsecase{
			CacheRepo: cacheRepo,
			Repo:      repo,
		},
		UpdateUseCase: &usecases.UpdateUserUsecase{
			CacheRepo: cacheRepo,
			Repo:      repo,
		},
	}

	r := mux.NewRouter()

	r.HandleFunc("/user/{username}", handler.GetUserByUsernameHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}

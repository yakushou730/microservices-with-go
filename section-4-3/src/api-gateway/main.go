package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yakushou730/microservices-with-go/section-4-3/src/api-gateway/handlers"
	"github.com/yakushou730/microservices-with-go/section-4-3/src/api-gateway/repositories"
)

func main() {
	r := mux.NewRouter()

	h := &handlers.Handler{
		SessionsRepo: &repositories.RestSessionsRepository{},
		UsersRepo:    &repositories.RestUsersRepository{},
	}

	r.HandleFunc("/authorize", h.Authorize).Methods("POST")

	r.HandleFunc("/restricted/resource-1", handlers.VerifyJWT(h.AddSessionData(h.Restricted1)))
	r.HandleFunc("/restricted/resource-2", handlers.VerifyJWT(h.AddSessionData(h.Restricted2)))

	// certPath := "server.pem"
	// keyPath := "server.key"
	//
	// err := http.ListenAndServeTLS(":8443", certPath, keyPath, r)
	err := http.ListenAndServe(":8443", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

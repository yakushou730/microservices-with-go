package main

import (
	"log"
	"net/http"
	"net/http/pprof"

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
	// BASIC PPROF HANDLERS
	r.HandleFunc("/debug/pprof", pprof.Index)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	// Manually add support for paths linked to by index page at /debug/pprof/
	r.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	r.Handle("/debug/pprof/mutex", pprof.Handler("mutex"))
	r.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	r.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	r.Handle("/debug/pprof/block", pprof.Handler("block"))

	r.HandleFunc("/user/{username}", handler.GetUserByUsernameHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}

package main

import (
	"github.com/elvismdnin/matchGateway/web"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	port := "8000"
    router := mux.NewRouter()

	web.CreateGame().AddRoute(router)
	web.ServeSPA(router)

	srv := &http.Server{
		Handler: router,
		Addr:    "localhost:" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

    log.Println("Server opened at " + port)
    log.Fatal(srv.ListenAndServe())
}
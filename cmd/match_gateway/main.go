package main

import (
	intweb "github.com/elvismdnin/match_gateway/internal/web"
	"github.com/elvismdnin/match_gateway/web"
	"github.com/gorilla/mux"

	"log"
	"net/http"
	"time"
)

func main() {
	port := "8000"
    router := mux.NewRouter()

	intweb.InitSession()
	web.CreateGame().AddRoute(router)
	web.ServeSPA(router)

	srv := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

    log.Println("Server opened at " + port)
    log.Fatal(srv.ListenAndServe())
}
package web

import (
	"github.com/elvismdnin/matchGateway/internal/web"
	"github.com/gorilla/mux"
)

func ServeSPA(r *mux.Router)  {
	spa := web.SpaHandler{StaticPath: "static", IndexPath: "index.html"}
	r.PathPrefix("/").Handler(spa)
}
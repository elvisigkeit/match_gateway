package web

import (
	"fmt"
	"github.com/elvismdnin/matchGateway/internal/web"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateGame() web.Handler {
	return web.Handler {
		Route: func(r *mux.Route) {
			r.Path("/new").Methods("GET")
		},
		Func: func(w http.ResponseWriter, r *http.Request) {
			var encodeErr error

			_, _ = fmt.Fprint(w, "{success}")

			if encodeErr != nil {
				http.Error(w, encodeErr.Error(), http.StatusInternalServerError)
			}
		},
	}
}
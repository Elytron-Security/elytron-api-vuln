package router

import (
	"api/src/router/route"

	"github.com/gorilla/mux"
)

func Generate() *mux.Router {
	r := mux.NewRouter()

	return route.RouteConfig(r)
}
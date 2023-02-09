package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Routes struct {
	URI    string
	Metodo string
	Funcao func(http.ResponseWriter, *http.Request)
}

// config das rotas padrão
func Config(router *mux.Router) *mux.Router {

	routes := routesIndex
	for _, route := range routes {
		router.HandleFunc(route.URI, route.Funcao).Methods(route.Metodo)
	}
	// fileServer := http.FileServer(http.Dir("./assets/")) Se for adicionar css ou js depois
	// router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}

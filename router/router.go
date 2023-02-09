package router

import (
	"modulo/router/routes"

	"github.com/gorilla/mux"
)

// retorna um router com todas as rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return routes.Config(r)
}

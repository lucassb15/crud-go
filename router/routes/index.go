package routes

import (
	"modulo/controllers"
	"net/http"
)

var routesIndex = []Routes{
	// Metodo Get para carregar o index
	{
		URI:    "/",
		Metodo: http.MethodGet,
		Funcao: controllers.LoadIndex,
	},
	// Metodo Post para pegar os inputs do form
	{
		URI:    "/indexPost",
		Metodo: http.MethodPost,
		Funcao: controllers.LoadIndexPost,
	},
	{
		URI:    "/BirthDayCalc",
		Metodo: http.MethodPost,
		Funcao: controllers.BirthdayCalc,
	},
}

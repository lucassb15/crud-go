package controllers

import (
	"fmt"
	"modulo/models"
	"modulo/utils"
	"net/http"
)

type PegaData struct {
	DataD *models.DateAge
}

// Carrega a pagina localhost:8000
func LoadIndex(w http.ResponseWriter, r *http.Request) {
	utils.ReadTemplate(w, "index.html", nil)

}

// Pega as informações do Formulario e armazena no "UserData"
func LoadIndexPost(w http.ResponseWriter, r *http.Request) {

	UserData := models.User{
		Nome:  r.FormValue("nome"),
		Email: r.FormValue("email"),
	}
	dateBirth := models.DateAge{
		Date: r.FormValue("birthday"),
	}
	utils.ReadTemplate(w, "index.html", UserData)
	fmt.Println(UserData, dateBirth)

}
func BirthdayCalc(w http.ResponseWriter, r *http.Request) {
	utils.ReadTemplate(w, "BirthDayCalc.html", nil)
}
func (me *PegaData) Teste(numero int32) int32 {
	cavalo := 10
	fmt.Println(cavalo)
	return numero
}

package controllers

import (
	"fmt"
	"modulo/models"
	"modulo/utils"
	"net/http"
	"strconv"
	"time"
)

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

	utils.ReadTemplate(w, "index.html", UserData)
	fmt.Println(UserData)

}

func BirthdayCalc(w http.ResponseWriter, r *http.Request) {
	dateBirth := models.DateAge{
		Date: r.FormValue("birthday"),
	}

	// Converte o input do usuário para Int pegando apenas o ano
	fmt.Println(dateBirth)
	yearInput, err := strconv.ParseInt(r.FormValue("birthday")[0:4], 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(yearInput)

	// Pega a data atual e retira apenas o Ano
	CurrentTime := time.Now()
	CurrentTime.Year()
	YearNow := CurrentTime.Year()
	fmt.Println(YearNow)

	// Calcular a idade
	CalcDate := YearNow - int(yearInput)
	fmt.Println(CalcDate)

	// carregar template
	utils.ReadTemplate(w, "BirthDayCalc.html", dateBirth)

}

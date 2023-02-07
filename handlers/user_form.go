package handlers

import (
	"fmt"
	"modulo/controllers"
	"net/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "erro ao fazer parse do form: %v", err)
			return
		}

		err := controllers.Create(r.PostForm.Get("name"), r.PostForm.Get("email"))
		if err != nil {
			fmt.Fprintf(w, "erro ao salvar os dados: %v", err)
			return
		}
	}

	http.ServeFile(w, r, "handlers/templates/user_form.html")
}

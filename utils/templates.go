package utils

import (
	"html/template"
	"net/http"
)

var tmpl *template.Template

func CarregarTemplates() {
	tmpl = template.Must(template.ParseGlob("views/*.html")) // passa uma forma de identificar os arquivos que vai ser considerado templates fica na pasta "views"
}

func ReadTemplate(w http.ResponseWriter, template string, dados interface{}) {
	tmpl.ExecuteTemplate(w, template, dados) // os controllers utilizam essa função para executar os templates
}

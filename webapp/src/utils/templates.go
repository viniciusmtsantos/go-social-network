package utils

import (
	"net/http"
	"text/template"
)

var templates *template.Template

// CarregarTemplates insere os templates html na variavel templates
func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
	templates = template.Must(templates.ParseGlob("views/templates/*.html"))
}

// ExecutarTemplate renderiza uma pagina na tela
func ExecutarTemplate(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, template, dados)
}

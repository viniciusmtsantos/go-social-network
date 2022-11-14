package controllers

import (
	"net/http"
	"webapp/src/cookies"
)

// Fazer Logout remove os dados de autenticação salvos no browser
func FazerLogout(w http.ResponseWriter, r *http.Request) {
	cookies.Deletar(w)
	http.Redirect(w, r, "/login", 302)
}

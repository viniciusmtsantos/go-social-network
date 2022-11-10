package middlewares // pacote entre a requisião e a resposta

import (
	"api/src/authentication"
	"api/src/responses"
	"log"
	"net/http"
)

// Logger escreve informaçoes da requisição o terminal
func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n%s %s %s", r.Method, r.RequestURI, r.Host)
		proximaFuncao(w, r)
	}
}

// Autenticar verifica se o usuario fazendo a requisição está autenticado
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := authentication.ValidarToken(r); erro != nil {
			responses.Erro(w, http.StatusUnauthorized, erro)
			return
		}
		proximaFuncao(w, r)
	}
}

package rotas // criando um struct para todas as rotas da API

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa todas as rotas da API, todas tem a mesma estrutura
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Retorna o router com todas as rotas configuradas
// Configurar retorna todas as rotas dentro do router
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios // nosso router entrou sem nenhuma função
	rotas = append(rotas, rotaLogin)
	rotas = append(rotas, rotasPublicacoes...)

	for _, rota := range rotas {

		if rota.RequerAutenticacao {
			r.HandleFunc(rota.URI,
				middlewares.Logger(middlewares.Autenticar(rota.Funcao)),
			).Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.URI, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}
	}
	return r // e sai configurada
}

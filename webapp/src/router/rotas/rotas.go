package rotas // criando um struct para todas as rotas da API

import (
	"net/http"
	"webapp/src/middlewares"

	"github.com/gorilla/mux"
)

// Rota representa todas as rotas da aplicação web, todas tem a mesma estrutura
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

func Configurar(router *mux.Router) *mux.Router {
	rotas := rotasLogin
	rotas = append(rotas, rotasUsuarios...)
	rotas = append(rotas, rotaPaginaPrincipal) // uma unica rota
	rotas = append(rotas, rotasPublicacoes...) // uma unica rota

	for _, rota := range rotas {

		if rota.RequerAutenticacao {
			router.HandleFunc(rota.URI,
				middlewares.Logger(middlewares.Autenticar(rota.Funcao)),
			).Methods(rota.Metodo)

		} else {
			router.HandleFunc(rota.URI,
				middlewares.Logger(rota.Funcao),
			).Methods(rota.Metodo)
		}
	}

	filServer := http.FileServer(http.Dir("./assets"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", filServer))

	return router // e sai configurada
}

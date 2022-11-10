package controllers // todos que estão no pacote controllers vao lidar com as funções http, isto é, receber as requisições e devolver as responses.

import (
	"api/src/authentication"
	"api/src/data"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/safety"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CriarUsuario insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario models.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar("cadastro"); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := data.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	usuario.ID, erro = repositorio.Criar(usuario)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusCreated, usuario)
}

// BuscarUsuarios busca todos os usuários salvos no banco
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))
	db, erro := data.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	usuarios, erro := repositorio.Buscar(nomeOuNick)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, usuarios)
}

// BuscarUsuario busca um usuário salvo no banco
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := data.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	usuario, erro := repositorio.BuscarPorID(usuarioID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, usuario)
}

// AtualizarUsuario altera as informações de um usuário no banco
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	usuarioIDNoToken, erro := authentication.ExtrairUsuarioID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if usuarioID != usuarioIDNoToken {
		responses.Erro(w, http.StatusForbidden, errors.New("Não é possível atualizar um usuário que não seja o seu"))
		return
	}

	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario models.Usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar("edicao"); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := data.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	if erro = repositorio.Atualizar(usuarioID, usuario); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// DeletarUsuario exclui as informações de um usuário no banco
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	usuarioIDNoToken, erro := authentication.ExtrairUsuarioID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if usuarioID != usuarioIDNoToken {
		responses.Erro(w, http.StatusForbidden, errors.New("Não é possível deletar um usuário que não seja o seu"))
		return
	}

	db, erro := data.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	if erro = repositorio.Deletar(usuarioID); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// SeguirUsuario permite que um usuário siga outro
func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	seguidorID, erro := authentication.ExtrairUsuarioID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if seguidorID == usuarioID {
		responses.Erro(w, http.StatusForbidden, errors.New("Não é possível seguir você mesmo"))
		return
	}

	db, erro := data.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	if erro = repositorio.Seguir(usuarioID, seguidorID); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// PararDeSeguirUsuario permite que um usuário pare de seguir outro
func PararDeSeguirUsuario(w http.ResponseWriter, r *http.Request) {
	seguidorID, erro := authentication.ExtrairUsuarioID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if seguidorID == usuarioID {
		responses.Erro(w, http.StatusForbidden, errors.New("Não é possível parar de seguir você mesmo"))
		return
	}

	db, erro := data.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	if erro = repositorio.PararDeSeguir(usuarioID, seguidorID); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// BuscarSeguidores traz todos os seguidores de um usuário
func BuscarSeguidores(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := data.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	seguidores, erro := repositorio.BuscarSeguidores(usuarioID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, seguidores)
}

// BuscarSeguindo traz todos os usuários que um determinado usuário está seguindo
func BuscarSeguindo(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := data.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	usuarios, erro := repositorio.BuscarSeguindo(usuarioID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, usuarios)
}

// AtualizarSenha permite alterar a senha de um usuário
func AtualizarSenha(w http.ResponseWriter, r *http.Request) {
	usuarioIDNoToken, erro := authentication.ExtrairUsuarioID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if usuarioIDNoToken != usuarioID {
		responses.Erro(w, http.StatusForbidden, errors.New("Não é possível atualizar a senha de um usuário que não seja o seu"))
		return
	}

	corpoRequisicao, erro := io.ReadAll(r.Body)
	var senha models.Senha
	if erro = json.Unmarshal(corpoRequisicao, &senha); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := data.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	senhaSalvaNoBanco, erro := repositorio.BuscarSenha(usuarioID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = safety.VerificarSenha(senhaSalvaNoBanco, senha.Atual); erro != nil {
		responses.Erro(w, http.StatusUnauthorized, errors.New("A senha atual não condiz com a que está salva no banco"))
		return
	}

	senhaComHash, erro := safety.Hash(senha.Nova)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = repositorio.AtualizarSenha(usuarioID, string(senhaComHash)); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

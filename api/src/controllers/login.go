package controllers

import (
	"api/src/authentication"
	"api/src/data"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/safety"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

// Login é responsável por autenticar um usuário na API
func Login(w http.ResponseWriter, r *http.Request) {
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

	db, erro := data.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	usuarioSalvoNodata, erro := repositorio.BuscarPorEmail(usuario.Email)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = safety.VerificarSenha(usuarioSalvoNodata.Senha, usuario.Senha); erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	token, erro := authentication.CriarToken(usuarioSalvoNodata.ID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	usuarioID := strconv.FormatUint(usuarioSalvoNodata.ID, 10)

	responses.JSON(w, http.StatusOK, models.DadosAutenticacao{ID: usuarioID, Token: token})
}

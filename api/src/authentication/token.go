package authentication

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Criar Token retorna um token assinado com as permissoes do usuario
func CriarToken(usuarioID uint64) (string, error) { // a string é o token a ser retornado
	permissoes := jwt.MapClaims{} // map com as permissoes dentro do token
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioId"] = usuarioID
	// secret para fazzer a assinatuara deste token e garantir a autenticidade deste token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.SignedString([]byte(config.SecretKey)) // chave para ser utilizada e fazer a assinatura
}

// ValidarToken verifica se o token passado na requisição é valido
func ValidarToken(r *http.Request) error {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerficacao)
	if erro != nil {
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Token invalido")
}

// ExtraiUsuarioID retorna o usuarioId que está salvo no token
func ExtrairUsuarioID(r *http.Request) (uint64, error) {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerficacao)
	if erro != nil {
		return 0, erro
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		usuarioID, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permissoes["usuarioId"]), 10, 64)
		if erro != nil {
			return 0, erro
		}

		return usuarioID, nil
	}

	return 0, errors.New("Token invalido")
}

func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	// Bearer scdshcisduchidsuchsdicsd
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return " "
}

func retornarChaveDeVerficacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Metodo de assinatura inesperado! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}

package db

import (
	"github.com/Ignis-Divine/api-nodemcu/models"
	"golang.org/x/crypto/bcrypt"
)

//IntentoLogin revisa el usuario en la db
func IntentoLogin(email string, password string) (models.Usuario, bool) {
	usu, encontrado, _ := RevisoSiExisteUsuario(email)
	if encontrado == false {
		return usu, false
	}
	passwordBytes := []byte(password)
	passwordDB := []byte(usu.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return usu, false
	}
	return usu, true
}

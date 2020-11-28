package db

import (
	"golang.org/x/crypto/bcrypt"
)

//EncriptarPassword  es la rutina que permite encriptar el password del usuario
func EncriptarPassword(pass string) (string, error) {
	//el costo es el algoritmo es 2^costo, mayor costo, mejor encriptacion, si es mayor el costo mas tiempo tarda
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}

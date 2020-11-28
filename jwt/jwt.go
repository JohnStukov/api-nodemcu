package jwt

import (
	"github.com/Ignis-Divine/api-nodemcu/models"
	"time"
	//alias y paquete
	jwt "github.com/dgrijalva/jwt-go"
)

//GeneroJWT genera el encriptado con JWT
func GeneroJWT(t models.Usuario) (string, error) {
	miClave := []byte("qwerty_12345")
	//el password jamas se pone en el json
	payload := jwt.MapClaims{
		"_id":            t.ID.Hex(),
		"email":          t.Email,
		"nombre":         t.Nombre,
		"apellidos":      t.Apellidos,
		"fecha_registro": t.FechaRegistro,
		"imagen":         t.Imagen,
		"exp":            time.Now().Add(time.Hour * 24).Unix(),
	}
	//pasamos el header y el payload
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}

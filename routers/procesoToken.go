package routers

import (
	"errors"
	"github.com/Ignis-Divine/api-nodemcu/db"
	"github.com/Ignis-Divine/api-nodemcu/models"
	jwt "github.com/dgrijalva/jwt-go"
)

//variables que vamos a exportar en todas las rutas
//Email valor Email para ser usado en todos los endpoints
var Email string

//IDUsuario es el ID devuelto del modelo, que se usara en todos los endpoints
var IDUsuario string
var IDDatos string
var IDAlarma string

//ProcesoToken realiza el proceso para extraer los valores, el error siempre va al final de los parametros de salida
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("qwerty_12345")
	claims := &models.Claim{}

	/*quitamos la palabra standar Bearer
	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token no valido")
	}
	tk = strings.TrimSpace(splitToken[1])
	*/ //quitamos la palabra standar Bearer

	//esto hace la validacion misma del token
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, _ := db.RevisarSiExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}
	return claims, false, string(""), err
}

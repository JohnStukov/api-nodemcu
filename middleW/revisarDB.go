package middleW

import (
	"github.com/Ignis-Divine/api-nodemcu/db"
	"net/http"
)

//RevisarDB ES EL middleware q permite conocer el estado de la base de datos
func RevisarDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.RevisarConexion() == 0 {
			http.Error(w, "Conexión perdida con la Base de Datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}


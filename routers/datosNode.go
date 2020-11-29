package routers

import (
	"encoding/json"
	"github.com/Ignis-Divine/api-nodemcu/db"
	"github.com/Ignis-Divine/api-nodemcu/models"
	"log"
	"net/http"
	"strings"
)

func CrearRegistro(w http.ResponseWriter, r *http.Request) {
	x := r.Header.Get("Authorization")
	auth := strings.Replace(x, "Basic ", "", -1)
	defer r.Body.Close()
	var t models.Datos
	t.MacNodemcu=auth
	_, encontrado, _ := db.RevisarSiExisteNodemcu(auth)
	if encontrado == true {
		http.Error(w, "No existe este dispositivo", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}
	log.Println("json:", t)
	_, status, err := db.InsertoDatos(t)
	if err != nil {
		http.Error(w, "Error al intentar el registro de datos "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se logro el registro de los datos ", 400)
		return
	}
	//se registran los datos del nodemcu
	w.WriteHeader(http.StatusCreated)
}
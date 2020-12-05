package routers

import (
	"encoding/json"
	"github.com/Ignis-Divine/api-nodemcu/db"
	"github.com/Ignis-Divine/api-nodemcu/models"
	"net/http"
	"strconv"
	"strings"
)

func CrearRegistroDatos(w http.ResponseWriter, r *http.Request) {
	x := r.Header.Get("Authorization")
	auth := strings.Replace(x, "Basic ", "", -1)
	defer r.Body.Close()
	var t models.Datos
	t.MacNodemcu=auth
	_, encontrado, _ := db.RevisoSiExisteNodemcu(auth)
	if encontrado != true {
		http.Error(w, "No existe este dispositivo", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}
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

func ListarDatos(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	fecha := r.URL.Query().Get("fecha")
	hora := r.URL.Query().Get("hora")

	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "debe enviar el parametro pagina como entero mayor a cero", http.StatusBadRequest)
		return
	}
	pag := int64(pagTemp)

	result, status := db.ObtenerRegistros(IDUsuario, pag, fecha, hora)
	if status != false {
		http.Error(w, "error al leer datos", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}


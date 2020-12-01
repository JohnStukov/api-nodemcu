package routers

import (
	"encoding/json"
	"github.com/Ignis-Divine/api-nodemcu/db"
	"github.com/Ignis-Divine/api-nodemcu/models"
	"net/http"
	"strconv"
	"strings"
)

func CrearRegistroAlarma(w http.ResponseWriter, r *http.Request) {
	x := r.Header.Get("Authorization")
	auth := strings.Replace(x, "Basic ", "", -1)
	defer r.Body.Close()
	var a models.Alarma
	a.MacNodemcu=auth
	n, encontrado, _ := db.RevisarSiExisteNodemcu(auth)
	if encontrado != true {
		http.Error(w, "No existe este dispositivo", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}
	_, status, err := db.InsertoAlarma(a)
	if err != nil {
		http.Error(w, "Error al intentar el registro de datos "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se logro el registro de los datos ", 400)
		return
	}
	//se registran los datos del nodemcu y se enviá un mensaje de texto
	Txt(n.D, a.Tipo)
	w.WriteHeader(http.StatusCreated)
}

func ListarAlarmas(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	fecha := r.URL.Query().Get("fecha")
	hora := r.URL.Query().Get("hora")
	tipo := r.URL.Query().Get("tipo")

	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "debe enviar el parámetro pagina como entero mayor a cero", http.StatusBadRequest)
		return
	}
	pag := int64(pagTemp)

	result, status := db.ObtenerAlarmas(IDAlarma, pag, fecha, hora, tipo)
	if status != false {
		http.Error(w, "error al leer datos", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
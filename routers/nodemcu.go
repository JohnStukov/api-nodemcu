package routers

import (
	"encoding/json"
	"github.com/Ignis-Divine/api-nodemcu/db"
	"github.com/Ignis-Divine/api-nodemcu/models"
	"net/http"
	"strconv"
)

//Registro crea el registro del nodemcu
func RegistrarNodemcu(w http.ResponseWriter, r *http.Request) {
	var t models.Nodemcu
	//el Body del http.response es un Stream, solo se lee una ves y se destruye
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}
	if len(t.Dispositivo) == 0 {
		http.Error(w, "El nombre del dispoditivo es requerido ", 400)
		return
	}
	if len(t.Token) == 0 {
		http.Error(w, "La MAC del dispositivo es requerido ", 400)
		return
	}
	if len(t.Status) == 0 {
		http.Error(w, "El status del dispositivo es requerido ", 400)
		return
	}
	//revisamos que haya correos repetidos
	_, encontrado, _ := db.RevisoSiExisteNodemcu(t.Token)
	if encontrado == true {
		http.Error(w, "Ya existe un nodemcu con esta MAC registrado ", 400)
		return
	}
	//revisar si fue correcto el registro
	_, status, err := db.InsertoRegistroNodemcu(t)
	if err != nil {
		http.Error(w, "Error al intentar el registro de usuario "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se logro el registro del usuario ", 400)
		return
	}
	//se registra el nodemcu
	w.WriteHeader(http.StatusCreated)
}


func ModificarNodemcu(w http.ResponseWriter, r *http.Request) {
	var t models.Nodemcu
	//grabamos el body
	err := json.NewDecoder(r.Body).Decode(&t)
	//revisamos q el body no esta vacio o que no contiene errores de diseño el documento
	if err != nil {
		http.Error(w, "datos incorrectos "+err.Error(), 400)
		return
	}
	var status bool
	status, err = db.ModificoNodemcu(t, IDUsuario)
	if err != nil {
		http.Error(w, "ocurrio un error al modificar el registro, intente de nuevo "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "no se logro modificar el registro del usuario "+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}


func EliminarNodemcu(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}
	_, err := db.BuscoNodemcu(ID)
	if err != nil {
		http.Error(w, "dato no encontrado "+err.Error(), 400)
		return
	}
	err = db.BorroNodemcu(ID)
	if err != nil {
		http.Error(w, "ocurrio un error al eliminar el usuario "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

//VerUsuario permite extraer los valores del perfil
func VerNodemcu(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "debe enviar el parametro ID", http.StatusBadRequest)
		return
	}
	nodemcu, err := db.BuscoNodemcu(ID)
	if err != nil {
		http.Error(w, "error al encontrar el registro"+err.Error(), 400)
		return
	}
	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nodemcu)
}

func ListarNodemcus(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	busqueda := r.URL.Query().Get("busqueda")
	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "debe enviar el parametro pagina como entero mayor a cero", http.StatusBadRequest)
		return
	}
	pag := int64(pagTemp)
	result, status := db.ListoNodemcus(IDUsuario, pag, busqueda)
	if status == false {
		http.Error(w, "error al leer usuarios", http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
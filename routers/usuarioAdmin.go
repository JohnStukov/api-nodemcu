package routers

import (
	"encoding/json"
	"github.com/Ignis-Divine/api-nodemcu/db"
	"github.com/Ignis-Divine/api-nodemcu/models"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

//Registro crea el registro del usuario
func RegistrarUsuarioAdmin(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	//el Body del http.response es un Stream, solo se lee una ves y se destruye
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}
	//revisa si el correoesta vacio
	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido ", 400)
		return
	}
	//validamos la longitud de la contraseña
	if len(t.Password) < 8 {
		http.Error(w, "La contraseña debe ser minimo 6 caracteres ", 400)
		return
	}
	//revisamos que haya correos repetidos
	_, encontrado, _ := db.RevisoSiExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un usuario con este email ", 400)
		return
	}
	//revisar si fue correcto el registro
	_, status, err := db.InsertoRegistroUsuario(t)
	if err != nil {
		http.Error(w, "Error al intentar el registro de usuario "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se logro el registro del usuario ", 400)
		return
	}
	//se registra el usuario
	w.WriteHeader(http.StatusCreated)
}


func ModificarUsuario(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	//grabamos el body
	err := json.NewDecoder(r.Body).Decode(&t)
	//revisamos q el body no esta vacio o que no contiene errores de diseño el documento
	if err != nil {
		http.Error(w, "datos incorrectos "+err.Error(), 400)
		return
	}
	var status bool
	status, err = db.ModificoRegistro(t, IDUsuario)
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


func EliminarUsuario(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}
	_, err := db.BuscoUsuario(ID)
	if err != nil {
		http.Error(w, "dato no encontrado "+err.Error(), 400)
		return
	}
	err = db.BorroUsuario(ID)
	if err != nil {
		http.Error(w, "ocurrio un error al eliminar el usuario "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

//VerUsuario permite extraer los valores del perfil
func VerUsuario(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	usuario, err := db.BuscoUsuario(ID)
	if err != nil {
		http.Error(w, "error al encontrar el registro"+err.Error(), 400)
		return
	}
	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(usuario)
}


func ListarUsuarios(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	busqueda := r.URL.Query().Get("busqueda")

	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "debe enviar el parametro pagina como entero mayor a cero", http.StatusBadRequest)
		return
	}
	pag := int64(pagTemp)

	result, status := db.ListoUsuarios(IDUsuario, pag, busqueda)
	if status != false {
		http.Error(w, "error al leer usuarios", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

func SubirImagenUsuario(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("imagen")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo = "uploads/usuarios/" + IDUsuario + "." + extension
	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "error al subir la imagen! "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "error al copiar la imagen! "+err.Error(), http.StatusBadRequest)
		return
	}

	var usuario models.Usuario
	var status bool

	usuario.Imagen = IDUsuario + "." + extension
	status, err = db.ModificoRegistro(usuario, IDUsuario)
	if err != nil || status == false {
		http.Error(w, "error al grabar la imagen en la db! "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}


func ObtenerImagen(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "debe enviar el parametro id ", http.StatusBadRequest)
		return
	}

	perfil, err := db.BuscoUsuario(ID)
	if err != nil {
		http.Error(w, "usuario no encontrado ", http.StatusBadRequest)
		return
	}

	openFile, err := os.Open("uploads/usuarios/" + perfil.Imagen)
	if err != nil {
		http.Error(w, "imagen no encontrada ", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, openFile)
	if err != nil {
		http.Error(w, "error al copiar la imagen ", http.StatusBadRequest)

	}
}
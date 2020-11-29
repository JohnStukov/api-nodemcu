package handlers

import (
	"github.com/Ignis-Divine/api-nodemcu/middleW"
	"github.com/Ignis-Divine/api-nodemcu/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

//Manejadores, configuro el puerto, el handler y se pone en escucha el servidor
func Manejadores() {
	router := mux.NewRouter()
	//-----------------------------RUTAS------------------------------------//
	//-----------------------------LOGIN------------------------------------//
	router.HandleFunc("/login", middleW.RevisarDB(routers.Login)).Methods("POST")
	//-----------------------------USUARIOS---------------------------------//
	router.HandleFunc("/registro", middleW.RevisarDB(routers.RegistroUsuarioAdmin)).Methods("POST")
	router.HandleFunc("/verUsuario", middleW.RevisarDB(middleW.ValidoJWT(routers.VerUsuario))).Methods("GET")
	router.HandleFunc("/modificarUsuario", middleW.RevisarDB(middleW.ValidoJWT(routers.ModificarUsuario))).Methods("PUT")
	router.HandleFunc("/eliminarUsuario", middleW.RevisarDB(middleW.ValidoJWT(routers.EliminarUsuario))).Methods("DELETE")
	router.HandleFunc("/usuarios", middleW.RevisarDB(middleW.ValidoJWT(routers.ListarUsuarios))).Methods("GET")
	//-----------------------------NODEMCUS---------------------------------//
	router.HandleFunc("/nodemcu", middleW.RevisarDB(routers.CrearRegistro)).Methods("POST")
	//----------------------------------------------------------------------//

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "5050"
	}
	//da permisos a todo mundo
	handler := cors.AllowAll().Handler(router)
	//pone a escuchar el puerto elegido y pasa el control al handler
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}

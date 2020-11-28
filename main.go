package main

import (
	//go no entiende de carpetas locales, se tienen q colocar el github
	"github.com/Ignis-Divine/api-nodemcu/db"
	"github.com/Ignis-Divine/api-nodemcu/handlers"
	"log"
)

func main() {
	if db.RevisarConexion() == 0 {
		log.Fatal("Sin conexion a la db")
		return
	}
	handlers.Manejadores()
}

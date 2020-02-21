package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"TallerVI/controllers"
	"TallerVI/driver"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var db *sql.DB

func init() {
	gotenv.Load()
}

func main() {

	db = driver.ConnectDB()
	controller := controllers.Controller{}
	ruta := mux.NewRouter()

	ruta.HandleFunc("/turnos", controller.Turnos(db)).Methods("Get")
	ruta.HandleFunc("/turno/{id}", controller.Detalle(db)).Methods("Get")
	ruta.HandleFunc("/turnos", controller.NuevoTurno(db)).Methods("Post")
	ruta.HandleFunc("/turnos", controller.ModTurno(db)).Methods("Put")
	ruta.HandleFunc("/turno/{id}", controller.DelTurno(db)).Methods("Delete")

	fmt.Println("usar el puerto 8000")
	log.Fatal(http.ListenAndServe(":8000", ruta))
}

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"TallerVI/controllers"
	"TallerVI/driver"
	"TallerVI/models"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var agenda []models.Turno
var db *sql.DB

/*--------------------------------------------------------------------------*/
//funciones

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
	ruta.HandleFunc("/turno/{id}", borrarturno).Methods("Delete")

	fmt.Println("usar el puerto 8000")
	log.Fatal(http.ListenAndServe(":8000", ruta))
}

// IDEA: Endpoints

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

/*
func modturnos(w http.ResponseWriter, r *http.Request) {
	var a models.Turno
	json.NewDecoder(r.Body).Decode(&a)
	resultado, err := db.Exec("update turnos set fecha=$1, hora=$2, id_usuario=$3,id_cliente=$4,id_mascota=$5 where id_turno=$6 RETURNING id_turno", &a.Fecha, &a.Hora, &a.Veterinario, &a.Dueno, &a.Mascota, &a.ID)

	rowsUpdate, err := resultado.RowsAffected()
	logFatal(err)
	json.NewEncoder(w).Encode(rowsUpdate)
}
*/
func borrarturno(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)
	resultado, err := db.Exec("delete from turnos where id_turno=$1", parametro["id"])
	logFatal(err)
	rowsDelete, err := resultado.RowsAffected()
	logFatal(err)
	json.NewEncoder(w).Encode(rowsDelete)
}

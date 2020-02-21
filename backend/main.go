package main

import (
	"TallerVI/controllers"
	"TallerVI/driver"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

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
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	ruta := mux.NewRouter()

	ruta.HandleFunc("/", index).Methods("GET")
	ruta.HandleFunc("/turnos", controller.Turnos(db)).Methods("Get")
	ruta.HandleFunc("/turno/{id}", controller.Detalle(db)).Methods("Get")
	ruta.HandleFunc("/turnos", controller.NuevoTurno(db)).Methods("Post")
	ruta.HandleFunc("/turnos", controller.ModTurno(db)).Methods("Put")
	ruta.HandleFunc("/turno/{id}", controller.DelTurno(db)).Methods("Delete")

	fmt.Println("usar el puerto 8000")
	log.Fatal(http.ListenAndServe(":"+port, ruta))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "bienvenidos")
}
func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

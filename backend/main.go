package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

type Turno struct {
	ID          int    `json:"id_turno"`
	Fecha       string `json:"fecha"`
	Hora        string `json:"hora"`
	Veterinario int    `json:"id_usuario"`
	Dueno       int    `json:"id_cliente"`
	Mascota     int    `json:"id_mascota"`
}

var agenda []Turno
var db *sql.DB

/*--------------------------------------------------------------------------*/
//funciones

func init() { gotenv.Load() }
func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	ruta := mux.NewRouter()

	ruta.HandleFunc("/", index).Methods("Get")
	ruta.HandleFunc("/turno/{id}", detalle).Methods("Get")
	ruta.HandleFunc("/login", login).Methods("Post")
	ruta.HandleFunc("/singup", singup).Methods("Post")
	ruta.HandleFunc("/turnos", turnos).Methods("Get")
	ruta.HandleFunc("/turnos", nuevoturno).Methods("Post")
	ruta.HandleFunc("/turnos", modturnos).Methods("Put")
	ruta.HandleFunc("/turno/{id}", borrarturno).Methods("Delete")

	fmt.Println("usar el puerto 8000")
	log.Fatal(http.ListenAndServe(":8000", ruta))
}

// IDEA: Endpoints
func index(w http.ResponseWriter, r *http.Request) {

}

func detalle(w http.ResponseWriter, r *http.Request) {
	var a Turno
	parametros := mux.Vars(r)

	rows := db.QueryRow("select * from turnos where id_turno=$1", parametros["id"])
	err := rows.Scan(&a.ID, &a.Fecha, &a.Hora, &a.Veterinario, &a.Dueno, &a.Mascota)
	logFatal(err)
	json.NewEncoder(w).Encode(a)
}

func login(w http.ResponseWriter, r *http.Request) {

}

func singup(w http.ResponseWriter, r *http.Request) {

}

func turnos(w http.ResponseWriter, r *http.Request) {
	var a Turno
	agenda = []Turno{}

	rows, err := db.Query("select * from turnos")
	logFatal(err)
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&a.ID, &a.Fecha, &a.Hora, &a.Veterinario, &a.Dueno, &a.Mascota)
		logFatal(err)

		agenda = append(agenda, a)
	}
	json.NewEncoder(w).Encode(agenda)
}

func nuevoturno(w http.ResponseWriter, r *http.Request) {

	var a Turno
	var aID int
	json.NewDecoder(r.Body).Decode(&a)
	err := db.QueryRow("insert into turnos (fecha, hora, id_usuario,id_cliente,id_mascota)	values ($1,$2,$3,$4,$5)	RETURNING id_turno;", a.Fecha, a.Hora, a.Veterinario, a.Dueno, a.Mascota).Scan(&aID)
	logFatal(err)

	json.NewEncoder(w).Encode(aID)
}

func modturnos(w http.ResponseWriter, r *http.Request) {
	var a Turno
	json.NewDecoder(r.Body).Decode(&a)
	resultado, err := db.Exec("update turnos set fecha=$1, hora=$2, id_usuario=$3,id_cliente=$4,id_mascota=$5 where id_turno=$6 RETURNING id_turno", &a.Fecha, &a.Hora, &a.Veterinario, &a.Dueno, &a.Mascota, &a.ID)

	rowsUpdate, err := resultado.RowsAffected()
	logFatal(err)
	json.NewEncoder(w).Encode(rowsUpdate)
}

func borrarturno(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)
	resultado, err := db.Exec("delete from turnos where id_turno=$1", parametro["id"])
	logFatal(err)
	rowsDelete, err := resultado.RowsAffected()
	logFatal(err)
	json.NewEncoder(w).Encode(rowsDelete)
}

package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"github.com/subosito/gotenv"
)

type Turno struct {
	ID          int    `json:"id"`
	Horario     string `json:"horario"`
	Veterinario string `json:"veterinario"`
	Mascota     string `json:"mascota"`
	Dueno       string `json:"dueno"`
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
	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	logFatal(err)

	//	log.Println(pgUrl)
	db, err = sql.Open("postgres", pgUrl)
	logFatal(err)
	err = db.Ping()
	logFatal(err)
	ruta := mux.NewRouter()

	ruta.HandleFunc("/", index).Methods("Get")
	ruta.HandleFunc("/turno/{id}", detalle).Methods("Get")
	ruta.HandleFunc("/login", login).Methods("Get")
	ruta.HandleFunc("/login", login).Methods("Post")
	ruta.HandleFunc("/singup", singup).Methods("Get")
	ruta.HandleFunc("/singup", singup).Methods("Post")
	ruta.HandleFunc("/turnos", turnos).Methods("Get")
	ruta.HandleFunc("/turnos", nuevoturno).Methods("Post")
	ruta.HandleFunc("/turnos", modturnos).Methods("Put")
	ruta.HandleFunc("/turno/{id}", borrarturno).Methods("Delete")

	log.Fatal(http.ListenAndServe(":8000", ruta))
}

// IDEA: Endpoints
func index(w http.ResponseWriter, r *http.Request) {

}

func detalle(w http.ResponseWriter, r *http.Request) {

}

func login(w http.ResponseWriter, r *http.Request) {

}

func singup(w http.ResponseWriter, r *http.Request) {

}

func turnos(w http.ResponseWriter, r *http.Request) {

}

func nuevoturno(w http.ResponseWriter, r *http.Request) {
}

func modturnos(w http.ResponseWriter, r *http.Request) {
}

func borrarturno(w http.ResponseWriter, r *http.Request) {

}

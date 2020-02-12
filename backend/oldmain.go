package main

import (
	_ "database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
	_ "github.com/subosito/gotenv"

	"github.com/gorilla/mux"
)

/*--------------------------------------------------------------------------*/
// IDEA: Datos hardcode
type Turno struct {
	ID          int    `json:"id"`
	Horario     string `json:"horario"`
	Veterinario string `json:"veterinario"`
	Mascota     string `json:"mascota"`
	Dueno       string `json:"dueno"`
}

var agenda []Turno

/*--------------------------------------------------------------------------*/
//funciones
func main() {
	ruta := mux.NewRouter()

	agenda = append(agenda, Turno{ID: 1, Horario: "11:30", Veterinario: "Juan", Mascota: "Lola", Dueno: "Ignacio"},
		Turno{ID: 2, Horario: "12:30", Veterinario: "Juan", Mascota: "Nicolas", Dueno: "Guido"},
		Turno{ID: 3, Horario: "13:30", Veterinario: "Juan", Mascota: "Carmin", Dueno: "Esteban"},
		Turno{ID: 4, Horario: "13:30", Veterinario: "Ignacio", Mascota: "Martin", Dueno: "Luciano"})

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
	log.Println("index")
}

func detalle(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)
	conv, _ := strconv.Atoi(parametro["id"])
	for _, Turno := range agenda {
		if Turno.ID == conv {
			json.NewEncoder(w).Encode(&Turno)
		}
		/*  else if Turno.ID!=conv{
		    json.NewEncoder(w).Encode("No hay turno registrado")
		  } */

	}
}

func login(w http.ResponseWriter, r *http.Request) {
	log.Println("login")
}

func singup(w http.ResponseWriter, r *http.Request) {
	log.Println("singup")
}

func turnos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(agenda)
}

func nuevoturno(w http.ResponseWriter, r *http.Request) {
	var nturno Turno
	json.NewDecoder(r.Body).Decode(&nturno)
	agenda = append(agenda, nturno)
	json.NewEncoder(w).Encode(agenda)
}

func modturnos(w http.ResponseWriter, r *http.Request) {
	var modturno Turno
	json.NewDecoder(r.Body).Decode(&modturno)

	for i, item := range agenda {
		if item.ID == modturno.ID {
			agenda[i] = modturno
		}
	}

}

func borrarturno(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)
	conv, _ := strconv.Atoi(parametro["id"])
	for i, Turno := range agenda {
		if Turno.ID == conv {
			agenda = append(agenda[:i], agenda[i+1:]...)
		}
		/*  else if Turno.ID!=conv{
		    json.NewEncoder(w).Encode("No hay turno registrado")
		  } */

	}

	json.NewEncoder(w).Encode(agenda)
}

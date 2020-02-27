package controllers

import (
	"TallerVI/models"
	"TallerVI/repository"
	"TallerVI/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func enableCors(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

//	enableCors(&w, r) hack: agregar linea para habilitar el cors
type Controller struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var agenda []models.Turno
var turnosrep = repository.RepositorioTurnos{}

func (c Controller) Turnos(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w, r)
		var a models.Turno
		var error models.Error
		agenda = []models.Turno{}
		agenda, err := turnosrep.VerTurnos(db, a, agenda)
		if err != nil {
			error.Mensaje = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, agenda)
	}
}

func (c Controller) Detalle(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w, r)
		var a models.Turno
		var error models.Error

		parametros := mux.Vars(r)
		agenda = []models.Turno{}
		//turnosrep := repository.RepositorioTurnos{}
		id, _ := strconv.Atoi(parametros["id"])
		a, err := turnosrep.DetalleTurno(db, a, id)

		if err != nil {
			if err == sql.ErrNoRows {
				error.Mensaje = "Turno no encontrado"
				utils.SendError(w, http.StatusNotFound, error)
				return
			} else {
				error.Mensaje = "Server error"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, a)
	}

}

func (c Controller) ModTurno(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w, r)
		var a models.Turno
		var error models.Error

		json.NewDecoder(r.Body).Decode(&a)

		if a.ID == 0 || a.Fecha == "" || a.Hora == "" || a.Veterinario == 0 || a.Dueno == 0 || a.Mascota == 0 {
			error.Mensaje = "Error. Por favor complete todos los campos."
			utils.SendError(w, http.StatusBadRequest, error) //400
			return
		}

		//turnosrep := repository.RepositorioTurnos{}
		rowsUpdated, err := turnosrep.ModTurno(db, a)

		if err != nil {
			error.Mensaje = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error) //500
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, rowsUpdated)
	}
}

func (c Controller) DelTurno(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w, r)
		var error models.Error
		parametro := mux.Vars(r)
		//turnosrep := repository.RepositorioTurnos{}
		id, _ := strconv.Atoi(parametro["id"])
		rowsDeleted, err := turnosrep.BorTurno(db, id)

		if err != nil {
			error.Mensaje = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error) //500
			return
		}
		if rowsDeleted == 0 {
			error.Mensaje = "Turno no encotrado"
			utils.SendError(w, http.StatusInternalServerError, error) //404
			return
		}
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, rowsDeleted)

	}
}

//nuevas funciones
//------------------------------------------------------------------------------------------------------------------------
//Nuevo turno
func (c Controller) NuevoTurno(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w, r)
		var a models.Turno
		var aID int
		var error models.Error

		json.NewDecoder(r.Body).Decode(&a)

		if a.Fecha == "" || a.Hora == "" || a.Veterinario == 0 || a.Dueno == 0 || a.Mascota == 0 {
			error.Mensaje = "Error. Por favor complete todos los campos."
			utils.SendError(w, http.StatusBadRequest, error) //400
			return
		}
		//turnosrep := repository.RepositorioTurnos{}
		aID, err := turnosrep.NuevoTurno(db, a)

		if err != nil {
			error.Mensaje = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, aID)
	}
}

func (c Controller) RegTurno(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w, r)
		var a models.RegTurno
		var aID int
		var error models.Error

		json.NewDecoder(r.Body).Decode(&a)
		if a.FechaHora == "" || a.Veterinario == "" || a.Cliente == "" || a.Mascota == "" {
			error.Mensaje = "Error. Por favor complete todos los campos"
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}
		aID, err := turnosrep.RegTurno(db, a)
		if err != nil {
			fmt.Println(err)
			error.Mensaje = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, aID)
	}
}

//------------------------------------------------------------------------------------------------------------------------
func (c Controller) NuevoCliente(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w, r)
		var a models.NCliente
		var CliID int
		var error models.Error

		json.NewDecoder(r.Body).Decode(&a)

		if a.Descripcion == "" || a.Documento == "" {
			error.Mensaje = "Error. Por favor complete los campos requeridos."
			utils.SendError(w, http.StatusBadRequest, error) //400
			return
		}
		//turnosrep := repository.RepositorioTurnos{}
		CliID, err := turnosrep.NuevoCliente(db, a)

		if err != nil {
			fmt.Println(err)
			error.Mensaje = "Server error. Cliente ya cargado"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, CliID)
	}

}

func (c Controller) NuevaMascota(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w, r)
		var m models.NMascota
		var MasID int
		var error models.Error

		json.NewDecoder(r.Body).Decode(&m)

		if m.Descripcion == "" || m.IDCliente == 0 {
			error.Mensaje = "Error. Por favor complete los campos requeridos."
			utils.SendError(w, http.StatusBadRequest, error) //400
			return
		}
		//turnosrep := repository.RepositorioTurnos{}
		MasID, err := turnosrep.NuevaMascota(db, m)

		if err != nil {
			fmt.Println(err)
			error.Mensaje = "Server error. Cliente ya cargado"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, MasID)
	}

}

func (c Controller) NuevoUsuario(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w, r)
		var n models.NUsuario
		var UserID int
		var error models.Error

		json.NewDecoder(r.Body).Decode(&n)

		if n.Descripcion == "" || n.Documento == "" {
			error.Mensaje = "Error. Por favor complete los campos requeridos."
			utils.SendError(w, http.StatusBadRequest, error) //400
			return
		}
		//turnosrep := repository.RepositorioTurnos{}
		UserID, err := turnosrep.NuevoUsuario(db, n)

		if err != nil {
			fmt.Println(err)
			error.Mensaje = "Server error. Cliente ya cargado"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, UserID)
	}

}

var clientes []models.NUsuario

func (c Controller) ListadoClientes(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w, r)
		var a models.NUsuario
		var error models.Error
		clientes = []models.NUsuario{}
		clientes, err := turnosrep.ListadoClientes(db, a, clientes)
		if err != nil {
			error.Mensaje = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, clientes)

	}
}

var mascotas []models.NMascota

func (c Controller) Mascotas(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w, r)
		var a models.NMascota
		var error models.Error
		mascotas = []models.NMascota{}
		mascotas, err := turnosrep.Mascotas(db, a, mascotas)
		if err != nil {
			error.Mensaje = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, mascotas)
	}
}

var usuarios []models.NUsuario

func (c Controller) Usuarios(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w, r)
		var a models.NUsuario
		var error models.Error
		usuarios = []models.NUsuario{}
		usuarios, err := turnosrep.Usuarios(db, a, usuarios)
		if err != nil {
			fmt.Println(err)
			error.Mensaje = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, usuarios)
	}
}

/*
func (c Controller) MascotasClientes(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dmasc models.NMascotas
		var error models.Error

		parametro := mux.Vars(r)
		var cli string
		cli = parametro
		mascotas = []models.NMascotas{}
		mascotas, err := turnosrep.MascotasClientes(db, dmasc, mascotas, cli)
		if err == sql.ErrNoRows {
			error.Mensaje = "Dueño no encontrado"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		} else {
			error.Mensaje = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, dmasc)
	}
}
*/

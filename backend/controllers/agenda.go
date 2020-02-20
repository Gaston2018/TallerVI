package controllers

import (
	"TallerVI/models"
	"TallerVI/repository"
	"TallerVI/utils"
	"database/sql"
	"log"
	"net/http"
)

type Controller struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var agenda []models.Turno

func (c Controller) Turnos(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var a models.Turno
		var error models.Error
		agenda = []models.Turno{}
		turnosrep := repository.RepositorioTurnos{}
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

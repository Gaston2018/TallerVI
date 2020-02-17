package controllers

import (
	"TallerVI/models"
	"database/sql"
	"encoding/json"
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
		agenda = []models.Turno{}

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
}

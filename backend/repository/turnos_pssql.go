package repository

import (
	"TallerVI/models"
	"database/sql"
)

type RepositorioTurnos struct{}

func (b RepositorioTurnos) VerTurnos(db *sql.DB, a models.Turno, agenda []models.Turno) ([]models.Turno, error) {

	rows, err := db.Query("select * from turnos")
	//db.Query("select t.id_turno, t.fecha, t.hora, u.nombre, c.nombre, m.nombre from turnos as t join usuarios as u on t.id_usuario = u.id_usuario join mascotas as m on t.id_mascota = m.id_mascota join clientes as c on t.id_cliente = c.id_cliente")
	if err != nil {
		return []models.Turno{}, err
	}

	for rows.Next() {
		err = rows.Scan(&a.ID, &a.Fecha, &a.Hora, &a.Veterinario, &a.Dueno, &a.Mascota)

		agenda = append(agenda, a)
	}
	if err != nil {
		return []models.Turno{}, err
	}

	return agenda, nil

}

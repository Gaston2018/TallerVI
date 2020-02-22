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

func (b RepositorioTurnos) DetalleTurno(db *sql.DB, a models.Turno, id int) (models.Turno, error) {
	rows := db.QueryRow("select * from turnos where id_turno=$1", id)
	err := rows.Scan(&a.ID, &a.Fecha, &a.Hora, &a.Veterinario, &a.Dueno, &a.Mascota)
	return a, err
}

func (b RepositorioTurnos) NuevoTurno(db *sql.DB, a models.Turno) (int, error) {
	err := db.QueryRow("insert into turnos (fecha, hora, id_usuario,id_cliente,id_mascota)	values ($1,$2,$3,$4,$5)	RETURNING id_turno;", a.Fecha, a.Hora, a.Veterinario, a.Dueno, a.Mascota).Scan(&a.ID)

	if err != nil {
		return 0, err
	}
	return a.ID, nil

}

func (b RepositorioTurnos) ModTurno(db *sql.DB, a models.Turno) (int64, error) {
	result, err := db.Exec("update turnos set fecha=$1, hora=$2, id_usuario=$3,id_cliente=$4,id_mascota=$5 where id_turno=$6 RETURNING id_turno", &a.Fecha, &a.Hora, &a.Veterinario, &a.Dueno, &a.Mascota, &a.ID)

	if err != nil {
		return 0, err
	}

	rowsUpdated, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}
	return rowsUpdated, nil
}

func (b RepositorioTurnos) BorTurno(db *sql.DB, id int) (int64, error) {

	resultado, err := db.Exec("delete from turnos where id_turno=$1", id)

	if err != nil {
		return 0, err
	}

	rowsDelete, err := resultado.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsDelete, nil
}

package repository

import (
	"TallerVI/models"
	"database/sql"
)

type RepositorioTurnos struct{}

//consultas usuarios
func (b RepositorioTurnos) Usuarios(db *sql.DB, a models.NUsuario, usuarios []models.NUsuario) ([]models.NUsuario, error) {
	rows, err := db.Query("Select * from test_usuarios")
	if err != nil {
		return []models.NUsuario{}, err
	}
	for rows.Next() {
		err = rows.Scan(&a.IDUsuario, &a.Descripcion, &a.Telefono, &a.Direccion, &a.Documento)
		usuarios = append(usuarios, a)
	}
	if err != nil {
		return []models.NUsuario{}, err
	}

	return usuarios, nil

}

func (b RepositorioTurnos) NuevoUsuario(db *sql.DB, a models.NUsuario) (int, error) {
	err := db.QueryRow("insert into test_usuarios (descripcion, telefono, direccion, documento) values ($1,$2,$3,$4) RETURNING ID_Usuario", a.Descripcion, a.Telefono, a.Direccion, a.Documento).Scan(&a.IDUsuario)

	if err != nil {
		return 0, err
	}
	return a.IDUsuario, nil

}

func (b RepositorioTurnos) ModUsuario(db *sql.DB, a models.NUsuario) (int64, error) {
	result, err := db.Exec("update test_usuarios set descripcion=$1, telefono=$2, direccion=$3,documento=$4 where id_usuario=$5 RETURNING id_turno", &a.Descripcion, &a.Telefono, &a.Direccion, &a.Documento, &a.IDUsuario)

	if err != nil {
		return 0, err
	}

	rowsUpdated, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}
	return rowsUpdated, nil
}

//consultas clientes
func (b RepositorioTurnos) ListadoClientes(db *sql.DB, a models.NUsuario, clientes []models.NUsuario) ([]models.NUsuario, error) {

	rows, err := db.Query("select * from test_clientes")

	if err != nil {
		return []models.NUsuario{}, err
	}

	for rows.Next() {
		err = rows.Scan(&a.IDUsuario, &a.Descripcion, &a.Telefono, &a.Direccion, &a.Documento)

		clientes = append(clientes, a)
	}
	if err != nil {
		return []models.NUsuario{}, err
	}

	return clientes, nil

}

func (b RepositorioTurnos) NuevoCliente(db *sql.DB, a models.NCliente) (int, error) {
	err := db.QueryRow("insert into test_clientes (descripcion, telefono, direccion, documento) values ($1,$2,$3,$4)	RETURNING ID_Cliente", a.Descripcion, a.Telefono, a.Direccion, a.Documento).Scan(&a.IDCliente)

	if err != nil {
		return 0, err
	}
	return a.IDCliente, nil

}

func (b RepositorioTurnos) ModCliente(db *sql.DB, a models.NCliente) (int64, error) {
	result, err := db.Exec("update test_clientes set descripcion=$1, telefono=$2, direccion=$3,documento=$4 where id_cliente=$5 RETURNING id_turno", &a.Descripcion, &a.Telefono, &a.Direccion, &a.Documento, &a.IDCliente)

	if err != nil {
		return 0, err
	}

	rowsUpdated, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}
	return rowsUpdated, nil
}

//consultas mascotas
func (b RepositorioTurnos) Mascotas(db *sql.DB, a models.NMascota, mascotas []models.NMascota) ([]models.NMascota, error) {
	rows, err := db.Query("select * from test_mascotas")

	if err != nil {
		return []models.NMascota{}, err
	}

	for rows.Next() {
		err = rows.Scan(&a.IDMascota, &a.Descripcion, &a.Tipo, &a.IDCliente)

		mascotas = append(mascotas, a)
	}
	if err != nil {
		return []models.NMascota{}, err
	}

	return mascotas, nil

}

func (b RepositorioTurnos) NuevaMascota(db *sql.DB, a models.NMascota) (int, error) {
	err := db.QueryRow("insert into test_mascotas (descripcion, tipo, id_cliente) values ($1,$2,$3)	RETURNING ID_Cliente", a.Descripcion, a.Tipo, a.IDCliente).Scan(&a.IDMascota)

	if err != nil {
		return 0, err
	}
	return a.IDMascota, nil

}

func (b RepositorioTurnos) ModMascota(db *sql.DB, a models.NMascota) (int64, error) {
	result, err := db.Exec("update test_mascotas set descripcion=$1, tipo=$2, id_cliente=$3 where id_cliente=$4 RETURNING id_turno", &a.Descripcion, &a.Tipo, &a.IDCliente, &a.IDMascota)

	if err != nil {
		return 0, err
	}

	rowsUpdated, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}
	return rowsUpdated, nil
}

//consultas turnos
func (b RepositorioTurnos) VerTurnos(db *sql.DB, a models.RegTurno, agenda []models.RegTurno) ([]models.RegTurno, error) {

	rows, err := db.Query("select * from test_turnos")

	if err != nil {
		return []models.RegTurno{}, err
	}

	for rows.Next() {
		err = rows.Scan(&a.IDturno, &a.FechaHora, &a.Veterinario, &a.Cliente, &a.Mascota)

		agenda = append(agenda, a)
	}
	if err != nil {
		return []models.RegTurno{}, err
	}

	return agenda, nil

}

func (b RepositorioTurnos) DetalleTurno(db *sql.DB, a models.RegTurno, id int) (models.RegTurno, error) {
	rows := db.QueryRow("select * from test_turnos where id_turno=$1", id)
	err := rows.Scan(&a.IDturno, &a.FechaHora, &a.Veterinario, &a.Cliente, &a.Mascota)
	return a, err
}

func (b RepositorioTurnos) RegTurno(db *sql.DB, a models.RegTurno) (int, error) {
	err := db.QueryRow("insert into test_turnos (fechahora, veterinario ,mascota,cliente)	values ($1,$2,$3,$4)	RETURNING id_turno;", a.FechaHora, a.Veterinario, a.Mascota, a.Cliente).Scan(&a.IDturno)

	if err != nil {
		return 0, err
	}
	return a.IDturno, nil
}

func (b RepositorioTurnos) ModTurno(db *sql.DB, a models.RegTurno) (int64, error) {
	result, err := db.Exec("update test_turnos set fechahora=$1, veterinario=$2,mascota=$3,cliente=$4 where id_turno=$5 RETURNING id_turno", a.FechaHora, a.Veterinario, a.Mascota, a.Cliente, &a.IDturno)

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

	resultado, err := db.Exec("delete from test_turnos where id_turno=$1", id)

	if err != nil {
		return 0, err
	}

	rowsDelete, err := resultado.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsDelete, nil
}

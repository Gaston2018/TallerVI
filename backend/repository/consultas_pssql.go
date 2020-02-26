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

func (b RepositorioTurnos) NuevoTurno(db *sql.DB, a models.Turno) (int, error) {
	err := db.QueryRow("insert into turnos (fecha, hora, id_usuario,id_cliente,id_mascota)	values ($1,$2,$3,$4,$5)	RETURNING id_turno;", a.Fecha, a.Hora, a.Veterinario, a.Dueno, a.Mascota).Scan(&a.ID)

	if err != nil {
		return 0, err
	}
	return a.ID, nil

}

func (b RepositorioTurnos) DetalleTurno(db *sql.DB, a models.Turno, id int) (models.Turno, error) {
	rows := db.QueryRow("select * from turnos where id_turno=$1", id)
	err := rows.Scan(&a.ID, &a.Fecha, &a.Hora, &a.Veterinario, &a.Dueno, &a.Mascota)
	return a, err
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

//nuevas funciones

//creacion de clientes
func (b RepositorioTurnos) NuevoCliente(db *sql.DB, a models.NCliente) (int, error) {
	err := db.QueryRow("insert into test_clientes (descripcion, telefono, direccion, documento) values ($1,$2,$3,$4)	RETURNING ID_Cliente", a.Descripcion, a.Telefono, a.Direccion, a.Documento).Scan(&a.IDCliente)

	if err != nil {
		return 0, err
	}
	return a.IDCliente, nil

}

//creacion de mascotas
func (b RepositorioTurnos) NuevaMascota(db *sql.DB, a models.NMascota) (int, error) {
	err := db.QueryRow("insert into test_mascotas (descripcion, tipo, id_cliente) values ($1,$2,$3)	RETURNING ID_Cliente", a.Descripcion, a.Tipo, a.IDCliente).Scan(&a.IDMascota)

	if err != nil {
		return 0, err
	}
	return a.IDMascota, nil

}

//Creacion de usuarios
func (b RepositorioTurnos) NuevoUsuario(db *sql.DB, a models.NUsuario) (int, error) {
	err := db.QueryRow("insert into test_usuarios (descripcion, telefono, direccion, documento) values ($1,$2,$3,$4) RETURNING ID_Usuario", a.Descripcion, a.Telefono, a.Direccion, a.Documento).Scan(&a.IDUsuario)

	if err != nil {
		return 0, err
	}
	return a.IDUsuario, nil

}

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

//listado de Mascotas
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

/* en progreso
//filtar mascotas segun dueño
func (b RepositorioTurnos) MascotasClientes(db *sql.DB, dmasc models.NMascota, mascotas []models.NMascotas, cli string) (models.NMascota, error) {
	rows, err := db.Query("select * from test_mascotas where id_cliente=(select id_cliente from test_clientes	where descripcion='$1')", cli)

	//	if err != nil {
	//		return 0, err
		}

	for rows.Next() {
		err = rows.Scan(&dmasc.IDMascota, &dmasc.Descripcion, &dmasc.Tipo, &dmasc.IDCliente)
		mascotas = append(dmasc, m)

	}
	return m, err
}
*/
/*nuevo input de tunos en progreso

func (b RepositorioTurnos) cargarturnoTurno(db *sql.DB, a models.NuevoTurno) (int, error) {
	//var comodin models.RegTurno
	var vet int
	var cli int
	var masc int

	err := db.QueryRow("select id_usuario from usuarios where descripcion = '$1' RETURNING id_usuario", a.Veterinario).Scan(&vet)

	if err != nil {
		return 104, err
	}

	err = db.QueryRow("select id_cliente from clientes where descripcion = '$1' RETURNING id_cliente", a.Cliente).Scan(&cli)

	if err != nil {
		return 104, err
	}

	err = db.QueryRow("select id_mascota from mascotas where descripcion = '$1' RETURNING id_mascota", a.Mascota).Scan(&masc)

	if err != nil {
		return 104, err
	}

	err = db.QueryRow("insert into test_agenda (fecha_hora, id_usuario,id_cliente,id_mascota)	values ($1,$2,$3,$4,$5)	RETURNING id_turno", a.FechaHora, vet, cli, masc).Scan(&a.IDturno)

	if err != nil {
		return 104, err
	}
	return a.IDturno, nil

}
*/

package models

type NUsuario struct {
	Descripcion string `json:"nombre"`
	Telefono    string `json:"telefono"`
	Direccion   string `json:"direccion"`
	Documento   string `json:"documento"`
	IDUsuario   int    `json:"id_usuario"`
}

type NCliente struct { //Nuevo cliente
	Descripcion string `json:"nombre"`
	Telefono    string `json:"telefono"`
	Direccion   string `json:"direccion"`
	Documento   string `json:"documento"`
	IDCliente   int    `json:"id_cliente"`
}

type NMascota struct { //nueva mascota
	Descripcion string `json:"nombre"`
	Tipo        string `json:"tipo"`
	IDCliente   int    `json:"cliente"`
	IDMascota   int    `json:"id_mascota"`
}

type RegTurno struct { //comodin para registrar
	IDturno     int    `json:"id_turno"`
	FechaHora   string `json:"fecha_hora"`
	Veterinario string `json:"Nombre_Veterinario"`
	Cliente     string `json:"nombre_cliente"`
	Mascota     string `json:"nombre_mascota"`
}

type Comod struct {
	Dato string `json:"dato"`
}

//EN PROCESO
/*
type NuevoTurno2 struct { //recepcion de turo
	FechaHora   string `json:"fecha_hora"`
	Veterinario string `json:"Nombre_Veterinario"`
	Cliente     string `json:"nombre_cliente"`
	Mascota     string `json:"nombre_mascota"`
	IDturno     int    `json:"id_turno"`
}
*/

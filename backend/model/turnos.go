Package models

type Turno struct {
	ID          int    `json:"id_turno"`
	Fecha       string `json:"fecha"`
	Hora        string `json:"hora"`
	Veterinario int    `json:"id_usuario"`
	Dueno       int    `json:"id_cliente"`
	Mascota     int    `json:"id_mascota"`
}

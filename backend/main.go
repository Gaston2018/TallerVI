package main

import (
	"TallerVI/controllers"
	"TallerVI/driver"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var db *sql.DB

func init() {
	gotenv.Load()
}

func main() {

	db = driver.ConnectDB()
	controller := controllers.Controller{}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	ruta := mux.NewRouter()

	ruta.HandleFunc("/", index).Methods("GET")
	//manejo de Turnos
	ruta.HandleFunc("/turnos", controller.Turnos(db)).Methods("Get")          //muestra todos los turnos
	ruta.HandleFunc("/turno/{id}", controller.Detalle(db)).Methods("Get")     //muestra un turno en especifico
	ruta.HandleFunc("/turnos", controller.NuevoTurno(db)).Methods("Post")     //crea un turno
	ruta.HandleFunc("/turnos", controller.ModTurno(db)).Methods("Put")        //actualiza turno, envia id por json
	ruta.HandleFunc("/turno/{id}", controller.DelTurno(db)).Methods("Delete") //borra turno
	//manejo de clientes
	ruta.HandleFunc("/nuevocliente", controller.NuevoCliente(db)).Methods("Post") // creacion de clientes
	//manejo de mascotas
	ruta.HandleFunc("/mascotas", controller.Mascotas(db)).Methods("Get")            //ver mascotas
	ruta.HandleFunc("/nuevamascota", controller.ListadoClientes(db)).Methods("Get") //enviar listado de clientes
	ruta.HandleFunc("/nuevamascota/usuarios", controller.Usuarios(db)).Methods("Get")
	ruta.HandleFunc("/nuevamascota", controller.NuevaMascota(db)).Methods("Post") //creacion de mascotas
	//manejo de usuarios
	ruta.HandleFunc("/nuevoveterinario", controller.NuevoUsuario(db)).Methods("Post") //creacion de usuarios
	//funciones
	//ruta.HandleFunc("/nuevoturno/mascotas/{cliente}", controller.MascotasClientes(db)).Methods("Get") //enviar mascotas en funcion del dueño
	/*Rutas pendientes
	  ruta.HandleFunc("/nuevoturno").Methods("Get") --> envia clientes y veterinarios
	  Fin*/
	fmt.Println("usar el puerto " + port)
	log.Fatal(http.ListenAndServe(":"+port, ruta))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "bienvenidos")
}
func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

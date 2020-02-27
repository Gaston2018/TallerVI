package main

import (
	"TallerVI/controllers"
	"TallerVI/driver"
	"TallerVI/models"
	"TallerVI/utils"
	"database/sql"
	"fmt"
	"io/ioutil"
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
	ruta.HandleFunc("/clima", clima).Methods("Get")
	//manejo de usuarios
	ruta.HandleFunc("/nuevoveterinario", controller.NuevoUsuario(db)).Methods("Post") //creacion de usuarios
	ruta.HandleFunc("/veterinarios", controller.Usuarios(db)).Methods("Get")
	//manejo de clientes
	ruta.HandleFunc("/nuevocliente", controller.NuevoCliente(db)).Methods("Post")
	ruta.HandleFunc("/clientes", controller.ListadoClientes(db)).Methods("Get")
	//manejo de mascotas
	ruta.HandleFunc("/mascotas", controller.Mascotas(db)).Methods("Get")                     //ver mascotas
	ruta.HandleFunc("/nuevamascota/clientes", controller.ListadoClientes(db)).Methods("Get") //enviar listado de clientes
	ruta.HandleFunc("/nuevamascota", controller.NuevaMascota(db)).Methods("Post")            //creacion de mascotas
	//manejo de Turnos
	ruta.HandleFunc("/turnos", controller.Turnos(db)).Methods("Get")                   //muestra todos los turnos
	ruta.HandleFunc("/turno/{id}", controller.Detalle(db)).Methods("Get")              //muestra un turno en especifico
	ruta.HandleFunc("/turnos", controller.RegTurno(db)).Methods("Post")                //crea un turno
	ruta.HandleFunc("/turnos", controller.ModTurno(db)).Methods("Put")                 //actualiza turno, envia id por json
	ruta.HandleFunc("/turno/{id}", controller.DelTurno(db)).Methods("Delete")          //borra turno
	ruta.HandleFunc("/turnos/clientes", controller.ListadoClientes(db)).Methods("Get") //enviar listado de clientes
	ruta.HandleFunc("/turnos/Usuarios", controller.Usuarios(db)).Methods("Get")
	//turnos/Mascotas
	//funciones
	// creacion de clientes
	//	ruta.HandleFunc("/nuevoturno/mascotas/{cliente}", controller.MascotasClientes(db)).Methods("Get") //enviar mascotas en funcion del dueño
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
func clima(w http.ResponseWriter, r *http.Request) {

	var error models.Error
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/forecast?id=3433955&APPID=af4c25b26d16e7216792d5c73ffc6584")
	if err != nil {
		error.Mensaje = "Server error"
		utils.SendError(w, http.StatusInternalServerError, error)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		error.Mensaje = "sin datos"
		utils.SendError(w, http.StatusInternalServerError, error)
	}
	w.Header().Set("Content-Type", "application/json")
	utils.SendSuccess(w, string(body))
}

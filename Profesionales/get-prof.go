package main

import (
	"net/http"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Profesioles struct
{
	nombre string 
}
func main() {
	fmt.Println("Go MySQL")

	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/tallervi")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	router.HandleFunc("/Profesionales", GetProfesionales).Methods("GET")

	//insert, err := db.Query("INSERT INTO profesionales VALUES ('Ezequiel', 'Nanton', '20898', 'paciente')")

	//if err != nil {
	//	panic(err.Error())
	//}

	//defer insert.Close()

	func GetProfesionales(w http.ResponseWriter, req *http.Request)
	{

	}

	fmt.Println("Successfully inserted into profesionales table")

}

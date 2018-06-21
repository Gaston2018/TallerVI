package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go MySQL")

	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/tallervi")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO profesionales VALUES ('Ezequiel', 'Nanton', '20898', 'paciente')")

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	fmt.Println("Successfully inserted into profesionales table")

}

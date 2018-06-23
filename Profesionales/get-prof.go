package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

type Profesional struct {
	Matricula int    `db:"Matricula" json:"Matricula"`
	Nombre    string `db:"Nombre" json:"Nombre"`
	Apellido  string `db:"Apellido" json:"Apellido"`
}

var dbmap = initDb()

func initDb() *gorp.DbMap {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/tallervi")
	checkErr(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbmap.AddTableWithName(Profesional{}, "profesionales").SetKeys(true, "Matricula")
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func main() {

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/profesionales", getProfesionales)

	router.Run()

}

func getProfesionales(c *gin.Context) {

	var profesionales []Profesional
	_, err := dbmap.Select(&profesionales, "SELECT * FROM profesionales")

	if err == nil {
		c.JSON(200, profesionales)
	} else {
		c.JSON(404, gin.H{"error": "no user(s) into the table"})
	}

}

package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

type Profesional struct {
	Matricula    int    `db:"matricula" json:"matricula"`
	Nombre       string `db:"nombre" json:"nombre"`
	Apellido     string `db:"apellido" json:"apellido"`
	Especialidad string `db:"especialidad" json:"especialidad"`
}

var dbmap = initDb()

func initDb() *gorp.DbMap {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/tallervi")
	checkErr(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbmap.AddTableWithName(Profesional{}, "profesionales").SetKeys(true, "matricula")
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
	router.POST("/postProfesionales", postProfesionales)

	router.Run()

}

func getProfesionales(c *gin.Context) {

	especialidad := c.Query("especialidad") // shortcut for c.Request.URL.Query().Get("lastname")
	query := "SELECT * FROM profesionales where 1=1 "
	fmt.Println(especialidad)
	if especialidad != "" {
		query += "and especialidad = '" + especialidad + "'"
	}
	apellido := c.Query("apellido")
	if apellido != "" {
		query += "and apellido = '" + apellido + "'"
	}
	var profesionales []Profesional
	fmt.Println(query)
	_, err := dbmap.Select(&profesionales, query)

	if err == nil {
		c.JSON(200, profesionales)
	} else {
		fmt.Println(err)
		c.JSON(404, gin.H{"error": "no user(s) into the table"})
	}

}

func postProfesionales(c *gin.Context) {
	var matricula Profesional
	c.Bind(&matricula)

	log.Println(matricula)

	if matricula.Nombre != "" && matricula.Apellido != "" {

		if insert, _ := dbmap.Exec(`INSERT INTO profesionales (nombre, apellido, matricula, especialidad) VALUES (?, ?, ?, ?)`, matricula.Nombre, matricula.Apellido, matricula.Matricula, matricula.Especialidad); insert != nil {
			Matricula, err := insert.LastInsertId()
			_ = Matricula
			if err == nil {
				content := &Profesional{
					Matricula:    matricula.Matricula,
					Nombre:       matricula.Nombre,
					Apellido:     matricula.Apellido,
					Especialidad: matricula.Especialidad,
				}
				c.JSON(201, content)
			} else {
				checkErr(err, "Insert failed")
				fmt.Println(err)
			}
		}

	} else {
		c.JSON(400, gin.H{"error": "Fields are empty"})
	}
}

func OptionsUser(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

type Profesional struct {
	Matricula    int64  `db:"matricula" json:"matricula"`
	Nombre       string `db:"nombre" json:"nombre"`
	Apellido     string `db:"apellido" json:"apellido"`
	Especialidad string `db:"especialidad" json:"especialidad"`
	Direccion    string `db:"direccion" json:"direccion"`
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
	//router.GET("/profesionales/:matricula", getProfesionales)
	router.POST("/profesionales", postProfesionales)
	router.DELETE("/profesionales/:matricula", delProfesionales)
	router.PUT("/profesionales/:matricula", putProfesionales)

	router.Run()
	fmt.Println("Levanto")

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
	nombre := c.Query("nombre")
	if nombre != "" {
		query += "and nombre = '" + nombre + "'"
	}
	domicilio := c.Query("domicilio")
	if domicilio != "" {
		query += "and domicilio = '" + domicilio + "'"
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
	var profesional Profesional
	c.BindJSON(&profesional)
	log.Print("matricula: ")
	log.Println(profesional)
	if profesional.Nombre != "" && profesional.Apellido != "" {

		insert, sqlErr := dbmap.Exec(`INSERT INTO profesionales (nombre, apellido, especialidad, direccion) VALUES (?, ?, ?, ?)`, profesional.Nombre, profesional.Apellido, profesional.Especialidad, profesional.Direccion)
		fmt.Println("INSERT")
		fmt.Println(insert)
		fmt.Println(sqlErr)
		var matricula int64
		matricula, err := insert.LastInsertId()
		if err == nil {
			content := &Profesional{
				Nombre:       profesional.Nombre,
				Apellido:     profesional.Apellido,
				Especialidad: profesional.Especialidad,
				Direccion:    profesional.Direccion,
			}
			content.Matricula = matricula
			c.JSON(201, content)
		} else {
			fmt.Println(err)
			checkErr(err, "Insert failed")
		}
	}
}

func delProfesionales(c *gin.Context) {
	id := c.Params.ByName("matricula")

	var profesional Profesional
	err := dbmap.SelectOne(&profesional, "SELECT * FROM profesionales WHERE matricula=?", id)

	if err == nil {
		_, err = dbmap.Delete(&profesional)

		if err == nil {
			c.JSON(200, gin.H{"matricula #" + id: "deleted"})
		} else {
			checkErr(err, "Delete failed")
		}

	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}

	// curl -i -X DELETE http://localhost:8080/api/v1/users/1
}

func putProfesionales(c *gin.Context) {
	id := c.Params.ByName("matricula")
	var profesional Profesional
	err := dbmap.SelectOne(&profesional, "SELECT * FROM profesionales WHERE matricula=?", id)

	if err == nil {
		var json Profesional
		c.Bind(&json)

		matricula, _ := strconv.ParseInt(id, 0, 64)

		user := Profesional{
			Matricula:    matricula,
			Nombre:       json.Nombre,
			Apellido:     json.Apellido,
			Especialidad: json.Especialidad,
			Direccion:    json.Direccion,
		}

		if user.Nombre != "" && user.Apellido != "" {
			_, err = dbmap.Update(&matricula)

			if err == nil {
				c.JSON(200, user)
			} else {
				checkErr(err, "Updated failed")
			}

		} else {
			c.JSON(400, gin.H{"error": "fields are empty"})
		}

	} else {
		fmt.Println(err)
		c.JSON(404, gin.H{"error": "user not found"})
	}

	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://localhost:8080/api/v1/users/1
}

func OptionsUser(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}

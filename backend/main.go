package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := paths()
	r.Run(":3000")
}

func paths() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	//rutas de acceso
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "Index.html", nil)
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "Login.html", nil)
	})

	r.GET("/turnos", func(c *gin.Context) {
		c.HTML(http.StatusOK, "turnos.html", nil)
	})
	r.Static("/public", "./public")

	return r

}

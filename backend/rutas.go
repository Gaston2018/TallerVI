/*
package main
import (
  "github.com/gin-gonic/gin"
  "net/http"
)


func paths() *gin.Engine{
  r:=gin.Default()
  r.LoadHTMLGlob("templates/*.html")
  //rutas de acceso
  r.GET("/", func(c *gin.Context){
    c.HTML(http.StatusOK, "Index.html", nil)
  })

  r.GET("/login", func (c *gin.Context)  {
      c.HTML(http.StatusOK, "Login.html", nil)
  })

  r.GET("/turnos", func (c *gin.Context)  {
    c.HTML(http.StatusOK,"turnos.html", nil)
  })
  r.Static("/public", "./public")

  return r

}
*/

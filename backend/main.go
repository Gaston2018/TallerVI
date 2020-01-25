package main
import (
  "github.com/gin-gonic/gin"
  "net/http"
)
func main(){
  r:=gin.Default ()
  r.LoadHTMLGlob("Frontend/**/*.html")
  r.GET("/", func(c *gin.Context){
    c.HTML(http.StatusOK, "Index.html", nil)
  })

  r.GET("/login", func (c *gin.Context)  {
      c.HTML(http.StatusOK, "Login.html", nil)
  })

  r.Run(":3000")
}

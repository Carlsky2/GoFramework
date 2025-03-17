package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hola mundo")
	})

	r.GET("/saludo/:nombre", func(c *gin.Context) { //ruta de la pagina
		nombre := c.Param("nombre")                //obtener el parametro nombre
		c.String(http.StatusOK, "hola %s", nombre) //salida de texto plano
	})
}

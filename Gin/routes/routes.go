package routes

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// una estrucura de tipo usuario
// con dos campos nombre y email
type Usuario struct {
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
}

// arreglo de usuarios
//var usuarios []Usuario

func SetRoutes(r *gin.Engine) {
	//servir contenido estatico
	r.Static("/static", "./static")

	//cargar plantillas
	r.LoadHTMLGlob("templates/*.html")

	//peticiones

	//rutas estaticas
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	//iterar cada pagina que esta pasando
	r.GET("/:page", func(c *gin.Context) {
		page := c.Param("page")

		if !strings.HasSuffix(page, ".html") { // si no tienen la extension html se la agregamos
			page += ".html"
		}

		if _, err := os.Stat("templates/" + page); err == nil { // si la pagina existe
			c.HTML(http.StatusOK, page, nil) //cargamos la pagina
		} else {
			c.HTML(http.StatusNotFound, "404.html", nil) // si no existe cargamos la pagina 404
		}
	})

	//rutas dinamicas
	/*
		r.GET("/saludo/:nombre", func(c *gin.Context) { //ruta de la pagina
			nombre := c.Param("nombre")                //obtener el parametro nombre
			c.String(http.StatusOK, "hola %s", nombre) //salida de texto plano
		})

		//rutas con post
		r.POST("/usuarios", func(c *gin.Context) { //defimos la ruta con post  este
			var nuevoUsuario Usuario

			if err := c.BindJSON(&nuevoUsuario); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Error al codificar el json"})
				return
			}

			if nuevoUsuario.Nombre == "" || nuevoUsuario.Email == "" {
				c.JSON(http.StatusBadRequest, gin.H{"Error": "Nombre y email son requeridos"})
			}

			usuarios = append(usuarios, nuevoUsuario)

			c.JSON(http.StatusOK, gin.H{"mensaje": "Usuario registrado", "datos": usuarios})

		})
	*/

}

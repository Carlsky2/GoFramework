package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// una estrucura de tipo usuario
// con dos campos nombre y email
type Usuario struct {
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
}

// arreglo de usuarios
var usuarios []Usuario

func SetRoutes(r *gin.Engine) {
	//peticiones

	//rutas estaticas
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hola mundo")
	})

	//rutas dinamicas
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
}

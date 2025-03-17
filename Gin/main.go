package main

import (
	"Gin/routes"

	"github.com/gin-gonic/gin" // importar el framework de gin
)

func main() {
	r := gin.Default() //define la ruta por defecto del proyecto
	routes.SetRoutes(r)
	r.Run(":8080") // url del proyecto en este caso localhost:8080
}

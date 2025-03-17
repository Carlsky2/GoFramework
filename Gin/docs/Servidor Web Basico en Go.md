# sercvidor web basico en go

'''go
package main

import (
	"github.com/gin-gonic/gin" // importar el framework de gin
)

func main() {
	r := gin.Default() //define la ruta por defecto del proyecto

	r.GET("/", func(c *gin.Context) { //ruta de la pagina
		//c.String(200, "hola mundo") //primera ruta, escuchando una ruta de la raiz del proyecto
		//200 es el codigo de estado de la pagina

		//salida con json - esto es util para APIS
		c.JSON(200, gin.H{
			"message": "hola mundo"})
		})
		
	r.Run(":8080") // url del proyecto en este caso localhost:8080
}
'''
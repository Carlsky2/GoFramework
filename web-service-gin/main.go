package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album representa la información de un álbum.
type album struct {
	ID      string  `json:"id"`
	Titulo  string  `json:"titulo"`
	Artista string  `json:"artista"`
	Precio  float64 `json:"precio"`
}

// Slice de albums para almacenar datos de álbumes de ejemplo.
var albums = []album{
	{ID: "1", Titulo: "Kenopsia", Artista: "quannnic", Precio: 32000},
	{ID: "2", Titulo: "VOL.1", Artista: "Novulent", Precio: 64500},
	{ID: "3", Titulo: "Exodo", Artista: "Peso deidad", Precio: 128000},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

// getAlbums responde con la lista de todos los álbumes en formato JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums agrega un álbum a partir del JSON recibido en el cuerpo de la solicitud.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Llama a BindJSON para enlazar el JSON recibido a la variable newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Agrega el nuevo álbum al slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID localiza el álbum cuyo valor de ID coincide con el parámetro
// enviado por el cliente y luego devuelve ese álbum como respuesta.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Recorre la lista de álbumes buscando un álbum cuyo valor de ID
	// coincida con el parámetro.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "álbum no encontrado"})
}

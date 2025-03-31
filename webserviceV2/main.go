package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album representa la información de un álbum.
type album struct {
	ID      string  `json:"id" binding:"required"`
	Titulo  string  `json:"titulo" binding:"required"`
	Artista string  `json:"artista" binding:"required"`
	Precio  float64 `json:"precio" binding:"required,gt=0"`
}

// Slice de albums para almacenar datos de ejemplo.
var albums = []album{
	{ID: "1", Titulo: "Kenopsia", Artista: "quannnic", Precio: 32000},
	{ID: "2", Titulo: "VOL.1", Artista: "Novulent", Precio: 64500},
	{ID: "3", Titulo: "Exodo", Artista: "Peso deidad", Precio: 128000},
}

func main() {
	router := gin.Default()
	router.Use(gin.Logger(), gin.Recovery())

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.PUT("/albums/:id", updateAlbum)
	router.DELETE("/albums/:id", deleteAlbum)

	router.Run("localhost:8080")
}

// getAlbums responde con la lista de todos los álbumes en formato JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums agrega un nuevo álbum con validación.
func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID busca un álbum por su ID.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "álbum no encontrado"})
}

// updateAlbum modifica un álbum existente.
func updateAlbum(c *gin.Context) {
	id := c.Param("id")
	var updatedAlbum album

	if err := c.ShouldBindJSON(&updatedAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, a := range albums {
		if a.ID == id {
			albums[i] = updatedAlbum
			c.IndentedJSON(http.StatusOK, updatedAlbum)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "álbum no encontrado"})
}

// deleteAlbum elimina un álbum por su ID.
func deleteAlbum(c *gin.Context) {
	id := c.Param("id")

	for i, a := range albums {
		if a.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "álbum eliminado"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "álbum no encontrado"})
}

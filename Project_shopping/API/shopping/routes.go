package shopping

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Routes() {
	router := gin.Default()
	fmt.Println("In Routes")
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", addAlbum)
	router.PUT("/albums/:id", updateAlbumByID)
	router.DELETE("/albums/:id", deleteAlbumByID)

	router.Run("localhost:8080")
	fmt.Printf("starting server at 8080")

}

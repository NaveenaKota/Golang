package shopping

import (
	"Project_shopping/dataservice"
	"Project_shopping/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var albums []models.Album

/*var albums = []models.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}*/

func getAlbums(c *gin.Context) {
	fmt.Println("To get all the albums present")
	db := dataservice.ConnectToDb()
	albums = dataservice.GetAlbumsFromDb(db, albums)
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			fmt.Println(a.ID, id)
			db := dataservice.ConnectToDb()
			dataservice.GetAlbumByIdFromDb(db, albums, a.ID)
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
func addAlbum(c *gin.Context) {
	var newAlbum models.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	// Add the new album to the slice.
	fmt.Println(newAlbum)
	db := dataservice.ConnectToDb()
	dataservice.PostAddAlbumToDb(db, newAlbum)
	albums = append(albums, newAlbum)
	fmt.Println(albums)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func deleteAlbumByID(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	for index, a := range albums {
		fmt.Println("In for id is:", a.ID)
		if a.ID == id {
			db := dataservice.ConnectToDb()
			dataservice.DeleteAlbumFromDb(db, a.ID)
			albums = append(albums[:index], albums[index+1:]...)
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func updateAlbumByID(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	// Parse the request body to get the updated album data
	var updatedAlbum models.Album
	if err := c.ShouldBindJSON(&updatedAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for index, a := range albums {
		if a.ID == id {
			fmt.Println("In for id is:", a.ID)
			db := dataservice.ConnectToDb()
			dataservice.UpdateAlbumInDb(db, updatedAlbum)
			// Update the album's data with the new data from the request body
			albums[index].Title = updatedAlbum.Title
			albums[index].Artist = updatedAlbum.Artist
			albums[index].Price = updatedAlbum.Price

			c.IndentedJSON(http.StatusOK, albums[index])
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Mona Lisa", Artist: "Leonardo da Vinci", Price: 860_000_000},
	{ID: "2", Title: "The Starry Night", Artist: "Vincent van Gogh", Price: 100_000_000},
	{ID: "3", Title: "The Persistence of Memory", Artist: "Salvador Dalí", Price: 60_000_000},
	{ID: "4", Title: "Girl with a Pearl Earring", Artist: "Johannes Vermeer", Price: 90_000_000},
	{ID: "5", Title: "The Scream", Artist: "Edvard Munch", Price: 120_000_000},
	{ID: "6", Title: "Guernica", Artist: "Pablo Picasso", Price: 200_000_000},
	{ID: "7", Title: "The Night Watch", Artist: "Rembrandt", Price: 500_000_000},
	{ID: "8", Title: "The Kiss", Artist: "Gustav Klimt", Price: 150_000_000},
	{ID: "9", Title: "Water Lilies", Artist: "Claude Monet", Price: 80_000_000},
	{ID: "10", Title: "American Gothic", Artist: "Grant Wood", Price: 8_000_000},
	{ID: "11", Title: "No. 5, 1948", Artist: "Jackson Pollock", Price: 140_000_000},
	{ID: "12", Title: "Portrait of Adele Bloch-Bauer I", Artist: "Gustav Klimt", Price: 135_000_000},
	{ID: "13", Title: "The Arnolfini Portrait", Artist: "Jan van Eyck", Price: 100_000_000},
	{ID: "14", Title: "The Birth of Venus", Artist: "Sandro Botticelli", Price: 200_000_000},
	{ID: "15", Title: "Whistler’s Mother", Artist: "James McNeill Whistler", Price: 30_000_000},
	{ID: "16", Title: "Les Demoiselles d'Avignon", Artist: "Pablo Picasso", Price: 179_000_000},
	{ID: "17", Title: "A Sunday Afternoon on the Island of La Grande Jatte", Artist: "Georges Seurat", Price: 110_000_000},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	fmt.Println("Print ")
	err := c.BindJSON(&newAlbum)
	if err != nil {
		return
	}
	fmt.Println("Print ")
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	var id = c.Param("id")
	var albumList []album
	for _, album := range albums {
		if album.ID == id {
			albumList = append(albumList, album)
		}
	}
	c.IndentedJSON(http.StatusFound, albumList)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

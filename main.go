package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID        string  `json:"id"`
	Title     string  `json:"title"`
	Artist    string  `json:"artist"`
	Price     float64 `json:"price"`
	Published int
}

// albums slice (Array) to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	// Initialize a Gin router using Default.
	router := gin.Default()

	// Use the GET function to associate the GET HTTP method and /albums path with a handler function.
	// Note that you’re passing the name of the getAlbums function. This is different from passing the result of the function, which you would do by passing getAlbums() (note the parenthesis).
	router.GET("/albums", getAlbums)

	// Use the Run function to attach the router to an http.Server and start the server.
	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	// Call Context.IndentedJSON to serialize the struct into JSON and add it to the response.

	// The function’s first argument is the HTTP status code you want to send to the client. Here, you’re passing the StatusOK constant from the net/http package to indicate 200 OK.

	// Note that you can replace Context.IndentedJSON with a call to Context.JSON to send more compact JSON. In practice, the indented form is much easier to work with when debugging and the size difference is usually small.
	c.IndentedJSON(http.StatusOK, albums)
}

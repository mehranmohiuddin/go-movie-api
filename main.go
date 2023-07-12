package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mehranmohiuddin/go-movie-api/model"
)

func main() {
	r := gin.Default()

	movie := model.Movie{
		Name:     "Avengers: Endgame",
		Director: "Russo Brothers",
		Year:     2019,
	}

	var movies = []model.Movie{
		movie,
	}

	r.GET("/movies", func(c *gin.Context) {
		c.JSON(http.StatusOK, movies)
	})
	r.Run()
}

package controllers

import (
	"github.com/abdulbasit108/go-postgresql-basic-api/initializers"
	"github.com/abdulbasit108/go-postgresql-basic-api/models"
	"github.com/gin-gonic/gin"
)

func CreateMovie(c *gin.Context) {

	var body struct {
		Title    string
		Year     int
		Genre    string
		Director string
	}

	c.Bind(&body)

	movie := models.Movies{Title: body.Title, Year: body.Year, Genre: body.Genre, Director: body.Director}
	result := initializers.DB.Create(&movie)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{
		"movie": movie,
	})
}

func GetMovies(c *gin.Context) {
	var movies []models.Movies
	initializers.DB.Find(&movies)

	c.JSON(200, gin.H{
		"movies": movies,
	})
}

func GetMovie(c *gin.Context) {
	id := c.Param("id")

	var movie models.Movies
	initializers.DB.First(&movie, id)

	c.JSON(200, gin.H{
		"movie": movie,
	})
}

func UpdateMovie(c *gin.Context) {
	// Get ID from Params
	id := c.Param("id")

	// Get Body from Request
	var body struct {
		Title    string
		Year     int
		Genre    string
		Director string
	}
	c.Bind(&body)

	// Find Object
	var movie models.Movies
	initializers.DB.First(&movie, id)

	// Update Object
	initializers.DB.Model(&movie).Updates(models.Movies{Title: body.Title, Year: body.Year, Genre: body.Genre, Director: body.Director})

	c.JSON(200, gin.H{
		"movie": movie,
	})
}

func DeleteMovie(c *gin.Context) {
	// Get ID from Params
	id := c.Param("id")

	initializers.DB.Delete(&models.Movies{}, id)

	c.Status(200)

}

package routes

import (
	"github.com/abdulbasit108/go-postgresql-basic-api/controllers"
	"github.com/gin-gonic/gin"
)

func MovieRoutes(r *gin.Engine) {
	r.POST("/createMovie", controllers.CreateMovie)

	r.GET("/getMovies", controllers.GetMovies)
	r.GET("getMovie/:id", controllers.GetMovie)

	r.PUT("updateMovie/:id", controllers.UpdateMovie)

	r.DELETE("/deleteMovie/:id", controllers.DeleteMovie)
}

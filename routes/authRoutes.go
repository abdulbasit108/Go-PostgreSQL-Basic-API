package routes

import (
	"github.com/abdulbasit108/go-postgresql-basic-api/controllers"
	"github.com/abdulbasit108/go-postgresql-basic-api/middlewares"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/signup", controllers.Signup)
	r.POST("/signin", controllers.Signin)
	r.GET("/testAuth", middlewares.RequireAuth, controllers.Validate)
}

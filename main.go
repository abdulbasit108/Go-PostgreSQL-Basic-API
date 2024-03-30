package main

import (
	"github.com/abdulbasit108/go-postgresql-basic-api/initializers"
	"github.com/abdulbasit108/go-postgresql-basic-api/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.ApplyMigrations()
}

func main() {
	r := gin.Default()

	routes.MovieRoutes(r)

	routes.AuthRoutes(r)

	r.Run()
}

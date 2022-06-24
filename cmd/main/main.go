package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/hplxsarthak/go-sql-project/pkg/routes"
	"github.com/gin-contrib/cors"

)

func main() {
	// Create a Server
	r := gin.Default()

	// Middleware for cors
	r.Use(cors.Default())

	// Call all the routes created in routes folder
	routes.RegisterMedRoutes(r)

	// Start the Server
	r.Run()
}
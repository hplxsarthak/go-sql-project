package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hplxsarthak/go-sql-project/pkg/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// Create a Server
	r := gin.Default()

	// Middleware for cors
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET","POST","PUT", "PATCH","DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	routes.LoginRoutes(r)

	// Call all the routes created in routes folder for med
	routes.RegisterMedRoutes(r)

	// Call all the routes created in routes folder for user
	routes.RegisterUsersRoutes(r)

	

	// Start the Server
	r.Run()
}

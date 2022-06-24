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

	r.Use(cors.Default())
	routes.RegisterMedRoutes(r)

	// Start the Server
	r.Run()
}
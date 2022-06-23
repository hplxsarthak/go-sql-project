package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/hplxsarthak/go-sql-project/pkg/routes"

)

func main() {
	// Create a Server
	r := gin.Default()
	routes.RegisterMedRoutes(r)

	// Start the Server
	r.Run()
}
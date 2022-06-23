package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hplxsarthak/go-sql-project/pkg/controllers"
)

// Function which contains all routes for our purpose
var RegisterMedRoutes = func (router *gin.Engine) {
	router.GET("/med/", controllers.GetMed)
	router.POST("/med/", controllers.CreateMed)
	router.GET("/med/:id", controllers.GetMedById)
	router.PUT("/med/:id", controllers.UpdateMed)
	router.DELETE("/med/:id", controllers.DeleteMed)
}
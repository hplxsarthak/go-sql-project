package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hplxsarthak/go-sql-project/pkg/controllers"
	"github.com/hplxsarthak/go-sql-project/pkg/middleware"
)

// Function which contains all routes for our purpose
var RegisterMedRoutes = func (router *gin.Engine) {
	router.GET("/med/",middleware.ValidateToken(),controllers.GetMed)
	router.POST("/med/", middleware.ValidateToken(),controllers.CreateMed)
	router.GET("/med/:id", middleware.ValidateToken(),controllers.GetMedById)
	router.PUT("/med/:id",middleware.ValidateToken(), controllers.UpdateMed)
	router.DELETE("/med/:id", middleware.ValidateToken(),controllers.DeleteMed)

	// router.GET("/med/search", controllers.SearchMed)
	router.GET("/med/search&pages",middleware.ValidateToken(), controllers.SearchPageMed)
}
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hplxsarthak/go-sql-project/pkg/controllers"
)

// Function which contains all routes for our purpose
var RegisterUsersRoutes = func (router *gin.Engine) {
	router.GET("/user/", controllers.GetUser)
	router.POST("/user/", controllers.CreateUser)
	router.GET("/user/:id", controllers.GetUserById)
	router.PUT("/user/:id", controllers.UpdateUser)
	router.DELETE("/user/:id", controllers.DeleteUser)

	// router.GET("/med/search", controllers.SearchMed)
	router.GET("/user/search&pages", controllers.SearchPageMed)
}
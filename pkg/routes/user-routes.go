package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hplxsarthak/go-sql-project/pkg/controllers"
	"github.com/hplxsarthak/go-sql-project/pkg/middleware"
)

// Function which contains all routes for our purpose
var RegisterUsersRoutes = func (router *gin.Engine) {
	router.GET("/user/", middleware.ValidateRole(),controllers.GetUser)
	router.POST("/user/", middleware.ValidateRole(),controllers.CreateUser)
	router.GET("/user/:id", middleware.ValidateRole(),controllers.GetUserById)
	router.PUT("/user/:id", middleware.ValidateRole(),controllers.UpdateUser)
	router.DELETE("/user/:id", middleware.ValidateRole(),controllers.DeleteUser)

	// router.GET("/user/search", controllers.SearchUser)
	router.GET("/user/search&pages", middleware.ValidateRole(),controllers.SearchPageUser)

}
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hplxsarthak/go-sql-project/pkg/controllers"
	"github.com/hplxsarthak/go-sql-project/pkg/middleware"
)

// Function which contains all routes for our purpose
var RegisterUsersRoutes = func (router *gin.Engine) {
	router.GET("/user/", middleware.ValidateToken(),controllers.GetUser)
	router.POST("/user/", middleware.ValidateToken(),controllers.CreateUser)
	router.GET("/user/:id", middleware.ValidateToken(),controllers.GetUserById)
	router.PUT("/user/:id", middleware.ValidateToken(),controllers.UpdateUser)
	router.DELETE("/user/:id", middleware.ValidateToken(),controllers.DeleteUser)

	// router.GET("/user/search", controllers.SearchUser)
	router.GET("/user/search&pages", middleware.ValidateToken(),controllers.SearchPageUser)

}
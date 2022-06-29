package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hplxsarthak/go-sql-project/pkg/controllers"
)

// Function which contains all routes for our purpose
var LoginRoutes = func (router *gin.Engine) {
	router.POST("/api/login", controllers.Login)
	router.POST("/api/logout", controllers.Logout)

}
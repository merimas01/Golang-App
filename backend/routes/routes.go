package routes

import (
	"Golang-App/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, userCtrl *controllers.UserController) {
	// Generic CRUD routes from BaseController
	router.POST("/users", userCtrl.Create)
	router.PUT("/users/:id", userCtrl.Update)
	router.GET("/users", userCtrl.GetAll)
	router.GET("/users/:id", userCtrl.GetByID)
	router.DELETE("/users/:id", userCtrl.Delete)

	// User-specific routes

}

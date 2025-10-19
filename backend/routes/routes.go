package routes

import (
	"Golang-App/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, userCtrl *controllers.UserController) {
	// Generic CRUD routes from BaseController
	router.POST("/users", userCtrl.CreateUser)
	router.PUT("/users/:id", userCtrl.UpdateUser)
	router.GET("/users", userCtrl.GetAllUsers)
	router.GET("/users/:id", userCtrl.GetUserByID)
	router.DELETE("/users/:id", userCtrl.DeleteUser)

	// User-specific routes

}

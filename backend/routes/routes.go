package routes

import (
	"Golang-App/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, useCtrl *controllers.UserController) {
	router.POST("/users", useCtrl.CreateUser)
	router.PUT("/users/:id", useCtrl.UpdateUser)
	router.GET("/users", useCtrl.GetAllUsers)
	router.GET("/users/:id", useCtrl.GetUserByID)
	router.DELETE("/users/:id", useCtrl.DeleteUser)
}

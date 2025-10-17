package main

import (
	"Golang-App/config"
	"Golang-App/controllers"
	"Golang-App/models"
	"Golang-App/repository"
	"Golang-App/routes"
	"Golang-App/services"
	"log"

	_ "Golang-App/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title User Service API
// @version 1.0
// @description User microservice in Go
// @host localhost:8080
// @BasePath /
func main() {
	db := config.Connect()
	db.AutoMigrate(&models.User{})

	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userCtrl := &controllers.UserController{Service: userService}

	router := gin.Default()
	//default route: http://localhost:8080
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello!"})
	})

	routes.RegisterUserRoutes(router, userCtrl)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("User Service running on port 8080")
	router.Run(":8080")
}

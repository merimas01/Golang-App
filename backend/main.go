package main

import (
	"Golang-App/config"
	"Golang-App/controllers"
	"Golang-App/models"
	"Golang-App/routes"
	"Golang-App/services"
	"log"

	_ "Golang-App/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	db := config.Connect()

	err := db.Migrator().HasTable(&models.User{})
	if !err {
		db.AutoMigrate(&models.User{})
	}

	// Generic repository
	userRepo := &services.BaseService[models.User, models.UserInsert, models.UserUpdate]{DB: db}

	// UserService
	userService := services.NewUserService(userRepo)

	// UserController using the new constructor
	userCtrl := controllers.NewUserController(userService)

	router := gin.Default()

	// Default route
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello!"})
	})

	// Register user routes
	routes.RegisterUserRoutes(router, userCtrl)

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("User Service running on port 8080")
	router.Run(":8080")
}

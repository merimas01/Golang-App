package main

import (
	"Golang-App/config"
	"Golang-App/controllers"
	"Golang-App/models"
	searchobjects "Golang-App/models/search_objects"
	"Golang-App/routes"
	"Golang-App/seed"
	"Golang-App/services"
	"fmt"
	"log"

	_ "Golang-App/docs"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	config.LoadConfig()

	dsnNoDB := config.GetDSNWithoutDB()

	tempDB, err := gorm.Open(mysql.Open(dsnNoDB), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to MySQL server:", err)
	}

	dbName := config.GetDbName()
	createDBQuery := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName)
	if err := tempDB.Exec(createDBQuery).Error; err != nil {
		log.Fatal("Failed to create database:", err)
	}

	dsn := config.GetDSN()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	fmt.Println("Connected to database:", dbName)

	// ovde se dodaju sve tabele koje zelimo migrirati u bazu
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to auto-migrate:", err)
	}

	if viper.GetString("SEED_DATA") == "true" {
		seed.SeedUsers(db)
	}

	userRepo := &services.BaseService[models.User, models.UserInsert, models.UserUpdate, searchobjects.BaseSearchObject]{DB: db}
	userService := services.NewUserService(userRepo)
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

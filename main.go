package main

import (
	"ddtservice_agri/controller"
	"ddtservice_agri/database"
	"ddtservice_agri/middleware"
	"ddtservice_agri/schema"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&schema.User{})
	database.Database.AutoMigrate(&schema.Role{})
	database.Database.AutoMigrate(&schema.Employee{})
	database.Database.AutoMigrate(&schema.Estate{})
	database.Database.AutoMigrate(&schema.Division{})
	database.Database.AutoMigrate(&schema.UserRole{})
	database.Database.AutoMigrate(&schema.UserDivision{})
	database.Database.AutoMigrate(&schema.Articles{})
}

func serveApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register())
	publicRoutes.POST("/login", controller.Login)
	publicRoutes.GET("/profile/:ID", controller.GetUserProfile)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.POST("/article", controller.AddEntry)
	protectedRoutes.POST("/user/division/:ID", controller.UserAddDivisions)
	protectedRoutes.POST("/employee", controller.AddEmployee)

	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}

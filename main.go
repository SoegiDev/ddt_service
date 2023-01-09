package main

import (
	"ddtservice_agri/controller"
	"ddtservice_agri/database"
	"ddtservice_agri/middleware"
	"ddtservice_agri/schema"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
	testGetDateTime()
}

func testGetDateTime() {
	current_time := time.Now()

	// individual elements of time can
	// also be called to print accordingly
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		current_time.Year(), current_time.Month(), current_time.Day(),
		current_time.Hour(), current_time.Minute(), current_time.Second())

	// formatting time using
	// custom formats
	fmt.Println(current_time.Format("2006-01-02 15:04:05"))
	fmt.Println(current_time.Format("2006-January-02"))
	fmt.Println(current_time.Format("2006-01-02 3:4:5 pm"))

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
	database.Database.AutoMigrate(&schema.Company{})
	database.Database.AutoMigrate(&schema.Article{})
	database.Database.AutoMigrate(&schema.Application{})
	database.Database.AutoMigrate(&schema.ActivityLog{})
	database.Database.AutoMigrate(&schema.RoleApplication{})
	database.Database.AutoMigrate(&schema.Account{})
	database.Database.AutoMigrate(&schema.Application{})
	database.Database.AutoMigrate(&schema.AccountRoleApplication{})

}

func serveApplication() {
	gin.DisableConsoleColor()
	// Logging to a file.
	f, _ := os.Create("logger.log")
	// Use the following code if you need to write the logs to file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(middleware.JSONLogMiddleware())

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)
	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	//User
	protectedRoutes.GET("/user/:ID", controller.UserGetProfiles)
	protectedRoutes.POST("/user/:ID/assign_roles", controller.UserAssignRole)
	protectedRoutes.POST("/user/:ID/assign_role_app", controller.UserAssignRoleApplication)
	protectedRoutes.POST("/user/:ID/assign_division", controller.UserAssignDivisions)
	protectedRoutes.PATCH("/user/:ID/account_edit", controller.UserAccountUpdate)
	// Article
	protectedRoutes.POST("/article", controller.ArticleAddNew)
	protectedRoutes.GET("/article/:ID", controller.ArticleFindById)
	protectedRoutes.GET("/article", controller.AllArticles)
	protectedRoutes.PATCH("/article/:ID/edit", controller.ArticleUpdate)

	// Estate
	protectedRoutes.POST("/estate", controller.EstateAddNew)
	protectedRoutes.PATCH("/estate/:ID/edit", controller.EstateUpdate)

	// Division
	protectedRoutes.POST("/division", controller.DivisionAddNew)

	//Corporation
	publicRoutes.POST("/company", controller.CompanyAddNew)
	publicRoutes.PATCH("/company/:ID/edit", controller.CompanyUpdate)

	//Application
	publicRoutes.POST("/application", controller.ApplicationAddNew)
	publicRoutes.PATCH("/application/:ID/edit", controller.ApplicationUpdate)

	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}

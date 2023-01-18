package main

import (
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

	protect "ddtservice_agri/api/protect"
	public "ddtservice_agri/api/public"

	_ "ddtservice_agri/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	database.Database.AutoMigrate(&schema.Company{})
	database.Database.AutoMigrate(&schema.Estate{})
	database.Database.AutoMigrate(&schema.Division{})
	database.Database.AutoMigrate(&schema.UserRole{})
	database.Database.AutoMigrate(&schema.UserDivision{})
	database.Database.AutoMigrate(&schema.Article{})
	database.Database.AutoMigrate(&schema.Application{})
	database.Database.AutoMigrate(&schema.ActivityLog{})
	database.Database.AutoMigrate(&schema.RoleApplication{})
	database.Database.AutoMigrate(&schema.Account{})
	database.Database.AutoMigrate(&schema.Application{})
	database.Database.AutoMigrate(&schema.AccountRoleApplication{})
	database.Database.AutoMigrate(&schema.Gang{})

}

func serveApplication() {
	router := gin.New()
	gin.DisableConsoleColor()
	// Logging to a file.
	f, _ := os.Create("logger.log")
	// Use the following code if you need to write the logs to file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router.Use(gin.Recovery())
	router.Use(middleware.JSONLogMiddleware())

	public.Authentication(router)
	router.GET("/swagger/auth/*any", ginSwagger.WrapHandler(swaggerfiles.NewHandler(), ginSwagger.InstanceName("auth")))
	protect.Protected(router)
	router.GET("/swagger/api/*any", ginSwagger.WrapHandler(swaggerfiles.NewHandler(), ginSwagger.InstanceName("api")))

	/*port := os.Getenv("HTTP_PLATFORM_PORT")
	if port == "" {
	    port = "8080"
	}*/

	port := os.Getenv("HTTP_PLATFORM_PORT")

	// default back to 8080 for local dev
	if port == "" {
		port = "8080"
	}

	router.Run("127.0.0.1:" + port)

}

package main

import (
	"ddtservice_agri/database"
	"ddtservice_agri/middleware"
	"ddtservice_agri/schema"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	protect "ddtservice_agri/api/protect"
	public "ddtservice_agri/api/public"

	_ "ddtservice_agri/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	loadDatabase()
	serveApplication()
	//loadEnv()

	//testGetDateTime()
}

// func testGetDateTime() {
// 	current_time := time.Now()

// 	// individual elements of time can
// 	// also be called to print accordingly
// 	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
// 		current_time.Year(), current_time.Month(), current_time.Day(),
// 		current_time.Hour(), current_time.Minute(), current_time.Second())

// 	// formatting time using
// 	// custom formats
// 	fmt.Println(current_time.Format("2006-01-02 15:04:05"))
// 	fmt.Println(current_time.Format("2006-January-02"))
// 	fmt.Println(current_time.Format("2006-01-02 3:4:5 pm"))
// }

func loadEnv() {
	if gin.Mode() == "debug" {
		err := godotenv.Load(".env.local")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	} else {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
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
	fmt.Println(os.Getenv("PORT"))
	fmt.Println(os.Getenv("GIN_MODE"))
	router := gin.New()
	gin.DisableConsoleColor()
	// Logging to a file.
	f, _ := os.Create("logger.log")
	// Use the following code if you need to write the logs to file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	router.Use(middleware.CORSMiddleware())
	router.Use(gin.Recovery())
	router.Use(middleware.JSONLogMiddleware())

	router.GET("/", GetHome)
	public.Authentication(router)
	router.GET("/swagger/auth/*any", ginSwagger.WrapHandler(swaggerfiles.NewHandler(), ginSwagger.InstanceName("auth")))
	protect.Protected(router)
	router.GET("/swagger/api/*any", ginSwagger.WrapHandler(swaggerfiles.NewHandler(), ginSwagger.InstanceName("api")))

	port := "8080"
	if os.Getenv("ASPNETCORE_PORT") != "" { // get enviroment variable that set by ACNM
		port = os.Getenv("ASPNETCORE_PORT")
	}
	fmt.Println(os.Getenv("GIN_ENV"))
	router.Run(fmt.Sprintf(":%v", port))
	// _ = router.Run()
}

func GetHome(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"Oke": "Status Server Don", "status": gin.Mode()})
}

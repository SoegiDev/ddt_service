package main

import (
	"ddtservice_agri/database"
	"ddtservice_agri/middleware"
	"ddtservice_agri/schema"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	protect "ddtservice_agri/api/protect"
	public "ddtservice_agri/api/public"

	_ "ddtservice_agri/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	gin.SetMode(gin.DebugMode)
	loadDatabase()
	serveApplication()
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

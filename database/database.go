package database

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() {
	var err error
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")
	//dsn: = "sqlserver://<user>:<password>@<server_host>?database=<database_name>"
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&encrypt=disable&connection+timeout=60", username, password, host, databaseName)
	fmt.Println(dsn)
	Database, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}
}

func Connect_PGSQL() {
	var err error
	var host, username, password, databaseName, port string
	if gin.Mode() == gin.DebugMode {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file Debug")
		}
		host = os.Getenv("DB_HOST")
		username = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		databaseName = os.Getenv("DB_NAME")
		port = os.Getenv("DB_PORT")
	} else {
		os.Setenv("DB_HOST", "localhost")
		os.Setenv("DB_USER", "postgres")
		os.Setenv("DB_PASSWORD", "password")
		os.Setenv("DB_NAME", "user_agri_prod")
		os.Setenv("DB_PORT", "5432")
		os.Setenv("GIN_ENV", "production")
		os.Setenv("SERVICE_NAME", "user-service")
		os.Setenv("TOKEN_TTL", "36000")
		os.Setenv("JWT_PRIVATE_KEY", "useragri_2023")
		host = os.Getenv("DB_HOST")
		username = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		databaseName = os.Getenv("DB_NAME")
		port = os.Getenv("DB_PORT")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Lagos", host, username, password, databaseName, port)
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}
}
